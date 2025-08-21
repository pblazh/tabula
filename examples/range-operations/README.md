# Range Operations Example

This example demonstrates working with cell ranges for aggregation and analysis.

## What it does

1. Calculates total sales for each product across all quarters
2. Calculates quarterly totals across all products
3. Calculates quarterly averages
4. Determines best performing quarter for each product

## Example Files

### Input

```csv
Month,Q1,Q2,Q3,Q4,,
Product A,1000,1200,1100,1300,,
Product B,800,900,950,1000,,
Product C,1500,1400,1600,1550,,
,,,,,,
,,,,,,
#csvssfile:./script.csvs
```

### Script

```csvs
// Add total column
let F1 = "Total";
let F2 = SUM(B2:E2);
let F3 = SUM(B3:E3);
let F4 = SUM(B4:E4);

// Add quarterly totals row
let A5 = "Quarter Total";
let B5 = SUM(B2:B4);
let C5 = SUM(C2:C4);
let D5 = SUM(D2:D4);
let E5 = SUM(E2:E4);
let F5 = SUM(F2:F4);

// Add average row
let A6 = "Quarter Average";
let B6 = AVERAGE(B2:B4);
let C6 = AVERAGE(C2:C4);
let D6 = AVERAGE(D2:D4);
let E6 = AVERAGE(E2:E4);
let F6 = AVERAGE(F2:F4);

// Calculate best performing quarter for each product
let G1 = "Best Quarter";
let G2 = IF(B2 > C2, 
             IF(B2 > D2, IF(B2 > E2, "Q1", "Q4"), 
                         IF(D2 > E2, "Q3", "Q4")), 
             IF(C2 > D2, IF(C2 > E2, "Q2", "Q4"), 
                         IF(D2 > E2, "Q3", "Q4")));
let G3 = IF(B3 > C3, 
             IF(B3 > D3, IF(B3 > E3, "Q1", "Q4"), 
                         IF(D3 > E3, "Q3", "Q4")), 
             IF(C3 > D3, IF(C3 > E3, "Q2", "Q4"), 
                         IF(D3 > E3, "Q3", "Q4")));
let G4 = IF(B4 > C4, 
             IF(B4 > D4, IF(B4 > E4, "Q1", "Q4"), 
                         IF(D4 > E4, "Q3", "Q4")), 
             IF(C4 > D4, IF(C4 > E4, "Q2", "Q4"), 
                         IF(D4 > E4, "Q3", "Q4")));
```

### Output

```csv
Month           , Q1   , Q2   , Q3   , Q4   , Total , Best Quarter
Product A       , 1000 , 1200 , 1100 , 1300 , 4600  , Q4
Product B       , 800  , 900  , 950  , 1000 , 3650  , Q4
Product C       , 1500 , 1400 , 1600 , 1550 , 6050  , Q3
Quarter Total   , 3300 , 3500 , 3650 , 3850 , 14300 ,
Quarter Average , 1100 , 1166 , 1216 , 1283 , 4766  ,
#csvssfile:./script.csvs
```
