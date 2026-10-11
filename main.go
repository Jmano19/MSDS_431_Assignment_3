package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <input.csv> <output.jl>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	if inputPath == outputPath {
		fmt.Fprintln(os.Stderr, "Error: input and output paths must differ.")
		os.Exit(1)
	}

	if err := ConvertCSVToJSONL(inputPath, outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Conversion failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversion successful: %s -> %s\n",
		inputPath, outputPath)
}
