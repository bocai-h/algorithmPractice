package main

import "fmt"

func main() {
	sources := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	dp := getDp1(sources)
	des := generateLIS(sources, dp)
	for _, item := range des {
		fmt.Printf("%v ", item)
	}
}

func getDp1(sArray []int) (dp []int) {
	//dp record when index to i end, max asc length
	dp = make([]int, len(sArray))
	for i, sitem := range sArray {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if sArray[j] < sitem {
				dp[i] = mathMax(dp[i], dp[j]+1)
			}
		}
	}
	return dp
}

// generate Longest increasing Sequence
func generateLIS(sArray []int, dp []int) (des []int) {
	//find max dp element
	len := 0
	index := 0
	for i, ditem := range dp {
		if ditem > len {
			len = ditem
			index = i
		}
	}
	des = make([]int, len)

	len--
	des[len] = sArray[index]
	for i := index; i >= 0; i-- {
		if sArray[i] < sArray[index] && dp[i] == dp[index]-1 {
			len--
			des[len] = sArray[i]
			index = i
		}
	}
	return des
}

func mathMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// func findMaxAscSeq(sources *[]int) (des []int) {
// 	sLen := len(*sources)
// 	des = make([]int, 0, sLen)
// 	d := make([]int, sLen)
// 	for i, n := range *sources {
// 		d[i] = 1
// 		for j := 0; j < i; j++ {
// 			if n > (*sources)[j] {
// 				if d[i] < d[j]+1 {
// 					d[i] = d[j] + 1
// 				}
// 			}
// 		}
// 	}
// 	fmt.Printf("d=%v\n", d)
// 	maxCount := -1
// 	maxI := -1
// 	for i, count := range d {
// 		if count > maxCount {
// 			maxCount = count
// 			maxI = i
// 		}
// 	}
// 	curMaxNum := (*sources)[maxI]
// 	curI := maxI

// 	des = append(des, (*sources)[curI])

// 	for i := maxI - 1; i >= 0; i-- {
// 		if (*sources)[i] < curMaxNum && d[i] == d[curI]-1 {
// 			des = append(des, (*sources)[i])
// 			curI = i
// 			curMaxNum = (*sources)[i]
// 		}
// 	}
// 	return
// }
