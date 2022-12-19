package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFileName := "input.txt"
	data, err := os.ReadFile(inputFileName)
	check(err)
	lines := strings.Split(string(data), "\n")
	fmt.Println(lines)
}
