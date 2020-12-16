package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		println("Error, cant read file")
	}
	defer file.Close()
	var mainCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if evaluateString(scanner.Text()) {
			mainCount++
		}
	}
	fmt.Printf("FINAL RESULT:%d", mainCount)
}

func evaluateString(currentString string) bool {
	range_init, range_end := getRange(currentString)
	x := getLetter(currentString)
	count := strings.Count(currentString, x) - 1
	if count >= range_init && count <= range_end {
		return true
	}
	fmt.Printf("init:%d, end:%d\n", range_init, range_end)
	fmt.Printf("letter:%s\n", x)
	return false
}

func getLetter(curString string) string {
	fspace := strings.Index(curString, " ")
	runes := []rune(curString)

	return string(runes[fspace+1 : fspace+2])
}

func getRange(curString string) (ninit, nend int) {
	fdash := strings.Index(curString, "-")
	fspace := strings.Index(curString, " ")
	runes := []rune(curString)

	ninit, _ = strconv.Atoi(string(runes[0:fdash]))
	nend, _ = strconv.Atoi(string(runes[fdash+1 : fspace]))

	return
}
