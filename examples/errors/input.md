# Error Examples

This file demonstrates all types of errors that can occur in Tabula.

## CSV Errors

### Malformed CSV - Wrong Number of Fields

```csv
name,age,city
Alice,30,NYC
Bob,25,LA,Extra
Charlie,35
```

```tabula
let A1 = "test";
```

### CSV with Parsing Error

```csv
one,two,three
1,2,3
4,5,6
```

```tabula
let A2 = ;
```

## Table Errors

### Malformed Table Header

| name  | age | another |
| ----- | --- | ------- |
| Alice | 30  |

```tabula
let A1 = 100;
```

### Malformed Table Body

| name  | age | city |
| ----- | --- | ---- | ----- |
| Alice | 30  |
| Bob   | 25  | NYC  | Extra |

```tabula
let B2 = 200;
```

## Parsing Errors

### Unexpected Token

```csv
x,y,z
1,2,3
```

```tabula
let A1 = + ;
```

### Missing Assignment Value

```csv
a,b,c
10,20,30
```

```tabula
let B1 = ;
```

### Missing Assignment Identifier

```csv
col1,col2
100,200
```

```tabula
let = 5;
```

### Not Terminated String

```csv
data
500
```

```tabula
let A1 = "unterminated;
```

### Invalid Range with Variables

```csv
x,y
1,2
```

```tabula
A:B;
```

## Function Errors

### Wrong Arity - Too Few Arguments

```csv
n1,n2,n3
10,20,30
```

```tabula
let D1 = SUM();
```

### Wrong Arity - Too Many Arguments

```csv
val
5
```

```tabula
let B1 = NOT(true, false);
```

### Invalid Argument Type

```csv
text,num
hello,42
```

```tabula
let C1 = ROUND(A1, 2);
```

### Invalid Argument - String in Numeric Function

```csv
a,b
5,10
```

```tabula
let C1 = AVERAGE("text", B1);
```

## Operator Errors

### Type Mismatch in Addition

```csv
str,num
hello,100
```

```tabula
let C1 = A1 + B1;
```

### Type Mismatch in Comparison

```csv
text,number
world,50
```

```tabula
let C1 = A1 > B1;
```

### Unsupported Operator

```csv
x,y
1,2
```

```tabula
let C1 = A1 & B1;
```

## Reference Errors

### Undefined Variable

```csv
col
10
```

```tabula
let A1 = x;
```

## Date Format Errors

### Invalid Date Format

```csv
date
2024-13-45
```

```tabula
let B1 = DATE(A1);
```

### Unsupported Format Specifier

```csv
value
hello
```

```tabula
fmt A1 = "%v";
let x = A1;
```

### Format with No Placeholders

```csv
data
big
```

```tabula
fmt A1 = "no placeholder here";
let x = A1;
```

### Format with Multiple Placeholders

```csv
text
world
```

```tabula
fmt A1 = "%s %s";
let x = A1;
```

## Parsing Format Errors

### Cannot Parse Integer

```csv
text
not a number
```

```tabula
fmt A1 = "%d";
let x = A1
```

### Cannot Parse Boolean

```csv
empty

```

```tabula
fmt A1 = "%t";
let x = A1
```

### Cannot Parse Float

```csv
text
not_a_float
```

```tabula
fmt A1 = "%f";
let x = A1
```

## Circular Dependency

### Self Reference

```csv
a,b
1,2
```

```tabula
let A1 = A1 + 1;
```

### Circular Reference Chain

```csv
x,y,z
0,0,0
```

```tabula
let A1 = B1;
let B1 = C1;
let C1 = A1;
```

## Edge Cases

### Division by Zero (if implemented)

```csv
n
10
```

```tabula
let B1 = A1 / 0;
```

### Complex Expression Error

```csv
a,b,c
1,2,3
```

```tabula
let D1 = SUM(A1:C1) + AVERAGE("text", 5, true);
```

## Multiple Errors in One Block

```csv
one,two
10,20,30
40,50
```

```tabula
let A1 = ;
let B1 = Z99;
let C1 = SUM();
fmt D1 = 123;
let E1 = "unterminated;
```

## Combined Table and Script Errors

| num | text  |
| --- | ----- | ----- |
| 5   | hello | extra |

```tabula
let A1 = B1 + ;
let C1 = INVALID_FUNC(1, 2, 3);
```

| one | two | three |
| --- | --- | ----- |
| 11  | 20  | 191   |

<!-- Tabula: can not parse: unexpected = at tabula_a215f12c9339.md:1:7 -->

```tabula
let C2 = A2 + B2 * 9;
```

something
