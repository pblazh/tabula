# CSVSS Functions Reference

CSVSS provides a comprehensive set of built-in functions for data manipulation and calculations.

## Numeric Functions

### SUM
Adds up numbers or ranges of cells.
```
SUM(number1, number2, ...)
SUM(range)
```
Examples:
```
let A1 = SUM(1, 2, 3);          # Result: 6
let A1 = SUM(B1:D1);            # Sum range B1 to D1
let A1 = SUM(B1, C2:D2, E3);    # Mix individual cells and ranges
```

### ADD
Adds exactly two numbers.
```
ADD(number1, number2)
```
Example:
```
let A1 = ADD(5, 3);             # Result: 8
```

### PRODUCT
Multiplies numbers together.
```
PRODUCT(number1, number2, ...)
```
Example:
```
let A1 = PRODUCT(2, 3, 4);      # Result: 24
let A1 = PRODUCT(B1:D1);        # Multiply range B1 to D1
```

### AVERAGE
Calculates the arithmetic mean of numbers.
```
AVERAGE(number1, number2, ...)
AVERAGE(range)
```
Examples:
```
let A1 = AVERAGE(10, 20, 30);   # Result: 20
let A1 = AVERAGE(B1:D1);        # Average of range B1 to D1
```

### ABS
Returns the absolute value of a number.
```
ABS(number)
```
Example:
```
let A1 = ABS(-5);               # Result: 5
let A1 = ABS(B1);               # Absolute value of B1
```

### POWER
Raises a number to a specified power.
```
POWER(base, exponent)
```
Example:
```
let A1 = POWER(2, 3);           # Result: 8 (2^3)
let A1 = POWER(B1, 2);          # Square B1
```

### CEILING
Rounds a number up to the nearest integer or specified factor.
```
CEILING(number, factor)
```
Example:
```
let A1 = CEILING(4.3, 1);       # Result: 5
let A1 = CEILING(15, 10);       # Result: 20
```

### FLOOR
Rounds a number down to the nearest integer or specified factor.
```
FLOOR(number, factor)
```
Example:
```
let A1 = FLOOR(4.7, 1);         # Result: 4
let A1 = FLOOR(15, 10);         # Result: 10
```

### INT
Converts a number to an integer by truncating decimal places.
```
INT(number)
```
Example:
```
let A1 = INT(4.9);              # Result: 4
let A1 = INT(-3.2);             # Result: -3
```

## String Functions

### CONCATENATE
Joins multiple strings together.
```
CONCATENATE(text1, text2, ...)
```
Example:
```
let A1 = CONCATENATE("Hello", " ", "World");  # Result: "Hello World"
let A1 = CONCATENATE(B1, " - ", C1);          # Join B1 and C1 with " - "
```

### LEN
Returns the length of a string.
```
LEN(text)
```
Example:
```
let A1 = LEN("Hello");          # Result: 5
let A1 = LEN(B1);               # Length of text in B1
```

### UPPER
Converts text to uppercase.
```
UPPER(text)
```
Example:
```
let A1 = UPPER("hello");        # Result: "HELLO"
let A1 = UPPER(B1);             # Convert B1 to uppercase
```

### LOWER
Converts text to lowercase.
```
LOWER(text)
```
Example:
```
let A1 = LOWER("HELLO");        # Result: "hello"
let A1 = LOWER(B1);             # Convert B1 to lowercase
```

### TRIM
Removes leading and trailing spaces from text.
```
TRIM(text)
```
Example:
```
let A1 = TRIM("  hello  ");     # Result: "hello"
let A1 = TRIM(B1);              # Remove spaces from B1
```

### EXACT
Tests whether two strings are exactly the same.
```
EXACT(text1, text2)
```
Example:
```
let A1 = EXACT("hello", "hello");   # Result: true
let A1 = EXACT("Hello", "hello");   # Result: false
let A1 = EXACT(B1, C1);             # Compare B1 and C1
```

## Logical Functions

### IF
Returns one value if a condition is true, another if false.
```
IF(condition, value_if_true, value_if_false)
```
Examples:
```
let A1 = IF(B1 > 10, "High", "Low");           # Conditional text
let A1 = IF(B1 == "", 0, B1);                  # Default value if empty
let A1 = IF(C1 > 90, "A", IF(C1 > 80, "B", "C")); # Nested conditions
```

### AND
Returns true if all conditions are true.
```
AND(condition1, condition2)
```
Example:
```
let A1 = AND(B1 > 0, C1 > 0);      # True if both B1 and C1 are positive
let A1 = AND(B1 == "Yes", C1 > 10); # Multiple condition types
```

### OR
Returns true if any condition is true.
```
OR(condition1, condition2)
```
Example:
```
let A1 = OR(B1 > 100, C1 > 100);   # True if either B1 or C1 > 100
let A1 = OR(B1 == "", C1 == "");   # True if either is empty
```

### NOT
Returns the opposite of a logical value.
```
NOT(condition)
```
Example:
```
let A1 = NOT(B1 > 10);              # True if B1 is NOT > 10
let A1 = NOT(B1 == "");             # True if B1 is NOT empty
```

### TRUE
Returns the logical value TRUE.
```
TRUE()
```
Example:
```
let A1 = TRUE();                    # Result: true
```

### FALSE
Returns the logical value FALSE.
```
FALSE()
```
Example:
```
let A1 = FALSE();                   # Result: false
```

## Function Usage Tips

### Combining Functions
Functions can be nested and combined:
```
let A1 = UPPER(TRIM(B1));                    # Trim then uppercase
let A1 = IF(LEN(B1) > 0, B1, "Empty");      # Check if not empty
let A1 = SUM(ABS(B1), ABS(C1));             # Sum of absolute values
```

### Using with Cell Ranges
Many functions work with cell ranges:
```
let total = SUM(A1:A10);            # Sum column A, rows 1-10
let avg = AVERAGE(B2:D5);           # Average of rectangular range
let result = PRODUCT(C1:C3);        # Multiply values in range
```

### Error Handling
Functions validate their arguments and return appropriate errors:
- Type mismatches (passing text to numeric functions)
- Invalid ranges
- Wrong number of arguments

Always ensure your data types match the function requirements for reliable results.