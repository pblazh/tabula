---
aliases:
tags:
created: 2026.02.16 09:20
updated: 2026.02.20 08:16
---

# Tabula

Table

| Title  | Price | amount | Total |
| ------ | ----- | ------ | ----- |
| Apples | $10   | 3      | $0    |
| Pears  | $15   | 7      | $0    |
|        |       | Total  | $0    |

```tabula
let D2 = B2 * C2;
let D3 = B3 * C3;
let D4 = SUM(D2:D3);
```

Some content between
blocks

## CSV with an included script

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
#tabula:#include "script.tbl"
```

Some content between
blocks

## CSV with an inline script

```csv
10 , 30 , 3010
d  , e  , f
```

```tabula
let A1 = 10;
let B1 = 30; let C1 = A1 + B1 * 100;
```

## CSV with an inclide script with an include

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
```

```tabula
#include "script.tbl"
```

ended
