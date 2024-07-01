package fileops

import (
	"bufio"
	"os"
	"strings"
)

type FileOperations struct {
	Filename  string
	ShowLines bool
	ShowWords bool
	ShowBytes bool
}

func (ops *FileOperations) GetFileByte() (int64, error) {
	file, err := os.Open(ops.Filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func (ops *FileOperations) CountLines() (int, error) {
	file, err := os.Open(ops.Filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, scanner.Err()
}

func (ops *FileOperations) CountWords() (int, error) {
	file, err := os.Open(ops.Filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	return wordCount, scanner.Err()
}

func (ops *FileOperations) CountAll() (int, int, int64, error) {
	file, err := os.Open(ops.Filename)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()
	lineCount := 0
	wordCount := 0
	byteCount := int64(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		byteCount += int64(len(line)) + 1

		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)
		for lineScanner.Scan() {
			wordCount++
		}

		if err := lineScanner.Err(); err != nil {
			return 0, 0, 0, err
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return lineCount, wordCount, byteCount, nil
}
