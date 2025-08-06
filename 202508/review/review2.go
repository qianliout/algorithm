package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// FrameType defines the type of a protocol frame.
const (
	FrameTypeData byte = 0
	FrameTypeEnd  byte = 1
)

// Frame represents a protocol data frame.
type Frame struct {
	Type      byte
	Key       string
	SessionID uint64
	Data      []byte
}

// Writer implements io.WriteCloser for sending data over a stream.
type Writer struct {
	conn      *Conn
	key       string
	sessionID uint64
	closed    bool
	mutex     sync.Mutex
}

// Write sends data as part of a stream.
func (w *Writer) Write(data []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.closed {
		return 0, io.ErrClosedPipe
	}

	// Data must be copied because the caller might reuse the buffer.
	dataCopy := make([]byte, len(data))
	copy(dataCopy, data)

	frame := &Frame{
		Type:      FrameTypeData,
		Key:       w.key,
		SessionID: w.sessionID,
		Data:      dataCopy,
	}

	select {
	case w.conn.sendChan <- frame:
		return len(data), nil
	case <-w.conn.closeChan:
		return 0, io.ErrClosedPipe
	}
}

// Close signals the end of the stream.
func (w *Writer) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.closed {
		return nil
	}
	w.closed = true

	frame := &Frame{
		Type:      FrameTypeEnd,
		Key:       w.key,
		SessionID: w.sessionID,
	}

	select {
	case w.conn.sendChan <- frame:
		return nil
	case <-w.conn.closeChan:
		return io.ErrClosedPipe
	}
}

// Stream represents a single, logical flow of data.
type Stream struct {
	key        string
	sessionID  uint64
	pipeReader *io.PipeReader
	pipeWriter *io.PipeWriter
}

// newStream creates a new stream with an associated io.Pipe.
func newStream(key string, sessionID uint64) *Stream {
	pr, pw := io.Pipe()
	return &Stream{
		key:        key,
		sessionID:  sessionID,
		pipeReader: pr,
		pipeWriter: pw,
	}
}

// sessionManager manages all active streams for a connection.
type sessionManager struct {
	sync.RWMutex
	streams        map[uint64]*Stream
	newStreamsChan chan *Stream
	sessionCounter uint64
	closed         bool
}

// newSessionManager creates a new session manager.
func newSessionManager() *sessionManager {
	return &sessionManager{
		streams:        make(map[uint64]*Stream),
		newStreamsChan: make(chan *Stream, 256), // Buffered channel to prevent deadlocks
	}
}

// dispatchFrame finds the appropriate stream and forwards the frame data.
func (sm *sessionManager) dispatchFrame(frame *Frame) {
	sm.RLock()
	stream, exists := sm.streams[frame.SessionID]
	sm.RUnlock()

	if !exists {
		if frame.Type == FrameTypeEnd {
			// END frame for a non-existent or already closed stream, ignore.
			return
		}

		sm.Lock()
		// Double-checked lock pattern to prevent race condition on stream creation
		if stream, exists = sm.streams[frame.SessionID]; !exists {
			stream = newStream(frame.Key, frame.SessionID)
			sm.streams[frame.SessionID] = stream
			// Unlock before sending to channel to avoid holding lock during potential block
			sm.Unlock()
			sm.newStreamsChan <- stream
		} else {
			sm.Unlock()
		}
	}

	if len(frame.Data) > 0 {
		// This provides backpressure if the user is not reading from the pipe.
		if _, err := stream.pipeWriter.Write(frame.Data); err != nil {
			// Reader side of the pipe was closed, likely by the user.
			sm.removeStream(stream.sessionID)
			stream.pipeWriter.CloseWithError(err)
		}
	}

	if frame.Type == FrameTypeEnd {
		stream.pipeWriter.Close() // Signals EOF to the reader
		sm.removeStream(stream.sessionID)
	}
}

// removeStream removes a stream from the manager.
func (sm *sessionManager) removeStream(sessionID uint64) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.streams, sessionID)
}

// closeAll closes all active streams and stops accepting new ones.
func (sm *sessionManager) closeAll(err error) {
	sm.Lock()
	defer sm.Unlock()
	if sm.closed {
		return
	}
	sm.closed = true
	for _, stream := range sm.streams {
		stream.pipeWriter.CloseWithError(err)
	}
	sm.streams = make(map[uint64]*Stream)
	close(sm.newStreamsChan)
}

// Conn is a multiplexed connection over a single TCP connection.
type Conn struct {
	conn net.Conn

	wg        sync.WaitGroup
	closeOnce sync.Once
	closeChan chan struct{}

	sendChan chan *Frame
	sessions *sessionManager
}

