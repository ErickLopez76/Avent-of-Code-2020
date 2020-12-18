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
		fmt.Println("----------------Start Evaluation-------------")
		fmt.Println(scanner.Text())
		if evaluateString(scanner.Text()) {
			mainCount++
		}
	}
	fmt.Printf("FINAL RESULT:%d\n", mainCount)
}

func evaluateString(currentString string) bool {
	range_init, range_end := getRange(currentString)
	x, expresionEvaluate := getLetterAndExpresion(currentString)
	fmt.Printf("Expresion to Evaluate: %v\n", expresionEvaluate)
	runes := []rune(expresionEvaluate)
	firstLetter := string(runes[range_init-1 : range_init])
	//	fmt.Printf("pre error")
	if range_end > len(runes) {
		return false
	}
	lastLetter := string(runes[range_end-1 : range_end])

	fmt.Printf("firstLetter:%s , lastLetter:%s\n", firstLetter, lastLetter)

	if (firstLetter == x && lastLetter != x) || (firstLetter != x && lastLetter == x) {
		fmt.Printf("cuenta 1\n")
		return true
	}
	fmt.Printf("init:%d, end:%d\n", range_init, range_end)
	fmt.Printf("letter:%s\n", x)
	return false
}

func getLetterAndExpresion(curString string) (letter, expresion string) {
	fspace := strings.Index(curString, " ")
	runes := []rune(curString)

	letter = string(runes[fspace+1 : fspace+2])
	expresion = string(runes[fspace+4:])
	return
}

func getRange(curString string) (ninit, nend int) {
	fdash := strings.Index(curString, "-")
	fspace := strings.Index(curString, " ")
	runes := []rune(curString)

	ninit, _ = strconv.Atoi(string(runes[0:fdash]))
	nend, _ = strconv.Atoi(string(runes[fdash+1 : fspace]))

	return
}
