package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fnsc/velocity-control/handlers"
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

	dailyLoadCountHandler := setHandlers()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		request, err := loader.ParseRequest(line)
		if err != nil {
			fmt.Fprintf(outputFile, "Error parsing line: %s\n", line)
			continue
		}

		response := dailyLoadCountHandler.Handle(request)
		reponseJson, err := json.Marshal(response)
		if err != nil {
			fmt.Fprintf(outputFile, "error generating response for line: %s", line)

			continue
		}

		outputFile.WriteString(string(reponseJson) + "\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func setHandlers() *handlers.DailyLoadCountHandler {
	dailyLimitHandler := handlers.NewDailyLimitHandler()
	weeklyLimitHandler := handlers.NewWeeklyLimitHandler()
	dailyLoadCountHandler := handlers.NewDailyLoadCountHandler()

	dailyLoadCountHandler.SetNext(&weeklyLimitHandler.BaseHandler)
	weeklyLimitHandler.SetNext(&dailyLimitHandler.BaseHandler)

	return dailyLoadCountHandler
}
