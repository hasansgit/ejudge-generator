package main

import (
	test_gen "ejudge-generator/internal/test-gen"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Test generator v0.3")
	fmt.Println("Welcome!")

	var genScript string
	fmt.Print("Put script file or path here:\t")
	fmt.Scanln(&genScript)
	genScript = strings.Trim(genScript, "'")

	var testQuantity int
	fmt.Print("Test quantity:\t")
	fmt.Scanln(&testQuantity)

	var solution string
	fmt.Print("Put solution if you have:\t")
	fmt.Scanln(&solution)
	solution = strings.Trim(solution, "'")

	test_gen.Gen(genScript, solution, testQuantity)
}
