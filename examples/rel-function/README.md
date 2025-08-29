# REL Function Example

This example demonstrates the REL (relative reference) function
for creating dynamic cell references.

## What REL does

REL creates relative cell references based on the target cell being assigned to:

- `REL(1, 0)` - 1 column right, same row
- `REL(-1, 0)` - 1 column left, same row
- `REL(0, -1)` - same column, 1 row up
- `REL(2, 1)` - 2 columns right, 1 row down

## Example Files

### Input

```csv
1,1,0,0,0,0
2,2,0,0,0,0
3,3,0,0,0,0
4,4,0,0,0,0
5,5,0,0,0,0
6,6,0,0,0,0
7,7,0,0,0,0
8,8,0,0,0,0
#tabulafile:./script.csvs
```

### Script

```csvs
// Various REL function examples

// Basic relative references
let A1 = REL(1, 0);              // References B1 (1 column right, same row)
let B2 = REL(-1, 0);             // References A2 (1 column left, same row)
let C3 = REL(0, -1);             // References C2 (same column, 1 row up)
let D4 = REL(2, 1);              // References F5 (2 columns right, 1 row down)

// REL in arithmetic expressions
let A5 = REL(1, 0) + REL(0, 1);           // Sum of B5 and A6
let B5 = SUM(REL(-1, 0), REL(1, 0));      // Sum of A5 and C5

// REL with conditional logic
let C5 = IF(REL(0, 1) > 10, REL(-1, 0), 0); // Conditional with relative refs

// REL with arithmetic expressions as arguments
let A6 = REL(SUM(1 , 1), 2 - 2);      // Same as REL(2, 0) - references C6
let B6 = REL(3 / 3, 4 / 2);           // Same as REL(1, 2) - references C8
```

### Output

```csv
1  , 1  , 0 , 0 , 0 , 0
2  , 2  , 0 , 0 , 0 , 0
3  , 3  , 0 , 0 , 0 , 0
4  , 4  , 0 , 0 , 0 , 0
11 , 11 , 0 , 0 , 0 , 0
0  , 0  , 0 , 0 , 0 , 0
7  , 7  , 0 , 0 , 0 , 0
8  , 8  , 0 , 0 , 0 , 0
#tabulafile:./script.csvs
```
