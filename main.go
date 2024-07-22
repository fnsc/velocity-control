package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fnsc/velocity-control/loader"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		request, err := loader.ParseRequest(line)
		if err != nil {
			fmt.Fprintf(outputFile, "Error parsing line: %s\n", line)
			continue
		}

		fmt.Println(request)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
