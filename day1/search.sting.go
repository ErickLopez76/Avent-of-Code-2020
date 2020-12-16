package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"A", "C", "C"}
	x := "C"

	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		fmt.Printf("Found %s at index %d in %v.\n", x, i, a)
	} else {
		fmt.Printf("Did not find %s in %v.\n", x, a)
	}

	num := []int{991, 354, 197, 687, 713, 509}
	sort.Ints(num)
	xn := 71
	j := sort.Search(len(num), func(i int) bool { return xn <= num[i] })
	if j < len(num) && num[j] == xn {
		fmt.Printf("Found %d at index %d in %v.\n", xn, j, num)
	} else {
		fmt.Printf("Did not find %d in %v.\n", xn, num)
	}
}
