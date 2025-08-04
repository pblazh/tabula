# Advanced Calculations Example

This example demonstrates advanced calculations with variables, business logic, and formatting.

## Files

- `input.csv` - Sample item data with quantities and totals
- `script.csvs` - CSVSS script with tax, discount, and total calculations
- `output.csv` - Expected output after applying transformations

## What it does

1. Defines business constants (tax rate, discount thresholds)
2. Calculates subtotal from item totals
3. Applies volume discounts for large orders
4. Calculates tax on discounted amount
5. Formats final amounts as currency

## Business Logic

- 8% tax rate
- 5% discount for orders over $1000
- Progressive calculation: subtotal → discount → tax → final total

## Usage

```bash
csvss -s script.csvs input.csv
```

## Features Demonstrated

- Variable definitions for constants
- Complex business logic with conditionals
- Multi-step calculations
- Range formatting with `fmt`
- Financial calculation patterns

