package main

import (
	"fmt"
	"markov"
	"os"
	"strconv"
)

func main() {
	file := os.Args[1]
	length, err := strconv.Atoi(os.Args[2])
	outputName := os.Args[3]
	flag := os.Args[4]

	if err != nil {
		fmt.Println("There was a problem converting the length to an integer.")
	}

	chain, stats := markov.Train(file)
	output := markov.Generate(chain, length)

	outputFile, err := os.Create(outputName)
	defer outputFile.Close()

	fmt.Fprintf(outputFile, output)
	fmt.Printf("Source word count: %d, unique keys: %d.\n", stats[0], stats[1])

	if flag == "-stats" {
		fmt.Fprintf(outputFile, "\n\nSource word count: %d, unique keys: %d.\n\n", stats[0], stats[1])

		for key, item := range chain	{
			fmt.Fprintf(outputFile, "[%s]:%v\n", key, item)
		}
	}
}
