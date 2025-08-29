# Function Combinations Example

This example demonstrates how to combine multiple Tabula functions to create
complex data transformations.

1. **Text Cleaning**: Combines UPPER and TRIM to clean and format names
2. **Data Validation**: Uses LEN and IF to check for empty values and provide defaults
3. **Mathematical Operations**: Combines ABS and SUM for absolute value calculations
4. **Complex Grading**: Nested IF statements with AVERAGE for grade assignment
5. **Advanced Math**: Combines POWER and CEILING for squared and rounded values

## Example Files

### Input

```csv
Name,Value1,Value2,Value3,,,,
  John  ,10,20,30,,,,
" Mary ",15,-5,25,,,,
Charlie,,35,40,,,,
#tabulafile:./script.csvs
```

### Script

```csvs
// Combining Functions Examples

// Clean and format text
let A2 = UPPER(TRIM(A2));                    // Trim then uppercase
let A3 = UPPER(TRIM(A3));
let A4 = UPPER(TRIM(A4));

// Check if not empty and provide defaults
let E1 = "Status";
let E2 = IF(LEN(A2) > 0, A2, "Empty");      // Check if not empty
let E3 = IF(LEN(A3) > 0, A3, "Empty");
let E4 = IF(LEN(A4) > 0, A4, "Empty");

// Sum of absolute values
let F1 = "Abs Sum";
let F2 = SUM(ABS(B2), ABS(C2), ABS(D2));             // Sum of absolute values
let F3 = SUM(ABS(B3), ABS(C3), ABS(D3));
let F4 = SUM(0, ABS(C4), ABS(D4));

// Complex nested conditions with math functions
let G1 = "Grade";
// Calculate average and assign grade for row 2
let avg_2 = AVERAGE(B2, C2, D2);
let G2 = IF(avg_2 > 24, "A",
         IF(avg_2 > 19, "B",
         IF(avg_2 > 14, "C", "F")));
// Calculate average and assign grade for row 3
let avg_3 = AVERAGE(B3, C3, D3);
let G3 = IF(avg_3 > 24, "A",
         IF(avg_3 > 19, "B",
         IF(avg_3 > 14, "C", "F")));
// Calculate average and assign grade for row 4
let avg_4 = AVERAGE(0, C4, D4);
let G4 = IF(avg_4 > 24, "A",
         IF(avg_4 > 19, "B",
         IF(avg_4 > 14, "C", "F")));

// Power and rounding combinations
let H1 = "Squared & Rounded";
let H2 = CEILING(POWER(B2, 2), 10);  // Square B2 and round up to nearest 10
let H3 = CEILING(POWER(B3, 2), 10);
let H4 = CEILING(POWER(0, 2), 10);
```

### Output

```csv
Name    , Value1 , Value2 , Value3 , Status  , Abs Sum , Grade , Squared & Rounded
JOHN    , 10     , 20     , 30     , JOHN    , 60      , B     , 100
MARY    , 15     , -5     , 25     , MARY    , 45      , F     , 230
CHARLIE ,        , 35     , 40     , CHARLIE , 75      , A     , 0
#tabulafile:./script.csvs
```
