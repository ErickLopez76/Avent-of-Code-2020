package main

import (
	"fmt"
	"regexp"
    "strconv"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	//var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	var validID = regexp.MustCompile(`^[1-9][0-9]*cm$`)

// [1-9][0-9]*["cm$"]
	fmt.Println(validID.MatchString("11m"))
	fmt.Println(validID.MatchString("11"))
	fmt.Println(validID.MatchString("123cm"))
	fmt.Println(validID.MatchString("11cm"))
    i, _ := strconv.Atoi("42ababa")
	fmt.Println(i)
}

