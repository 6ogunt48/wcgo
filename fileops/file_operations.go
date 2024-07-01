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
	ShowChars bool
}

func (ops *FileOperations) GetFileByte(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func (ops *FileOperations) CountLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount, scanner.Err()
}

func (ops *FileOperations) CountWords(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount, scanner.Err()
}

func (ops *FileOperations) CountChars(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}
	return charCount, scanner.Err()
}

func (ops *FileOperations) CountAll(file *os.File) (int, int, int64, int, error) {
	lineCount := 0
	wordCount := 0
	byteCount := int64(0)
	charCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		byteCount += int64(len(line)) + 1 // Adding 1 for the newline character
		charCount += len(line)

		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)
		for lineScanner.Scan() {
			wordCount++
		}
		if err := lineScanner.Err(); err != nil {
			return 0, 0, 0, 0, err
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, 0, err
	}

	return lineCount, wordCount, byteCount, charCount, nil
}
