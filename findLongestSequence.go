package main

import "fmt"

func main() {
	str := "ababcdefefddg134567890"
	len := maxUnique(str)
	fmt.Printf("str=%s, maxLongestSequence=%v\n", str, len)
}

func maxUnique(str string) (len int) {
	//pmap记录着遍历到i位置为止最近一次i出现的位置
	pmap := make(map[rune]int)
	//pre表示遍历到i位置为止,pre代表以str[i-1]结尾的最长无重复子串开始位置的前一个位置
	pre := -1
	//cur记录着当前长度
	cur := 0
	for i, s := range str {
		pre = mathMax(pre, pmap[s])
		cur = i - pre
		len = mathMax(len, cur)
		pmap[s] = i
	}
	return
}

func mathMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
