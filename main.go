package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fnsc/velocity-control/domain"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Error opening file: %s", err)

		return
	}

	defer file.Close()
	var requests []domain.Request

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var request domain.Request

		err := json.Unmarshal(scanner.Bytes(), &request)

		if err != nil {
			fmt.Println("error while deserializing your json:", err)

			continue
		}

		requests = append(requests, request)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading the file:", err)
	}

	for _, request := range requests {
		fmt.Println("%+v\n", request)
	}
}
