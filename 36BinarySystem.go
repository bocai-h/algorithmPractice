package main

import "fmt"

var stringTo10BinarySystemMap map[string]int
var tenBinaryToString map[int]string

func initMap() {
	stringTo10BinarySystemMap = make(map[string]int)
	tenBinaryToString = make(map[int]string)
	originData := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i, item := range originData {
		stringTo10BinarySystemMap[item] = i
		tenBinaryToString[i] = item
	}
}
func main() {
	initMap()
	result := Add("1b5", "2x")
	fmt.Printf("result=%s\n", result)
}

//Add add
func Add(n string, m string) (result string) {
	nLen := len(n)
	mLen := len(m)
	bits := nLen - mLen
	if bits < 0 {
		bits = -bits
	}
	if nLen > mLen {
		for i := 0; i < bits; i++ {
			m = fmt.Sprintf("0%s", m)
		}
	} else {
		for i := 0; i < bits; i++ {
			n = fmt.Sprintf("0%s", n)
		}
	}

	carry := 0
	for j := len(n) - 1; j >= 0; j-- {
		njstr := fmt.Sprintf("%c", n[j])
		mjstr := fmt.Sprintf("%c", m[j])
		addition := stringTo10BinarySystemMap[njstr] + stringTo10BinarySystemMap[mjstr] + carry
		//clear carry
		carry = 0
		r := 0
		if addition >= 36 {
			carry = 1
			r = addition - 36
		} else {
			r = addition
		}
		result = fmt.Sprintf("%s%s", tenBinaryToString[r], result)
	}
	if carry > 0 {
		result = fmt.Sprintf("%s%s", 1, result)
	}
	return
}
