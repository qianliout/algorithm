package main

import (
	"log"
	"net/http"

	// 启用 pprof 性能分析
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/wolfogre/go-pprof-practice/animal"
)

func main() {
	// ...

	runtime.GOMAXPROCS(1)
	// 启用 mutex 性能分析
	runtime.SetMutexProfileFraction(1)
	// 启用 block 性能分析
	runtime.SetBlockProfileRate(1)

	go func() {
		// 启动 http server. 对应 pprof 的一系列 handler 也会挂载在该端口下
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	// 运行各项动物的活动
	for {
		for _, v := range animal.AllAnimals {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}
