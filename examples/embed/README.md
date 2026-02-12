# Embedded Script Example

This example demonstrates utility bill calculations with embedded Tabula scripts directly in the CSV file using `#tabula:` comments.

## Key Features

- **Embedded Scripts**: Scripts are written directly in the CSV file using `#tabula:` comments
- **Self-contained**: No separate script files needed - everything is in one CSV file
- **Mixed Script Types**: You can use `#include` in embedded script with `#tabula:#include` file references in the same CSV file
- **Flexible Organization**: Scripts are processed in the order they appear, allowing fine-grained control over execution sequence

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
#tabula: fmt D2:D4,E2:E9 = "$%.2f"
#tabula: let E2:E4 = (REL(-2,0) - REL(-3,0)) * REL(-1,0)
#tabula: let E9 = SUM(E2:E8)
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
#tabula: fmt D2:D4,E2:E9 = "$%.2f"
#tabula: let E2:E4 = (REL(-2,0) - REL(-3,0)) * REL(-1,0)
#tabula: let E9 = SUM(E2:E8)
```
