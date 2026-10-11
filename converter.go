package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type HousingRecord struct {
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      float64 `json:"age"`
	Rooms    float64 `json:"rooms"`
	Bedrooms float64 `json:"bedrooms"`
	Pop      float64 `json:"pop"`
	HH       float64 `json:"hh"`
}

func ConvertCSVToJSONL(inputPath, outputPath string) error {
	input, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	defer input.Close()

	reader := csv.NewReader(input)
	reader.TrimLeadingSpace = true

	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("read CSV header: %w", err)
	}

	columns, err := findHousingColumns(headers)
	if err != nil {
		return err
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}

	// Remove an incomplete output file if conversion fails.
	defer output.Close()

	lineNumber := 1
	recordCount := 0

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		lineNumber++

		if err != nil {
			return fmt.Errorf("read CSV line %d: %w",
				lineNumber, err)
		}

		record, err := parseHousingRecord(row, columns)
		if err != nil {
			return fmt.Errorf("CSV line %d: %w",
				lineNumber, err)
		}

		encoded, err := json.Marshal(record)
		if err != nil {
			return fmt.Errorf("encode line %d: %w",
				lineNumber, err)
		}

		if _, err := output.Write(append(encoded, '\n')); err != nil {
			return fmt.Errorf("write output: %w", err)
		}

		recordCount++
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("close output: %w", err)
	}

	fmt.Printf("Converted %d records.\n", recordCount)
	return nil
}

func findHousingColumns(headers []string) (map[string]int, error) {
	aliases := map[string][]string{
		"value":    {"value", "median_house_value"},
		"income":   {"income", "median_income"},
		"age":      {"age", "housing_median_age"},
		"rooms":    {"rooms", "total_rooms"},
		"bedrooms": {"bedrooms", "total_bedrooms"},
		"pop":      {"pop", "population"},
		"hh":       {"hh", "households"},
	}

	indexes := make(map[string]int)
	for i, header := range headers {
		indexes[strings.ToLower(strings.TrimSpace(header))] = i
	}

	columns := make(map[string]int)
	for field, names := range aliases {
		found := false
		for _, name := range names {
			if index, ok := indexes[name]; ok {
				columns[field] = index
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("required column %q is missing", field)
		}
	}

	return columns, nil
}

func parseHousingRecord(row []string, columns map[string]int) (HousingRecord, error) {
	values := make(map[string]float64)

	for field, index := range columns {
		if index >= len(row) {
			return HousingRecord{}, fmt.Errorf(
				"missing value for %q", field)
		}

		value, err := strconv.ParseFloat(strings.TrimSpace(row[index]), 64)
		if err != nil {
			return HousingRecord{}, fmt.Errorf(
				"invalid number for %q: %q", field, row[index])
		}

		values[field] = value
	}

	return HousingRecord{
		Value:    values["value"],
		Income:   values["income"],
		Age:      values["age"],
		Rooms:    values["rooms"],
		Bedrooms: values["bedrooms"],
		Pop:      values["pop"],
		HH:       values["hh"],
	}, nil
}
