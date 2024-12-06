package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		dayStr := line[:2]
		day, err := strconv.Atoi(dayStr)
		if err != nil {
			continue
		}
		if start.Day() <= day && day <= end.Day() {
			result = append(result, line)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка сканирования файла")
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("ни одна строка не попала в указанный диапазон")
	}

	return result, nil
}
