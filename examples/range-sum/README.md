# Range Sum Example

This example demonstrates how to calculate a running sum of a column.

## Input

```csv
Date       , Expected payment , Expected sum , Real payment , Real sum , Difference
01.01.2023 , 1200.00 zł       , 0.00 zł  , 0.00 zł          , 0.00 zł  ,
20.02.2023 , 1200.00 zł       , 0.00 zł  , 2400.00 zł       , 0.00 zł  ,
01.03.2023 , 1200.00 zł       , 0.00 zł  , 0.00 zł          , 0.00 zł  ,
13.04.2023 , 1200.00 zł       , 0.00 zł  , 2400.00 zł       , 0.00 zł  ,
01.05.2023 , 1200.00 zł       , 0.00 zł  , 0.00 zł          , 0.00 zł  ,
02.06.2023 , 1200.00 zł       , 0.00 zł  , 2400.00 zł       , 0.00 zł  ,
03.07.2023 , 1200.00 zł       , 0.00 zł  , 1400.00 zł       , 0.00 zł  ,
01.08.2023 , 1200.00 zł       , 0.00 zł  , 0.00 zł          , 0.00 zł  ,
13.09.2023 , 1200.00 zł       , 0.00 zł  , 2400.00 zł       , 0.00 zł  ,
03.10.2023 , 1200.00 zł       , 0.00 zł  , 2400.00 zł       , 0.00 zł  ,
...
```

## Script

```
fmt B1:B37,C1:C,D1:D,E1:E,F1:F = "%.2f zł";       // setup format
let C2:C = SUM(REF("B2:" + REL(-1, 0)));          // calculate expected total for a period
let E2:E = SUM(REF("D2:" + REL(-1, 0)));          // calculate actual total
let F2:F = REF(REL(-1, 0)) - REF(REL(-3, 0));     // calculate a debt
```

## Open Ranges

The script uses open ranges to apply formulas and formats through the available data without hard-coding the last row.

- `B1:B37` is a regular closed range with explicit start and end cells.
- `C1:C`, `D1:D`, `E1:E`, and `F1:F` start at row 1 and continue to the last existing row in each column.
- `C2:C`, `E2:E`, and `F2:F` start at row 2 and continue to the last existing row, so the formulas are applied to every payment row.

Open ranges can also target rows or the whole remaining sheet:

```tabula
A1:1   // from A1 to the last existing cell in row 1
A1:    // from A1 to the bottom-right existing cell
:C1    // from the first cell in the last existing row back to C1
```

## Output

```csv
Date       , Expected payment , Expected sum , Real payment , Real sum    , Difference
01.01.2023 , 1200.00 zł       , 1200.00 zł   , 0.00 zł      , 0.00 zł     , -1200.00 zł
20.02.2023 , 1200.00 zł       , 2400.00 zł   , 2400.00 zł   , 2400.00 zł  , 0.00 zł
01.03.2023 , 1200.00 zł       , 3600.00 zł   , 0.00 zł      , 2400.00 zł  , -1200.00 zł
13.04.2023 , 1200.00 zł       , 4800.00 zł   , 2400.00 zł   , 4800.00 zł  , 0.00 zł
01.05.2023 , 1200.00 zł       , 6000.00 zł   , 0.00 zł      , 4800.00 zł  , -1200.00 zł
02.06.2023 , 1200.00 zł       , 7200.00 zł   , 2400.00 zł   , 7200.00 zł  , 0.00 zł
03.07.2023 , 1200.00 zł       , 8400.00 zł   , 1400.00 zł   , 8600.00 zł  , 200.00 zł
01.08.2023 , 1200.00 zł       , 9600.00 zł   , 0.00 zł      , 8600.00 zł  , -1000.00 zł
13.09.2023 , 1200.00 zł       , 10800.00 zł  , 2400.00 zł   , 11000.00 zł , 200.00 zł
03.10.2023 , 1200.00 zł       , 12000.00 zł  , 2400.00 zł   , 13400.00 zł , 1400.00 zł
```
