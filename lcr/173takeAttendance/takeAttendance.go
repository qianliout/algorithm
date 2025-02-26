package main

func main() {

}

func takeAttendance(records []int) int {

	// records = append(records, len(records))
	n := len(records)
	for i := 0; i < n; i++ {
		for records[records[i]] != records[i] {
			records[records[i]], records[i] = records[i], records[records[i]]
		}
	}
	for i, ch := range records {
		if i != ch {
			return i
		}
	}
	return n - 1
}
