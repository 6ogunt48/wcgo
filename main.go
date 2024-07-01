package main

import (
	"flag"
	"fmt"
	"wcgo/fileops"
)

func main() {
	ops := &fileops.FileOperations{}
	flag.BoolVar(&ops.ShowLines, "l", false, "show line count")
	flag.BoolVar(&ops.ShowWords, "w", false, "show word count")
	flag.BoolVar(&ops.ShowBytes, "c", false, "show byte count")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: wcgo [-l] [-w] [-c] <filename>")
		return
	}

	ops.Filename = flag.Args()[0]

	if !ops.ShowLines && !ops.ShowWords && !ops.ShowBytes {
		lines, words, bytes, err := ops.CountAll()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %8d %d %s\n", lines, words, bytes, ops.Filename)
		}
		return
	}

	if ops.ShowLines {
		lines, err := ops.CountLines()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", lines, ops.Filename)
		}
	}

	if ops.ShowWords {
		words, err := ops.CountWords()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", words, ops.Filename)
		}
	}

	if ops.ShowBytes {
		bytes, err := ops.GetFileByte()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%8d %s\n", bytes, ops.Filename)
		}
	}
}
