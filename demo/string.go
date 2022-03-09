package demo

func maxOfStr(str string) int {
	lastAc := make(map[rune]int)
	maxLen := 0
	start := 0
	for i, ch := range []rune(str) {
		if lastI, ok := lastAc[ch]; ok && lastI >= start {
			start = lastAc[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

func StringDemo() {
	maxOfStr("你好呀好呀")
}
