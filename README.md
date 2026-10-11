# CSV to JSON Lines Converter

## Overview

This Go command-line application converts California housing data from a comma-separated values (CSV) file into a newline-delimited JSON (JSONL) file. Each CSV record becomes an individual JSON object on its own line.

## Requirements

- Go installed on your system
- Git installed for version control
- The provided `housesInput.csv` dataset

## Project Structure

- `main.go`: Command-line argument validation and application entry point.
- `converter.go`: CSV parsing, column mapping, numeric validation, and JSON encoding.
- `converter_test.go`: Unit tests for column mapping and record parsing.
- `housesInput.csv`: Input dataset.
- `housesOutput.jl`: Generated JSON Lines output.
- `go.mod`: Go module definition.

## Build and Run

Initialize the project if necessary:

```bash
go mod init csv-to-jsonl
```

Format, test, and build:

```bash
gofmt -w main.go converter.go converter_test.go
go test ./...
go vet ./...
go build -o main.exe .
```

Run the application:

```bash
./main.exe housesInput.csv housesOutput.jl
```

On macOS or Linux, build with `go build -o main .` and run `./main housesInput.csv housesOutput.jl`.

## Output Format

The output is JSON Lines, not a single JSON array. Every line contains one JSON object. The program validates required columns and numeric values before encoding each record.

## Testing

Run `go test ./...` to execute the unit tests. The tests cover successful column mapping, missing required columns, valid records, and invalid numeric values.

Additional tests should cover empty files, malformed CSV records, output-file errors, and end-to-end conversion.

## Limitations

This version supports the seven housing fields required by the assignment. It is not yet a general-purpose converter for arbitrary CSV schemas. It uses floating-point values for all output fields.

## AI Assistance Disclosure

I used an AI assistant to help interpret the assignment requirements, develop an initial application structure, review Go coding practices, and propose unit tests and documentation. I reviewed the generated suggestions, ran the program and tests, and made adjustments as needed. The final implementation and understanding of the code are my responsibility.