// NewConn creates a new multiplexed connection from a standard net.Conn.
func NewConn(conn net.Conn) *Conn {
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetNoDelay(true)
		tcpConn.SetReadBuffer(1 << 20)
		tcpConn.SetWriteBuffer(1 << 20)
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(30 * time.Second)
	}

	c := &Conn{
		conn:      conn,
		closeChan: make(chan struct{}),
		sendChan:  make(chan *Frame, 256),
		sessions:  newSessionManager(),
	}

	c.wg.Add(2)
	go c.readLoop()
	go c.writeLoop()

	return c
}

// readLoop continuously reads from the underlying connection and dispatches frames.
func (c *Conn) readLoop() {
	defer c.wg.Done()
	defer c.Close() // If read fails, close the whole connection
	reader := bufio.NewReader(c.conn)
	for {
		frame, err := readFrame(reader)
		if err != nil {
			// Any error (including io.EOF from a closed conn) is fatal.
			return
		}

		select {
		case <-c.closeChan:
			return
		default:
			c.sessions.dispatchFrame(frame)
		}
	}
}

// writeLoop continuously sends frames from the send channel to the connection.
func (c *Conn) writeLoop() {
	defer c.wg.Done()
	writer := bufio.NewWriter(c.conn)
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case frame, ok := <-c.sendChan:
			if !ok {
				writer.Flush()
				return
			}
			if err := writeFrame(writer, frame); err != nil {
				c.Close()
				return
			}
			if len(c.sendChan) == 0 {
				if err := writer.Flush(); err != nil {
					c.Close()
					return
				}
			}
		case <-ticker.C:
			if err := writer.Flush(); err != nil {
				c.Close()
				return
			}
		case <-c.closeChan:
			// Drain remaining frames on shutdown
			for len(c.sendChan) > 0 {
				frame := <-c.sendChan
				if err := writeFrame(writer, frame); err != nil {
					break
				}
			}
			writer.Flush()
			return
		}
	}
}

// Send prepares a new stream for sending data.
func (c *Conn) Send(key string) (io.WriteCloser, error) {
	select {
	case <-c.closeChan:
		return nil, io.ErrClosedPipe
	default:
	}

	sessionID := atomic.AddUint64(&c.sessions.sessionCounter, 1)
	return &Writer{
		conn:      c,
		key:       key,
		sessionID: sessionID,
	}, nil
}

// Receive waits for and returns the next incoming stream.
func (c *Conn) Receive() (string, io.Reader, error) {
	select {
	case stream, ok := <-c.sessions.newStreamsChan:
		if !ok {
			return "", nil, io.EOF
		}
		return stream.key, stream.pipeReader, nil
	case <-c.closeChan:
		return "", nil, io.EOF
	}
}

// Close gracefully shuts down the connection.
func (c *Conn) Close() {
	c.closeOnce.Do(func() {
		close(c.closeChan) // Signal goroutines to stop
		c.conn.Close()     // Unblock blocking I/O operations
		c.wg.Wait()        // Wait for goroutines to finish
		c.sessions.closeAll(io.ErrClosedPipe)
		close(c.sendChan)
	})
}

// writeFrame serializes and writes a frame to the given writer.
func writeFrame(writer io.Writer, frame *Frame) error {
	keyBytes := []byte(frame.Key)
	keyLen := uint32(len(keyBytes))
	dataLen := uint32(len(frame.Data))

	// Header: Type (1) + KeyLen (4) + SessionID (8) + DataLen (4) = 17 bytes
	header := make([]byte, 17)
	header[0] = frame.Type
	binary.BigEndian.PutUint32(header[1:5], keyLen)
	binary.BigEndian.PutUint64(header[5:13], frame.SessionID)
	binary.BigEndian.PutUint32(header[13:17], dataLen)

	if _, err := writer.Write(header); err != nil {
		return err
	}
	if _, err := writer.Write(keyBytes); err != nil {
		return err
	}
	if _, err := writer.Write(frame.Data); err != nil {
		return err
	}
	return nil
}

