# Tabula Functions Reference

Tabula provides a comprehensive set of built-in functions for data
manipulation and calculations.

## Numeric Functions

### SUM

Adds up numbers or ranges of cells.

```tabula
SUM(number1, number2, ...)
SUM(range)
```

Examples:

```tabula
let A1 = SUM(1, 2, 3);          # Result: 6
let A1 = SUM(B1:D1);            # Sum range B1 to D1
let A1 = SUM(B1, C2:D2, E3);    # Mix individual cells and ranges
```

### ADD

Adds exactly two numbers.

```tabula
ADD(number1, number2)
```

Example:

```tabula
let A1 = ADD(5, 3);             # Result: 8
```

### PRODUCT

Multiplies numbers together.

```tabula
PRODUCT(number1, number2, ...)
```

Example:

```tabula
let A1 = PRODUCT(2, 3, 4);      # Result: 24
let A1 = PRODUCT(B1:D1);        # Multiply range B1 to D1
```

### AVERAGE

Calculates the arithmetic mean of numbers.

```tabula
AVERAGE(number1, number2, ...)
AVERAGE(range)
```

Examples:

```tabula
let A1 = AVERAGE(10, 20, 30);   # Result: 20
let A1 = AVERAGE(B1:D1);        # Average of range B1 to D1
```

### ABS

Returns the absolute value of a number.

```tabula
ABS(number)
```

Example:

```tabula
let A1 = ABS(-5);               # Result: 5
let A1 = ABS(B1);               # Absolute value of B1
```

### POWER

Raises a number to a specified power.

```tabula
POWER(base, exponent)
```

Example:

```tabula
let A1 = POWER(2, 3);           # Result: 8 (2^3)
let A1 = POWER(B1, 2);          # Square B1
```

### CEILING

Rounds a number up to the nearest integer or specified factor.

```tabula
CEILING(number, factor)
```

Example:

```tabula
let A1 = CEILING(4.3, 1);       # Result: 5
let A1 = CEILING(15, 10);       # Result: 20
```

### FLOOR

Rounds a number down to the nearest integer or specified factor.

```tabula
FLOOR(number, factor)
```

Example:

```tabula
let A1 = FLOOR(4.7, 1);         # Result: 4
let A1 = FLOOR(15, 10);         # Result: 10
```

### INT

Converts a number to an integer by truncating decimal places.

```tabula
INT(number)
```

Example:

```tabula
let A1 = INT(4.9);              # Result: 4
let A1 = INT(-3.2);             # Result: -3
```

## String Functions

### CONCATENATE

Joins multiple strings together.

```tabula
CONCATENATE(text1, text2, ...)
```

Example:

```tabula
let A1 = CONCATENATE("Hello", " ", "World");  # Result: "Hello World"
let A1 = CONCATENATE(B1, " - ", C1);          # Join B1 and C1 with " - "
```

### LEN

Returns the length of a string.

```tabula
LEN(text)
```

Example:

```tabula
let A1 = LEN("Hello");          # Result: 5
let A1 = LEN(B1);               # Length of text in B1
```

### UPPER

Converts text to uppercase.

```tabula
UPPER(text)
```

Example:

```tabula
let A1 = UPPER("hello");        # Result: "HELLO"
let A1 = UPPER(B1);             # Convert B1 to uppercase
```

### LOWER

Converts text to lowercase.

```tabula
LOWER(text)
```

Example:

```tabula
let A1 = LOWER("HELLO");        # Result: "hello"
let A1 = LOWER(B1);             # Convert B1 to lowercase
```

### TRIM

Removes leading and trailing spaces from text.

```tabula
TRIM(text)
```

Example:

```tabula
let A1 = TRIM("  hello  ");     # Result: "hello"
let A1 = TRIM(B1);              # Remove spaces from B1
```

### EXACT

Tests whether two strings are exactly the same.

```tabula
EXACT(text1, text2)
```

Example:

```tabula
let A1 = EXACT("hello", "hello");   # Result: true
let A1 = EXACT("Hello", "hello");   # Result: false
let A1 = EXACT(B1, C1);             # Compare B1 and C1
```

## Logical Functions

### IF

Returns one value if a condition is true, another if false.

```tabula
IF(condition, value_if_true, value_if_false)
```

Examples:

```tabula
let A1 = IF(B1 > 10, "High", "Low");           # Conditional text
let A1 = IF(B1 == "", 0, B1);                  # Default value if empty
let A1 = IF(C1 > 90, "A", IF(C1 > 80, "B", "C")); # Nested conditions
```

### AND

Returns true if all conditions are true.

```tabula
AND(condition1, condition2)
```

Example:

```tabula
let A1 = AND(B1 > 0, C1 > 0);      # True if both B1 and C1 are positive
let A1 = AND(B1 == "Yes", C1 > 10); # Multiple condition types
```

### OR

Returns true if any condition is true.

```tabula
OR(condition1, condition2)
```

