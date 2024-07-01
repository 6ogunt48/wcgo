package main

import (
	"flag"
	"fmt"
	"os"
	"wcgo/fileops"
)

func main() {
	ops := &fileops.FileOperations{}
	flag.BoolVar(&ops.ShowLines, "l", false, "show line count")
	flag.BoolVar(&ops.ShowWords, "w", false, "show word count")
	flag.BoolVar(&ops.ShowBytes, "c", false, "show byte count")
	flag.BoolVar(&ops.ShowChars, "m", false, "show character count")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: wcgo [-l] [-w] [-c] [-m] <filename>")
		return
	}

	ops.Filename = flag.Args()[0]

	file, err := os.Open(ops.Filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if !ops.ShowLines && !ops.ShowWords && !ops.ShowBytes && !ops.ShowChars {
		lines, words, bytes, chars, err := ops.CountAll(file)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %8d %8d %8d %s\n", lines, words, bytes, chars, ops.Filename)
		}
		return
	}

	if ops.ShowLines {
		lines, err := ops.CountLines(file)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", lines, ops.Filename)
		}
	}

	if ops.ShowWords {
		words, err := ops.CountWords(file)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", words, ops.Filename)
		}
	}

	if ops.ShowBytes {
		bytes, err := ops.GetFileByte(file)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", bytes, ops.Filename)
		}
	}

	if ops.ShowChars {
		chars, err := ops.CountChars(file)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", chars, ops.Filename)
		}
	}
}
