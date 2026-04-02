package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath, outputPath := "", ""
	if len(os.Args) >= 2 {
		inputPath = os.Args[1]
	}
	if len(os.Args) >= 3 {
		outputPath = os.Args[2]
	}

	scanner, inputFile, err := openInput(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open input: %v\n", err)
		os.Exit(1)
	}
	if inputFile != nil {
		defer inputFile.Close()
	}

	output, err := openOutput(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open output: %v\n", err)
		os.Exit(1)
	}
	if output != os.Stdout {
		defer output.Close()
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens := tokenize(line)
		codes, err := compileLine(tokens)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on line %q: %v\n", line, err)
			os.Exit(1)
		}
		if len(codes) == 0 {
			continue
		}

		fmt.Fprintln(output, formatLine(codes))
	}

	fmt.Fprintln(output, "0")
}

func openInput(path string) (*bufio.Scanner, *os.File, error) {
	if path == "" {
		return bufio.NewScanner(os.Stdin), nil, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewScanner(f), f, nil
}

func openOutput(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	}
	return os.Create(path)
}

func formatLine(codes []int) string {
	parts := make([]string, len(codes))
	for i, c := range codes {
		parts[i] = strconv.Itoa(c)
	}
	return strings.Join(parts, " ")
}
