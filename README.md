# Excel/CSV to JSON/Markdown Converter

This repository contains a Go program that converts data from Excel or CSV files to JSON or Markdown format. 

By [ruzhila.cn](http://ruzhila.cn/?from=github_mget), a campus for learning backend development through practice.

This is a tutorial code demonstrating how to use Golang for data transformation. Pull requests are welcome. 👏


## Description

The program takes two command-line arguments: the input file and the output file. It reads the data from the input file and writes it to the output file in the specified format. The format is determined by the file extension.

- If the output file has a `.json` extension, the program writes the data as an array of JSON objects. The keys of the objects are taken from the first row of the input file, and each subsequent row is written as a separate object.

- If the output file has a `.md` extension, the program writes the data as a Markdown table. 

## Usage

```bash
go run main.go input.xlsx output.json
go run main.go input.csv output.md
```

## Dependencies

This program uses the following Go libraries:

- `encoding/csv` for reading CSV files
- `encoding/json` for writing JSON files
- `os` for file operations
- `path/filepath` for handling file paths
- `github.com/qax-os/excelize` for reading Excel files
