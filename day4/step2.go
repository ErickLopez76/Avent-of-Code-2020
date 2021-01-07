package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	count := false
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
			if fields["byr"] == "" || fields["iyr"] == "" || fields["eyr"] == "" || fields["hgt"] == "" || fields["hcl"] == "" || fields["ecl"] == "" || fields["pid"] == "" {
				//fmt.Println(fields["byr"])
				count = false
				fmt.Println("No count")
			} else {
				count = true
				validIDcm := regexp.MustCompile(`^[1-9][0-9]*cm$`)
				validIDin := regexp.MustCompile(`^[1-9][0-9]*in$`)
				validHex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
				validColor := regexp.MustCompile(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`)
				validDig := regexp.MustCompile(`^[0-9]{9}$`)
				current_byr, _ := strconv.Atoi(fields["byr"])
				if len(fields["byr"]) != 4 || current_byr < 1920 || current_byr > 2002 {
					fmt.Println("byr fail!")
					count = false
				}

				current_iyr, _ := strconv.Atoi(fields["iyr"])
				if len(fields["iyr"]) != 4 || current_iyr < 2010 || current_iyr > 2020 {
					fmt.Println("iyr fail!")
					count = false
				}

				current_eyr, _ := strconv.Atoi(fields["eyr"])
				if len(fields["eyr"]) != 4 || current_eyr < 2020 || current_eyr > 2030 {
					fmt.Println("eyr fail!")
					count = false
				}

				if validIDcm.MatchString(fields["hgt"]) || validIDin.MatchString(fields["hgt"]) {

					if validIDcm.MatchString(fields["hgt"]) {
						hgtArray := strings.Split(fields["hgt"], "c")
						hgtInt, _ := strconv.Atoi(hgtArray[0])
						if hgtInt < 150 || hgtInt > 193 {
							fmt.Println("cm fail!")
							count = false
						}
					}
					if validIDin.MatchString(fields["hgt"]) {
						hgtArray := strings.Split(fields["hgt"], "i")
						hgtInt, _ := strconv.Atoi(hgtArray[0])
						if hgtInt < 59 || hgtInt > 76 {
							fmt.Println("in fail!")
							count = false
						}
					}
				} else {
					count = false
				}
				if !validHex.MatchString(fields["hcl"]) {
					fmt.Println("hcl fail!")
					count = false
				}

				if !validColor.MatchString(fields["ecl"]) {
					fmt.Println("ecl fail!")
					count = false
				}
				if !validDig.MatchString(fields["pid"]) {
					fmt.Println("pid fail!")
					count = false
				}

				if count {
					mainCount++
					fmt.Printf("Count: %v\n", lineCount)
				}
			}
			fmt.Println("-------------------- DATA ---------------")
			for key, value := range fields {
				fmt.Println(key, value)
			}
			//Clean fields
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
		fmt.Println("No count finish")
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

// regest for metric
// [1-9][0-9]*["cm$"]
