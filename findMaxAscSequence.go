package main

import "fmt"

func main() {
	sources := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	des := findMaxAscSeq(&sources)
	for i := len(des) - 1; i >= 0; i-- {
		fmt.Printf("%v ", des[i])
	}
}

func findMaxAscSeq(sources *[]int) (des []int) {
	sLen := len(*sources)
	des = make([]int, 0, sLen)
	d := make([]int, sLen)
	for i, n := range *sources {
		d[i] = 1
		for j := 0; j < i; j++ {
			if n > (*sources)[j] {
				if d[i] < d[j]+1 {
					d[i] = d[j] + 1
				}
			}
		}
	}
	fmt.Printf("d=%v\n", d)
	maxCount := -1
	maxI := -1
	for i, count := range d {
		if count > maxCount {
			maxCount = count
			maxI = i
		}
	}
	curMaxNum := (*sources)[maxI]
	curI := maxI

	des = append(des, (*sources)[curI])

	for i := maxI - 1; i >= 0; i-- {
		if (*sources)[i] < curMaxNum && d[i] == d[curI]-1 {
			des = append(des, (*sources)[i])
			curI = i
			curMaxNum = (*sources)[i]
		}
	}
	return
}