// readFrame reads and deserializes a frame from the given reader.
func readFrame(reader io.Reader) (*Frame, error) {
	header := make([]byte, 17)
	if _, err := io.ReadFull(reader, header); err != nil {
		return nil, err
	}

	frame := &Frame{}
	frame.Type = header[0]
	keyLen := binary.BigEndian.Uint32(header[1:5])
	frame.SessionID = binary.BigEndian.Uint64(header[5:13])
	dataLen := binary.BigEndian.Uint32(header[13:17])

	// Protection against malicious/corrupted frames
	if keyLen > 4096 || dataLen > (64<<20) { // 4KB key limit, 64MB data limit per frame
		return nil, fmt.Errorf("frame size limits exceeded: key=%d, data=%d", keyLen, dataLen)
	}

	keyBytes := make([]byte, keyLen)
	if _, err := io.ReadFull(reader, keyBytes); err != nil {
		return nil, err
	}
	frame.Key = string(keyBytes)

	frame.Data = make([]byte, dataLen)
	if _, err := io.ReadFull(reader, frame.Data); err != nil {
		return nil, err
	}

	return frame, nil
}

// ////////////////////////////////////////////
// /////// 接下来的代码为测试代码，请勿修改 /////////
// ////////////////////////////////////////////

// 连接到测试服务器，获得一个你实现的连接对象
func dial(serverAddr string) *Conn {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	return NewConn(conn)
}

// 启动测试服务器
func startServer(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("[WARNING] ln.Accept", err)
				return
			}
			go handle(NewConn(conn))
		}
	}()
	return ln
}

// 简单断言
func assertEqual[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("actual: %v, expected: %v", actual, expected))
	}
}

// 简单 case：单连接，双向传输少量数据
func testCase0() {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, "Whom shall I send? And who will go for us?"
And I said, "Here am I. Send me!"
Isaiah 6:8`
	)

	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(n)
		}
		err = writer.Close()
		if err != nil {
			panic(err)
		}

		// Wait for the client to close the connection, indicating it has received the data.
		// A receive call will block until a new stream arrives or the connection is closed by the peer.
		_, _, err = conn.Receive()
		if err != io.EOF {
			panic(fmt.Sprintf("server expected EOF from client, but got: %v", err))
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	assertEqual(string(dataB), data)
	conn.Close()
}

// 生成一个随机 key
func newRandomKey() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

// 读取随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func readRandomData(reader io.Reader, hash hash.Hash) (checksum string) {
	hash.Reset()
	var buf = make([]byte, 23<<20) // 调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 写入随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func writeRandomData(writer io.Writer, hash hash.Hash) (checksum string) {
	hash.Reset()
	const (
		dataSize = 500 << 20 // 一个 key 对应 500MB 随机二进制数据，dataSize 也可以是其他值，你的实现中不可假定 dataSize 为固定值
		bufSize  = 1 << 20   // 调用者写入时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	)
	var (
		buf  = make([]byte, bufSize)
		size = 0
	)
	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
		size += n
	}
	if size != dataSize {
		panic(size)
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 复杂 case：多连接，双向传输，大量数据，多个不同的 key
func testCase1() {
	var (
		mapKeyToChecksum = map[string]string{}
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		var (
			h         = sha256.New()
			_checksum = readRandomData(reader, h)
		)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)

		// 服务端向客户端连续进行 2 次传输
		for _, key := range []string{newRandomKey(), newRandomKey()} {
			writer, err := conn.Send(key)
			if err != nil {
				panic(err)
			}
			checksum := writeRandomData(writer, h)
			lock.Lock()
			mapKeyToChecksum[key] = checksum
			lock.Unlock()
			err = writer.Close() // 表明该 key 的所有数据已传输完毕
			if err != nil {
				panic(err)
			}
		}

		// Wait for the client to close the connection.
		_, _, err = conn.Receive()
		if err != io.EOF {
			panic(fmt.Sprintf("server expected EOF from client, but got: %v", err))
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String())
	// 客户端向服务端传输
	var (
		key = newRandomKey()
		h   = sha256.New()
	)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	checksum := writeRandomData(writer, h)
	lock.Lock()
	mapKeyToChecksum[key] = checksum
	lock.Unlock()
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 客户端等待服务端的多次传输
	keyCount := 0
	for {
		key, reader, err := conn.Receive()
		if err == io.EOF {
			// 服务端所有的数据均传输完毕，关闭连接
			break
		}
		if err != nil {
			panic(err)
		}
		_checksum := readRandomData(reader, h)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)
		keyCount++
	}
	assertEqual(keyCount, 2)
	conn.Close()
}

func main() {
	fmt.Println("=== Running Test Case 0 ===")
	testCase0()
	fmt.Println("Test Case 0 passed!")

	fmt.Println("=== Running Test Case 1 ===")
	testCase1()
	fmt.Println("Test Case 1 passed!")

	fmt.Println("All tests passed!")
}
