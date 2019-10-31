package learngo

func cases(str string) int {

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(str) {

		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength

}
func main() {
	cases("asdfsdfdsf")
}
