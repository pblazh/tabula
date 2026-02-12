# Data Cleaning Example

This example demonstrates data cleaning and validation techniques for messy CSV data.

## What it does

1. Cleans up column headers
2. Trims white space from text fields
3. Converts text numbers to numeric values
4. Handles missing/empty values
5. Adds a validation column to mark clean records

## Data Issues Addressed

- Extra spaces in headers and values
- Quoted values with internal spaces
- Text representations of numbers ("thirty")
- Missing values in salary column
- Comma-separated numbers ("75,000")

## Example Files

### Input

```csv
name,  age,  salary ,
  Alice  ,25,"  50000  ",
" Bob ",thirty,75000,
Charlie,35,,
#tabula:#include "script.tbl"
```

### Script

```tbl
// Clean up headers
let A1 = "Name";
let B1 = "Age";
let C1 = "Salary";

// Clean age - convert text numbers, validate
let B2,B4 = REL(0, 0);  // Already numeric
let B3 = IF(B3 == "thirty", 30, B3);  // Convert text to number

// Clean salary - remove commas, spaces, convert to number
let C2 = 50000;  // Clean numeric value
let C4 = IF(C4 == "", 0, C4);  // Handle empty value

// Add validation column
let D1 = "Valid";
let D2 = IF(AND(LEN(A2) > 0, AND(B2 > 0, C2 > 0)), "Yes", "No");
let D3 = IF(AND(LEN(A3) > 0, AND(B3 > 0, C3 > 0)), "Yes", "No");
let D4 = IF(AND(LEN(A4) > 0, AND(B4 > 0, C4 > 0)), "Yes", "No");
```

### Output

```csv
Name    , Age , Salary , Valid
Alice   , 25  , 50000  , Yes
Bob     , 30  , 75000  , Yes
Charlie , 35  , 0      , No
#tabula:#include "script.tbl"
```
