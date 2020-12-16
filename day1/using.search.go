package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("\n######## SearchInts not works in descending order  ######## ")
	intSlice := []int{55, 54, 53, 52, 51, 50, 48, 36, 15, 5} // sorted slice in descending
	x := 15
	pos := sort.SearchInts(intSlice, x)
	sort.Ints(intSlice)
	fmt.Printf("Found %d at index %d in %v\n", x, pos, intSlice)

	fmt.Println("\n######## Search works in descending order  ########")
	i := sort.Search(len(intSlice), func(i int) bool { return intSlice[i] == x })
	fmt.Printf("Found %d at index %d in %v\n", x, i, intSlice)

	fmt.Println("\n\n######## SearchStrings not works in descending order  ######## ")
	// sorted slice in descending
	strSlice := []string{"Washington", "Texas", "Ohio", "Nevada", "Montana", "Indiana", "Alaska"}
	y := "Montana"
	posstr := sort.SearchStrings(strSlice, y)
	fmt.Printf("Found %s at index %d in %v\n", y, posstr, strSlice)

	fmt.Println("\n######## Search works in descending order  ########")
	j := sort.Search(len(strSlice), func(j int) bool { return strSlice[j] <= y })
	fmt.Printf("Found %s at index %d in %v\n", y, j, strSlice)

	fmt.Println("\n######## Search works in ascending order  ########")
	fltSlice := []float64{10.10, 20.10, 30.15, 40.15, 58.95} // string slice in float64
	z := 40.15
	k := sort.Search(len(fltSlice), func(k int) bool { return fltSlice[k] >= z })
	fmt.Printf("Found %f at index %d in %v\n", z, k, fltSlice)
}
