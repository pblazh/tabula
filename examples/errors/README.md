# Error Examples

This example demonstrates types of errors that can occur in Tabula, showing how errors are detected, reported, and displayed in markdown output.

## Overview

Tabula validates input data, parses scripts, and evaluates expressions. When errors occur, they are reported as HTML comments in the markdown output, preserving the document structure while clearly indicating what went wrong.

## Error Categories

This example covers the following error types:

### Data Errors

- **CSV Errors**: Malformed CSV files (wrong number of fields, invalid formatting)
- **Table Errors**: Malformed markdown tables (mismatched columns, invalid structure)

### Script Errors

- **Parsing Errors**: Syntax errors in Tabula scripts (unexpected tokens, missing values, unterminated strings)
- **Invalid Ranges**: Incorrectly specified cell ranges
- **Reference Errors**: Undefined variables or invalid cell references

## Usage

Run the example to see error messages:

```bash
# Process the error examples
tabula -m -a -i input.md

```

## Error Message Format

Errors appear as HTML comments in the output:

```
<!-- Tabula: [error description] at [filename]:[line]:[column] -->
```
