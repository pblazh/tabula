# Date Handling Example

This example demonstrates date parsing, conversion, and formatting capabilities.

## What it does

1. Parses dates from various formats automatically
2. Converts dates using the TODATE function with custom formats
3. Formats dates in different output formats (ISO and US formats)
4. Handles different date components (date only, time only, datetime)

## Functions Used

- `REL(-1, 0)`: References the cell to the left (copies date values)
- `TODATE("02/01 06 AD", A5)`: Parses a date using a custom format
- `fmt C2:C5 = "01/02/2006 03:04PM"`: Formats dates in US format

## Example Files

### Input

```csv
date             , iso , us
2005-07-24 05:08 ,     ,
2005-07-24       ,     ,
8:30PM           ,     ,
24/07 25 AD      ,     ,
#tabulafile:./script.csvs
```

### Script

```csvs
// With a default format
let B2:B5 = REL(-1, 0);

// With a custom format
fmt C2:C4 = "01/02/2006 03:04PM";
let C2:C4 = REL(-1, 0);

// Manual formating
let C5 = FROMDATE("01/02/2006 03:04PM", TODATE("02/01 06 AD", A5));

```

### Output

```csv
date             , iso                 , us
2005-07-24 05:08 , 2005-07-24 05:08:00 , 07/24/2005 05:08AM
2005-07-24       , 2005-07-24 00:00:00 , 07/24/2005 12:00AM
8:30PM           , 0000-01-01 20:30:00 , 01/01/0000 08:30PM
24/07 25 AD      , 2025-07-24 00:00:00 , 07/24/2025 12:00AM
#tabulafile:./script.csvs
```

## Key Points

- The `iso` column copies dates and displays them in ISO datetime format
- The `us` column formats the same dates in US format with AM/PM
- Time-only values are combined with a default date (0000-01-01)
- Custom date formats can be parsed using the `TODATE` function with format specifications

