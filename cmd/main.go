package main

import {
	"fmt"
	"os"
	"encoding/json"
}

func parseTransaction(line string) (string, error) {
    var tx models.Transaction
    err := json.Unmarshal([]byte(line), &tx)

	if err != nil {
        return tx, fmt.Errorf("error parsing JSON: %v", err)
    }

	tx.Time, err = time.Parse(time.RFC3339, tx.Time.Format(time.RFC3339))

	if err != nil {
        return tx, fmt.Errorf("error parsing time: %v", err)
    }
    return tx, nil
}

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
		transaction, err := parseTransaction(line)

		if err != nil {
			fmt.Println("Error parsing transaction:", err)

			continue
		}

		fmt.Println(string(transaction))
	}
}
