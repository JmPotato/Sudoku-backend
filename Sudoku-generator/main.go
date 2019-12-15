package main

import (
	"fmt"

	generator "github.com/zwt/Sudoku-generator/generator"
)

func main() {
	board := generator.GenerateQuestion(25)
	for _, row := range board {
		fmt.Printf("%s", row)
	}
}
