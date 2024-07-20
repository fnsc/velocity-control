package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			fmt.Println("Error parsing transaction:", err)

			continue
		}

		fmt.Println(string(line))
	}
}
