package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Record map[string]string

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: csv2md <input-file> <output-file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	records, err := readRecords(inputFile)
	if err != nil {
		panic(err)
	}

	switch filepath.Ext(outputFile) {
	case ".json":
		err = writeJSON(records, outputFile)
	case ".md":
		err = writeMarkdown(records, outputFile)
	default:
		fmt.Println("Unsupported output file format")
		return
	}

	if err != nil {
		panic(err)
	}
}

func readRecords(inputFile string) ([]Record, error) {
	var records []Record
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch filepath.Ext(inputFile) {
	case ".csv":
		reader := csv.NewReader(file)
		headers, err := reader.Read()
		if err != nil {
			return nil, err
		}

		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}

			record := make(Record)
			for i, header := range headers {
				record[header] = row[i]
			}
			records = append(records, record)
		}
	case ".xlsx":
		f, err := excelize.OpenReader(file)
		if err != nil {
			return nil, err
		}

		rows, err := f.GetRows(f.GetSheetName(1))
		if err != nil {
			return nil, err
		}

		headers := rows[0]
		for _, row := range rows[1:] {
			record := make(Record)
			for i, header := range headers {
				record[header] = row[i]
			}
			records = append(records, record)
		}
	default:
		return nil, fmt.Errorf("unsupported input file format")
	}

	return records, nil
}

func writeJSON(records []Record, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	// with indentation
	encoder.SetIndent("", "  ")
	return encoder.Encode(records)
}

func writeMarkdown(records []Record, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if len(records) == 0 {
		return nil
	}

	headers := make([]string, 0, len(records[0]))
	for header := range records[0] {
		headers = append(headers, header)
	}

	file.WriteString("| " + strings.Join(headers, " | ") + " |\n")
	file.WriteString("| " + strings.Repeat("--- |", len(headers)) + "\n")

	for _, record := range records {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = record[header]
		}
		file.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	return nil
}
