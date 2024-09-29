package main

func main() {

}

func getEncryptedString(s string, k int) string {
	ss := []byte(s)
	n := len(s)
	k = k % n

	bre := ss[k:]
	after := ss[:k]
	bre = append(bre, after...)
	return string(bre)
}
