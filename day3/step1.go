package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//	defer file.Close()
	scanner := bufio.NewScanner(file)
	var array_file []string
	for scanner.Scan() {
		array_file = append(array_file, scanner.Text())
		//		array_file = scanner.
	}
	A := evaluateArray(array_file, 1, 1)
	B := evaluateArray(array_file, 3, 1)
	C := evaluateArray(array_file, 5, 1)
	D := evaluateArray(array_file, 7, 1)
	E := evaluateArray(array_file, 1, 2)

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("D:", D)
	fmt.Println("E:", E)
	println("result: ", A*B*C*D*E)
}

func evaluateArray(array_file []string, xp int, yp int) int {
	var mainCount int = 0
	lenLine := len(array_file[0])
	x := xp
	y := yp
	for y < len(array_file) {
		fmt.Printf("x:%d, y:%d\n", x, y)
		fmt.Println("print: ", string(array_file[y][x]))
		if string(array_file[y][x]) == "#" {
			mainCount++
		}
		x = (x + xp) % lenLine
		y = y + yp
	}
	return mainCount
}
