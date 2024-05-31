package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToInt(array []string) []int {
	var converted []int

	for _, n := range array {
		n, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			panic("Error converting element:" + err.Error())
		}
		converted = append(converted, n)
	}

	return converted
}

func formatMatrix(size int, data []int) [][]int {

	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			matrix[i][j] = data[i + j]
		}
	}

	return matrix
}

func readCSV(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(strings.NewReader(string(content)))
	reader.Comma = ';'

	matrix, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records := matrix[0]

	return records
}