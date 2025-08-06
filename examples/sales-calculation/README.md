# Sales Calculation Example

This example demonstrates basic runlength sum using REL function

## What it does

1. Adds a "Total" header in column D
2. Calculates the total for each product (Price × Quantity)
3. Adds a grand total row summing all individual totals

## Example Files

### Input

```csv
Product,Price,Quantity,
Apple,1.20,10,
Banana,0.80,15,
Cherry,2.50,8,
,,,
#csvss:./script.csvs
```

### Script

```csvs
// Add header for total column
let D1 = "Total";

// Calculate total for each row
let D2:D4 = REL(-2,0) * REL(-1,0);

// Add a grand total row
let A5 = "TOTAL";
let D5 = SUM(D2:D4);
```

### Output

```csv
Product , Price , Quantity , Total
Apple   , 1.20  , 10       , 12
Banana  , 0.80  , 15       , 12
Cherry  , 2.50  , 8        , 20
TOTAL   ,       ,          , 44
#csvss:./script.csvs
```
