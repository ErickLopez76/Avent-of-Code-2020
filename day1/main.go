package main

import (
	"fmt"
	"sort"
)

var myU = []int{1721, 979, 366, 299, 675, 1456}

func main() {
	//	sort.Slice(myU, func(i, j int) bool {
	//		return myU[i] < myU[j]
	//	})
	sort.Ints(myU)
	fmt.Println(myU)
	for _, v := range myU {
		diff2 := 2020 - v
		fmt.Println("Diff :%d", diff2)
		idx := sort.Search(len(myU), func(i int) bool { return diff2 <= myU[i] })
		if idx < len(myU) && myU[idx] == diff2 {
			fmt.Printf("Found\n")
			fmt.Println("Result: %d:", v*diff2)
		} else {
			fmt.Printf("Not Found\n")
		}

		fmt.Println("idx:", idx)
	}

}
