package test_gen

import (
	"fmt"
	"os"
	"os/exec"
)

func genAns(filename, solution string) {
	inputFile, _ := os.Open(filename + ".dat")
	defer inputFile.Close()
	outputFile, _ := os.Create(filename + ".ans")
	defer outputFile.Close()

	cmd := exec.Command(solution)
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile
	err := cmd.Run()
	if err != nil {
		os.Remove(filename + ".ans")
	}
}

func Gen(genScript, solution string, testQuantity int) {
	os.Mkdir("tests", 0755)
	for i := 1; i <= testQuantity; i++ {
		filename := fmt.Sprintf("tests/%03d", i)
		os.WriteFile(filename+".dat", []byte(GenData(genScript)), 0644)
		genAns(filename, solution)
	}
}
