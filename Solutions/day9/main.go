package main

import (
	"errors"
	"fmt"
	"main/lib"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Provide [part (1 or 2)] [filename]")
		return
	}
	part := os.Args[1]
	inputFile := os.Args[2]

	switch part {
	case "1":
		result, err := lib.Day9Part1(inputFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)

	case "2":
		result, err := 0, errors.New("not implemented") // Implement actual solution here
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(result)

	default:
		fmt.Println("Theres only part 1 and 2 😾")
	}
}
