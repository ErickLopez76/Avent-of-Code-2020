package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fields := map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}

	var currentSlice []string
	var currentField []string
	mainCount := 0
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			currentText := string(scanner.Text())
			currentSlice = strings.Split(currentText, " ")
			//			fmt.Printf("%q\n", strings.Split(currentText, " "))

			for i := 0; i < len(currentSlice); i++ {
				currentField = strings.Split(string(currentSlice[i]), ":")
				//				fmt.Printf("currentField: %q\n", currentField)
				fields[string(currentField[0])] = string(currentField[1])

			}
			//fmt.Println(fields["pid"])
		}
		if len(scanner.Text()) == 0 {
			fmt.Println("new line")
			if fields["byr"] == "" || fields["iyr"] == "" || fields["eyr"] == "" || fields["hgt"] == "" || fields["hcl"] == "" || fields["ecl"] == "" || fields["pid"] == "" {
				//fmt.Println(fields["byr"])
				fmt.Println("No count")
			} else {
				mainCount++
				fmt.Printf("Count: %v\n", lineCount)
			}
			fmt.Println("-------------------- DATA ---------------")
			for key, value := range fields {
				fmt.Println(key, value)
			}
			for key, _ := range fields {
				//fmt.Println(value)
				fields[key] = ""
			}

		}
		fmt.Println(lineCount)
		lineCount++
	}
	if fields["byr"] == "" || fields["iyr"] == "" || fields["eyr"] == "" || fields["hgt"] == "" || fields["hcl"] == "" || fields["ecl"] == "" || fields["pid"] == "" {
		//fmt.Println(fields["byr"])
		fmt.Println("No count")
	} else {
		mainCount++
		fmt.Printf("Count: %v\n", lineCount)
	}
	fmt.Println("-------------------- DATA ---------------")
	for key, value := range fields {
		fmt.Println(key, value)
	}

	fmt.Printf("Total: %v\n", mainCount)
}
