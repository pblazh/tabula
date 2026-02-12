# Functions Example

This example demonstrates the SUM function with various argument patterns.

## What it does

1. Sums individual cells: `SUM(A1, B1, C1, D1)`
2. Sums using ranges: `SUM(B2:D2)`
3. Sums with mixed ranges and cells: `SUM(D3:B3, A3)`

## Example Files

### Input

```csv
1,2,3,4,
1.1,2.1,3.1,4.1,
1,2,3,4,
#tabula:#include "script.tbl"
```

### Script

```tbl
let E1 = SUM(A1, B1, C1, D1);
let E2 = SUM(B2:D2);
let E3 = SUM(B3:D3, A3);
```

### Output

```csv
1   , 2   , 3   , 4   , 10
1.1 , 2.1 , 3.1 , 4.1 , 9.3
1   , 2   , 3   , 4   , 10
#tabula:#include "script.tbl"
```
