package main

import (
	"fmt"
	"log"
	"os"
	"scoreboard/ops"
)

func main() {
	fmt.Println("ScoreBoard")

	if len(os.Args) == 0 {
		log.Println("Please add the file path as an argument")
		os.Exit(1)
	}

	// Read File
	results, err := ops.ReadResults(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Calculate Scores
	scoreBoard, err := ops.CalculateScores(results)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Print Scoreboard
	fmt.Println(scoreBoard.PrintScoreBoard())
}