Example:

```tabula
let A1 = OR(B1 > 100, C1 > 100);   # True if either B1 or C1 > 100
let A1 = OR(B1 == "", C1 == "");   # True if either is empty
```

### NOT

Returns the opposite of a logical value.

```tabula
NOT(condition)
```

Example:

```tabula
let A1 = NOT(B1 > 10);              # True if B1 is NOT > 10
let A1 = NOT(B1 == "");             # True if B1 is NOT empty
```

### TRUE

Returns the logical value TRUE.

```tabula
TRUE()
```

Example:

```tabula
let A1 = TRUE();                    # Result: true
```

### FALSE

Returns the logical value FALSE.

```tabula
FALSE()
```

Example:

```tabula
let A1 = FALSE();                   # Result: false
```

## Special Functions

### REL (Relative Reference)

REL creates relative cell references based on the target cell being assigned to.

```tabula
REL(column_offset, row_offset)
```

The REL function calculates a cell reference relative to the target cell.
The offsets specify how many columns (positive = right, negative = left)
and rows (positive = down, negative = up) to move from the target cell.

Examples:

```tabula
let A1 = REL(1, 0);              # References B1 (1 column right, same row)
let B2 = REL(-1, 0);             # References A2 (1 column left, same row)
let C3 = REL(0, -1);             # References C2 (same column, 1 row up)
let D4 = REL(2, 1);              # References F5 (2 columns right, 1 row down)
```

REL can be used in arithmetic expressions and nested function calls:

```tabula
let A1 = REL(1, 0) + REL(0, 1);           # Sum of B1 and A2
let B1 = SUM(REL(-1, 0), REL(1, 0));      # Sum of A1 and C1
let C1 = IF(REL(0, 1) > 10, REL(-1, 0), 0); # Conditional with relative refs
```

REL with arithmetic expressions:

```tabula
let A1 = REL(SUM(1 , 1), 2 - 2);      # Same as REL(2, 0) - references C1
let B1 = REL(3 / 3, 4 / 2);      # Same as REL(1, 2) - references C3
```

Example of calculating running sum using REL:

```tabula
# Input CSV:
# 1,1,0
# 2,2,0
# 3,3,0
# 4,4,0

# Program:
let B1:B4 = REL(-1, 0) + REL(-2, 0);

# Result:
# 1,1,2
# 2,2,4
# 3,3,6
# 4,4,8
```

Important notes:

- REL arguments must be integers or expressions that evaluate to integers
- The resulting cell reference must be within the bounds of the input CSV
- REL cannot reference cells with negative coordinates

### CALL (External Program Execution)

CALL executes external programs and returns their stdout output as a string.

```tabula
CALL(command, argument1, argument2, ...)
```

The first argument is the command name, and subsequent arguments are passed
as parameters to the command.

Examples:

```tabula
let A1 = CALL("pwd");                     # Get current directory
let A1 = CALL("echo", "Hello World");     # Simple echo command
let A1 = CALL("date", "+%Y-%m-%d");       # Get formatted date
let A1 = CALL("whoami");                  # Get current username
let A1 = CALL("uname", "-s");             # Get system name
```

More complex examples:

```tabula
// Get disk usage
let A1 = CALL("df", "-h", ".");

// Execute custom script with parameters
let A1 = CALL("./my_script.sh", "param1", "param2");
```

Integration with CSV data:

```tabula
// Use CSV cell values as command arguments
let A2 = CALL("echo", B1);               # Echo content of B1
let A2 = CALL("curl", "-s", B1);         # Fetch URL from B1
let A2 = CALL("python", "script.py", B1, C1); # Pass B1 and C1 to Python script
```

Important notes:

- All arguments must be strings or expressions that evaluate to strings
- The command must be available in the system PATH or be an absolute path
- Output is captured from stdout and trailing whitespace is trimmed
- Newlines in the output are replaced with spaces to ensure CSV compatibility
- If the command fails, an error is returned with the failure message
- Commands run with the same environment and working directory as the Tabula process
- Be careful with security when using external commands, especially with user input

Security considerations:

- Validate input when using CSV data as command arguments
- Consider using absolute paths for scripts to avoid PATH-based attacks
- Be aware that external commands have access to the same environment as Tabula

### Combining Functions

Functions can be nested and combined:

```tabula
let A1 = UPPER(TRIM(B1));                    # Trim then uppercase
let A1 = IF(LEN(B1) > 0, B1, "Empty");      # Check if not empty
let A1 = SUM(ABS(B1), ABS(C1));             # Sum of absolute values
let A1 = UPPER(CALL("whoami"));              # Get uppercase username
let A1 = LEN(CALL("pwd"));                   # Get length of current directory path
let A1 = IF(CALL("test", "-f", B1) == "", "File not found", "File exists"); # Check file existence
```

### Using with Cell Ranges

Many functions work with cell ranges:

```tabula
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
