# Advanced Calculations Example

This example demonstrates advanced calculations with variables,
business logic, and formatting.

## Files

- `input.csv` - Sample item data with quantities and totals
- `script.csvs` - Tabula script with tax, discount, and total calculations
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
tabula -s script.csvs input.csv
```

## Features Demonstrated

- Variable definitions for constants
- Complex business logic with conditionals
- Multi-step calculations
- Range formatting with `fmt`
- Financial calculation patterns

## Example Files

### Input

```csv
Item,Quantity,Total
Widget A,5,250.00
Widget B,3,180.00
Widget C,8,320.00
Widget D,2,150.00
Widget E,6,300.00
Widget F,4,200.00
Widget G,7,350.00
Widget H,1,75.00
Widget I,9,450.00
,,
,,
,,
,,
,,
#tabulafile:./script.csvs
```

### Script

```csvs
// Define constants
let tax_rate = 0.08;
let discount_threshold = 1000;
let discount_rate = 0.05;

// Calculate subtotal
let subtotal = SUM(C2:C10);  // Assuming quantity * price in column C

// Apply volume discount
let discount = IF(subtotal > discount_threshold,
                  subtotal * discount_rate,
                  0);

// Calculate final amounts
let discounted_total = subtotal - discount;
let tax_amount = discounted_total * tax_rate;
let final_total = discounted_total + tax_amount;

// Output summary
let A12 = "Subtotal";
let B12 = subtotal;

let A13 = "Discount";
let B13 = discount;

let A14 = "Tax";
let B14 = tax_amount;

let A15 = "Total";
let B15 = final_total;

// Format currency
fmt B12:B15 = "%.2f";
```

### Output

```csv
Item     , Quantity , Total
Widget A , 5        , 250.00
Widget B , 3        , 180.00
Widget C , 8        , 320.00
Widget D , 2        , 150.00
Widget E , 6        , 300.00
Widget F , 4        , 200.00
Widget G , 7        , 350.00
Widget H , 1        , 75.00
Widget I , 9        , 450.00
         ,          ,
Subtotal , 2275     ,
Discount , 113.75   ,
Tax      , 172.9    ,
Total    , 2334.15  ,
#tabulafile:./script.csvs
```
