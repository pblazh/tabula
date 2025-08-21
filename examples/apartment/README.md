# Apartment Utilities Example

This example demonstrates utility bill calculations with monetary formatting
and relative cell references.

## What it does

1. Sets up monetary formatting for rate and payment columns
2. Calculates utility costs using the formula: (current - previous) × rate
3. Calculates the total cost by summing all payments
4. Uses REL function for relative cell references

## Example Files

### Input

```csv
2025.08.01    , previous  , current   , rate     , payment
electricity   , 21025.8   , 21200.3   , $4.32    , $0.00
water         , 82.102    , 89.519    , $56.88   , $0.00
gas           , 9791.021  , 9808.410  , $7.96    , $0.00
gas delivery  ,           ,           ,          , $122.20
water delivery,           ,           ,          , $20.00
maintenance   ,           ,           ,          , $409.49
apartment     ,           ,           ,          , $8000.00
              ,           ,           , total    , $0.00
#csvssfile:./script.csvs
```

### Script

```csvs
// Setting up a monetary format
fmt D2:D4,E2:E9 = "$%.2f"

// Calculating utility costs for metered services
let E2:E4 = (REL(-2,0) - REL(-3,0)) * REL(-1,0)

// Calculating the total cost
let E9 = SUM(E2:E8)
```

### Output

```csv
2025.08.01     , previous , current  , rate   , payment
electricity    , 21025.8  , 21200.3  , $4.32  , $753.84
water          , 82.102   , 89.519   , $56.88 , $421.88
gas            , 9791.021 , 9808.410 , $7.96  , $138.42
gas delivery   ,          ,          ,        , $122.20
water delivery ,          ,          ,        , $20.00
maintenance    ,          ,          ,        , $409.49
apartment      ,          ,          ,        , $8000.00
               ,          ,          , total  , $9865.83
#csvssfile:./script.csvs
```
