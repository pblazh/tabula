# Tabula Functions Reference

Tabula provides a comprehensive set of built-in functions for data
manipulation and calculations. These functions are mostly compatible with
Google Sheets, so if you need more information about function behavior,
you can refer to Google Sheets documentation.

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

### MAX

Returns the largest value among the arguments.

```tabula
MAX(number1, number2, ...)
MAX(range)
```

Example:

```tabula
let A1 = MAX(10, 25, 5);        # Result: 25
let A1 = MAX(B1:D1);            # Maximum value in range B1 to D1
```

### MAXA

Returns the largest value among the arguments, treating text as 0.

```tabula
MAXA(value1, value2, ...)
```

Example:

```tabula
let A1 = MAXA(10, "text", 25);  # Result: 25
let A1 = MAXA(B1:D1);           # Maximum value in range, text treated as 0
```

### MIN

Returns the smallest value among the arguments.

```tabula
MIN(number1, number2, ...)
MIN(range)
```

Example:

```tabula
let A1 = MIN(10, 25, 5);        # Result: 5
let A1 = MIN(B1:D1);            # Minimum value in range B1 to D1
```

### MINA

Returns the smallest value among the arguments, treating text as 0.

```tabula
MINA(value1, value2, ...)
```

Example:

```tabula
let A1 = MINA(10, "text", 25);  # Result: 0
let A1 = MINA(B1:D1);           # Minimum value in range, text treated as 0
```

### ROUND

Rounds a number to the nearest multiple of a specified significance.

```tabula
ROUND(number, significance)
```

Example:

```tabula
let A1 = ROUND(4.567, 0.01);    # Result: 4.57 (nearest 0.01)
let A1 = ROUND(4.567, 1);       # Result: 5 (nearest integer)
let A1 = ROUND(15, 10);         # Result: 20 (nearest 10)
```

### MOD

Returns the remainder after division.

```tabula
MOD(dividend, divisor)
```

Example:

```tabula
let A1 = MOD(10, 3);            # Result: 1
let A1 = MOD(15, 4);            # Result: 3
```

### SQRT

Returns the square root of a number.

```tabula
SQRT(number)
```

Example:

```tabula
let A1 = SQRT(16);              # Result: 4
let A1 = SQRT(2);               # Result: 1.414...
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

### FIND

Returns the position of one string within another string.

```tabula
FIND(within_text, find_text, start_num)
```

Example:

```tabula
let A1 = FIND("hello", "lo");       # Result: 3
let A1 = FIND("Hello World", "World"); # Result: 6
```

### LEFT

Returns the leftmost characters from a string.

```tabula
LEFT(text, num_chars)
```

Example:

```tabula
let A1 = LEFT("Hello World", 5);    # Result: "Hello"
let A1 = LEFT(B1, 3);               # First 3 characters of B1
```

### RIGHT

Returns the rightmost characters from a string.

```tabula
RIGHT(text, num_chars)
```

Example:

```tabula
let A1 = RIGHT("Hello World", 5);   # Result: "World"
let A1 = RIGHT(B1, 3);              # Last 3 characters of B1
```

### MID

Returns characters from the middle of a string.

```tabula
MID(text, start_num, num_chars)
```

Example:

```tabula
let A1 = MID("Hello World", 7, 5);  # Result: "World"
let A1 = MID(B1, 2, 3);             # 3 chars starting from position 2
```

### SUBSTITUTE

Replaces occurrences of old text with new text in a string.

```tabula
SUBSTITUTE(text, old_text, new_text, instance_num)
```

Example:

```tabula
let A1 = SUBSTITUTE("Hello World", "o", "0"); # Result: "Hell0 W0rld"
let A1 = SUBSTITUTE("Hello World", "o", "0", 1); # Result: "Hell0 World"
```

### VALUE

Converts text that represents a number to a number.

```tabula
VALUE(text)
```

Example:

```tabula
let A1 = VALUE("123");              # Result: 123
let A1 = VALUE("45.67");            # Result: 45.67
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

## Date Functions

### TODATE

Converts a string to a date using a specified format.

```tabula
TODATE(layout, value)
```

Example:

```tabula
let A1 = TODATE("2006-01-02", "2023-12-25");     # Parse date string
let A1 = TODATE("01/02/2006", "12/25/2023");     # Different format
```

### FROMDATE

Converts a date to a string using a specified format.

```tabula
FROMDATE(layout, date)
```

Example:

```tabula
let A1 = FROMDATE("2006-01-02", B1);             # Format date as string
let A1 = FROMDATE("January 2, 2006", B1);        # Custom format
```

### DAY

Returns the day of the month from a date.

```tabula
DAY(date)
```

Example:

```tabula
let A1 = DAY(TODATE("2006-01-02", "2023-12-25")); # Result: 25
```

### HOUR

Returns the hour from a date/time.

```tabula
HOUR(date)
```

Example:

```tabula
let A1 = HOUR(NOW());                            # Current hour
```

### MINUTE

Returns the minute from a date/time.

```tabula
MINUTE(date)
```

Example:

```tabula
let A1 = MINUTE(NOW());                          # Current minute
```

### MONTH

Returns the month from a date.

```tabula
MONTH(date)
```

Example:

```tabula
let A1 = MONTH(TODATE("2006-01-02", "2023-12-25")); # Result: 12
```

### SECOND

Returns the second from a date/time.

```tabula
SECOND(date)
```

Example:

