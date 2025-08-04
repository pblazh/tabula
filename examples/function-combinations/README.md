# Function Combinations Example

This example demonstrates how to combine multiple CSVSS functions to create complex data transformations.

## Files

- `input.csv` - Input data with names and numeric values (some messy data)
- `script.csvs` - Script showing various function combination patterns
- `output.csv` - Expected output after applying transformations

## What it does

1. **Text Cleaning**: Combines UPPER and TRIM to clean and format names
2. **Data Validation**: Uses LEN and IF to check for empty values and provide defaults
3. **Mathematical Operations**: Combines ABS and SUM for absolute value calculations
4. **Complex Grading**: Nested IF statements with AVERAGE for grade assignment
5. **Advanced Math**: Combines POWER and CEILING for squared and rounded values

## Usage

```bash
csvss -a input.csv
```

## Function Combinations Demonstrated

- `UPPER(TRIM(text))` - Clean and format text
- `IF(LEN(text) > 0, text, default)` - Provide defaults for empty values
- `SUM(ABS(val1), ABS(val2))` - Sum absolute values
- `CEILING(POWER(base, exp), factor)` - Power then round up
- Nested IF with AVERAGE for complex logic

## Features

- Chained function calls
- Complex conditional logic
- Mathematical transformations
- Data cleaning patterns
- Error-resistant formulas
