# Number Formatting Example

This example demonstrates number formatting using the `fmt` statement
for financial data presentation.

## What it does

1. Formats monetary amounts to 2 decimal places using `fmt`
2. Calculates percentage of total revenue for each item
3. Formats percentages with 1 decimal place and % symbol

## Example Files

### Input

```csv
Item,Amount,
Revenue,123456.789,
Expenses,98765.432,
Profit,24691.357,
#tabulafile:./script.tbl
```

### Script

```tbl
// Format amounts as currency with 2 decimal places
fmt B2:B4 = "%.2f";

// Format percentages
fmt C2:C4 = "%.1f%%";

// Add percentage calculations
let C1 = "Percentage";
let total_revenue = B2;
let C2 = 100.0;  // Revenue is 100%
let C3 = (B3 / total_revenue) * 100;
let C4 = (B4 / total_revenue) * 100;
```

### Output

```csv
Item     , Amount     , Percentage
Revenue  , 123456.789 , 100.0%
Expenses , 98765.432  , 80.0%
Profit   , 24691.357  , 20.0%
#tabulafile:./script.tbl
```
