# Grade Assignment Example

This example demonstrates conditional logic using nested IF statements to assign
letter grades based on numeric scores.

## What it does

1. Assigns letter grades (A, B, C, F) for each subject based on score thresholds:
   - A: 90+
   - B: 80-89
   - C: 70-79
   - F: Below 70
2. Calculates overall average and assigns an overall grade

## Example Files

### Input

```csv
Student,Math,Science,English,,,,
Alice,85,92,78,,,,
Bob,76,84,91,,,,
Carol,95,88,89,,,,
#tabula:#include "script.tbl"
```

### Script

```tbl
// Add grade columns
let E1 = "Math Grade";
let F1 = "Science Grade";
let G1 = "English Grade";
let H1 = "Overall";

// Assign letter grades based on scores
let E2 = IF(B2 > 89, "A", IF(B2 > 79, "B", IF(B2 > 69, "C", "F")));
let F2 = IF(C2 > 89, "A", IF(C2 > 79, "B", IF(C2 > 69, "C", "F")));
let G2 = IF(D2 > 89, "A", IF(D2 > 79, "B", IF(D2 > 69, "C", "F")));

let E3 = IF(B3 > 89, "A", IF(B3 > 79, "B", IF(B3 > 69, "C", "F")));
let F3 = IF(C3 > 89, "A", IF(C3 > 79, "B", IF(C3 > 69, "C", "F")));
let G3 = IF(D3 > 89, "A", IF(D3 > 79, "B", IF(D3 > 69, "C", "F")));

let E4 = IF(B4 > 89, "A", IF(B4 > 79, "B", IF(B4 > 69, "C", "F")));
let F4 = IF(C4 > 89, "A", IF(C4 > 79, "B", IF(C4 > 69, "C", "F")));
let G4 = IF(D4 > 89, "A", IF(D4 > 79, "B", IF(D4 > 69, "C", "F")));

// Calculate overall average and grade
let avg_H2 = AVERAGE(B2:D2);
let H2 = IF(avg_H2 > 89, "A",
         IF(avg_H2 > 79, "B",
         IF(avg_H2 > 69, "C", "F")));

let avg_H3 = AVERAGE(B3:D3);
let H3 = IF(avg_H3 > 89, "A",
         IF(avg_H3 > 79, "B",
         IF(avg_H3 > 69, "C", "F")));

let avg_H4 = AVERAGE(B4:D4);
let H4 = IF(avg_H4 > 89, "A",
         IF(avg_H4 > 79, "B",
         IF(avg_H4 > 69, "C", "F")));
```

### Output

```csv
Student , Math , Science , English , Math Grade , Science Grade , English Grade , Overall
Alice   , 85   , 92      , 78      , B          , A             , C             , B
Bob     , 76   , 84      , 91      , C          , B             , A             , B
Carol   , 95   , 88      , 89      , A          , B             , B             , A
#tabula:#include "script.tbl"
```