```tabula
let A1 = SECOND(NOW());                          # Current second
```

### YEAR

Returns the year from a date.

```tabula
YEAR(date)
```

Example:

```tabula
let A1 = YEAR(TODATE("2006-01-02", "2023-12-25")); # Result: 2023
```

### WEEKDAY

Returns the day of the week from a date (1=Sunday, 7=Saturday).

```tabula
WEEKDAY(date)
```

Example:

```tabula
let A1 = WEEKDAY(TODATE("2006-01-02", "2023-12-25")); # Day of week
```

### NOW

Returns the current date and time.

```tabula
NOW()
```

Example:

```tabula
let A1 = NOW();                                  # Current date/time
```

### DATE

Creates a date from year, month, and day values.

```tabula
DATE(year, month, day)
```

Example:

```tabula
let A1 = DATE(2023, 12, 25);                    # Create date Dec 25, 2023
let A1 = DATE(B1, C1, D1);                      # Use cell values
```

### DATEDIF

Calculates the difference between two dates in specified units.

```tabula
DATEDIF(start_date, end_date, unit)
```

Example:

```tabula
let A1 = DATEDIF(DATE(2020,1,1), DATE(2023,1,1), "Y"); # Years difference
let A1 = DATEDIF(B1, C1, "M");                  # Months difference
```

### DAYS

Returns the number of days between two dates.

```tabula
DAYS(start_date, end_date)
```

Example:

```tabula
let A1 = DAYS(DATE(2023,1,1), DATE(2023,12,31)); # Days in 2023
let A1 = DAYS(B1, C1);                          # Days between B1 and C1
```

### DATEVALUE

Converts a date string to a date value.

```tabula
DATEVALUE(date_text)
```

Example:

```tabula
let A1 = DATEVALUE("2023-12-25");               # Convert string to date
let A1 = DATEVALUE(B1);                         # Convert B1 string to date
```

## Count Functions

### COUNT

Counts the number of numeric values in a range.

```tabula
COUNT(value1, value2, ...)
COUNT(range)
```

Example:

```tabula
let A1 = COUNT(1, 2, "text", 4);                # Result: 3
let A1 = COUNT(B1:D1);                          # Count numbers in range
```

### COUNTA

Counts the number of non-empty values in a range.

```tabula
COUNTA(value1, value2, ...)
COUNTA(range)
```

Example:

```tabula
let A1 = COUNTA(1, "", "text", 4);              # Result: 3
let A1 = COUNTA(B1:D1);                         # Count non-empty cells
```

## Information Functions

### ISNUMBER

Tests whether a value is a number.

```tabula
ISNUMBER(value)
```

Example:

```tabula
let A1 = ISNUMBER(123);                         # Result: true
let A1 = ISNUMBER("text");                      # Result: false
let A1 = ISNUMBER(B1);                          # Test B1
```

### ISTEXT

Tests whether a value is text.

```tabula
ISTEXT(value)
```

Example:

```tabula
let A1 = ISTEXT("hello");                       # Result: true
let A1 = ISTEXT(123);                           # Result: false
let A1 = ISTEXT(B1);                            # Test B1
```

### ISLOGICAL

Tests whether a value is a logical value (true/false).

```tabula
ISLOGICAL(value)
```

Example:

```tabula
let A1 = ISLOGICAL(TRUE());                     # Result: true
let A1 = ISLOGICAL("text");                     # Result: false
let A1 = ISLOGICAL(B1);                         # Test B1
```

### ISBLANK

Tests whether a value is empty or blank.

```tabula
ISBLANK(value)
```

Example:

```tabula
let A1 = ISBLANK("");                           # Result: true
let A1 = ISBLANK("text");                       # Result: false
let A1 = ISBLANK(B1);                           # Test if B1 is blank
```

## Lookup Functions

### ADDRESS

```tabula
let x = ADDRESS(3, 2);                          # Result: "B3"
```

### COL

```tabula
let x = COL("B3");                          # Result: 2
```

### ROW

```tabula
let x = COL("B3");                          # Result: 3
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

### EXEC (External Program Execution)

EXEC executes external programs and returns their stdout output as a string.

```tabula
EXEC(command, argument1, argument2, ...)
```

The first argument is the command name, and subsequent arguments are passed
as parameters to the command.

Examples:

```tabula
let A1 = EXEC("pwd");                     # Get current directory
let A1 = EXEC("echo", "Hello World");     # Simple echo command
let A1 = EXEC("date", "+%Y-%m-%d");       # Get formatted date
let A1 = EXEC("whoami");                  # Get current username
let A1 = EXEC("uname", "-s");             # Get system name
```

More complex examples:

```tabula
// Get disk usage
let A1 = EXEC("df", "-h", ".");

// Execute custom script with parameters
let A1 = EXEC("./my_script.sh", "param1", "param2");
```

Integration with CSV data:

```tabula
// Use CSV cell values as command arguments
let A2 = EXEC("echo", B1);               # Echo content of B1
let A2 = EXEC("curl", "-s", B1);         # Fetch URL from B1
let A2 = EXEC("python", "script.py", B1, C1); # Pass B1 and C1 to Python script
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
let A1 = UPPER(EXEC("whoami"));              # Get uppercase username
let A1 = LEN(EXEC("pwd"));                   # Get length of current directory path
let A1 = IF(EXEC("test", "-f", B1) == "", "File not found", "File exists"); # Check file existence
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
