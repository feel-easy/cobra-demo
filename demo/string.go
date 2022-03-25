package demo

import "fmt"

func maxOfStr(str string) (int, string) {
	lastAc := make(map[rune]int)
	maxLen := 0
	start := 0
	maxStr := ""
	runeStr := []rune(str)
	for i, ch := range runeStr {
		if lastI, ok := lastAc[ch]; ok && lastI >= start {
			start = lastAc[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
			maxStr = string(runeStr[start : i+1])
		}
	}
	return maxLen, maxStr
}

func StringDemo() {
	fmt.Println(maxOfStr("hello"))
	h := make([]int, 10, 12)
	fmt.Println(h)
}
