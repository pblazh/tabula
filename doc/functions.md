# Tabula Functions Reference

Tabula provides a comprehensive set of built-in functions for data
manipulation and calculations. These functions are mostly compatible with
Google Sheets, so if you need more information about function behavior,
you can refer to Google Sheets documentation.

## Numeric Functions

### SUM

Adds up numbers or ranges of cells.

```tabula
SUM(values:number...):number
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
ADD(a:number, b:number):number
```

Example:

```tabula
let A1 = ADD(5, 3);             # Result: 8
```

### PRODUCT

Multiplies numbers together.

```tabula
PRODUCT(values:number...):number
```

Example:

```tabula
let A1 = PRODUCT(2, 3, 4);      # Result: 24
let A1 = PRODUCT(B1:D1);        # Multiply range B1 to D1
```

### AVERAGE

Calculates the arithmetic mean of numbers.

```tabula
AVERAGE(values:number...):number
```

Examples:

```tabula
let A1 = AVERAGE(10, 20, 30);   # Result: 20
let A1 = AVERAGE(B1:D1);        # Average of range B1 to D1
```

### ABS

Returns the absolute value of a number.

```tabula
ABS(value:number):number
```

Example:

```tabula
let A1 = ABS(-5);               # Result: 5
let A1 = ABS(B1);               # Absolute value of B1
```

### POWER

Raises a number to a specified power.

```tabula
POWER(base:number, exponent:number):number
```

Example:

```tabula
let A1 = POWER(2, 3);           # Result: 8 (2^3)
let A1 = POWER(B1, 2);          # Square B1
```

### CEILING

Rounds a number up to the nearest integer or specified factor.

```tabula
CEILING(value:number, significance:[number]):number
```

Example:

```tabula
let A1 = CEILING(4.3, 1);       # Result: 5
let A1 = CEILING(15, 10);       # Result: 20
```

### FLOOR

Rounds a number down to the nearest integer or specified factor.

```tabula
FLOOR(value:number, significance:[number]):number
```

Example:

```tabula
let A1 = FLOOR(4.7, 1);         # Result: 4
let A1 = FLOOR(15, 10);         # Result: 10
```

### INT

Converts a number to an integer by truncating decimal places.

```tabula
INT(value:number):number
```

Example:

```tabula
let A1 = INT(4.9);              # Result: 4
let A1 = INT(-3.2);             # Result: -3
```

### MAX

Returns the largest value among the arguments.

```tabula
MAX(values:number...):number
```

Example:

```tabula
let A1 = MAX(10, 25, 5);        # Result: 25
let A1 = MAX(B1:D1);            # Maximum value in range B1 to D1
```

### MAXA

Returns the largest value among the arguments, treating text as 0.

```tabula
MAXA(values:number|string...):number
```

Example:

```tabula
let A1 = MAXA(10, "text", 25);  # Result: 25
let A1 = MAXA(B1:D1);           # Maximum value in range, text treated as 0
```

### MIN

Returns the smallest value among the arguments.

```tabula
MIN(values:number...):number
```

Example:

```tabula
let A1 = MIN(10, 25, 5);        # Result: 5
let A1 = MIN(B1:D1);            # Minimum value in range B1 to D1
```

### MINA

Returns the smallest value among the arguments, treating text as 0.

```tabula
MINA(values:number|string...):number
```

Example:

```tabula
let A1 = MINA(10, "text", 25);  # Result: 0
let A1 = MINA(B1:D1);           # Minimum value in range, text treated as 0
```

### ROUND

Rounds a number to the nearest multiple of a specified significance.

```tabula
ROUND(value:number, significance:[number]):number
```

The `significance` argument is optional and defaults to 1 (rounding to the nearest integer).

Example:

```tabula
let A1 = ROUND(4.567, 0.01);    # Result: 4.57 (nearest 0.01)
let A1 = ROUND(4.567, 1);       # Result: 5 (nearest integer)
let A1 = ROUND(15, 10);         # Result: 20 (nearest 10)
```

### MOD

Returns the remainder after division.

```tabula
MOD(dividend:number, divisor:number):number
```

Example:

```tabula
let A1 = MOD(10, 3);            # Result: 1
let A1 = MOD(15, 4);            # Result: 3
```

### SQRT

Returns the square root of a number.

```tabula
SQRT(value:number):number
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
CONCATENATE(values:string...):string
```

Example:

```tabula
let A1 = CONCATENATE("Hello", " ", "World");  # Result: "Hello World"
let A1 = CONCATENATE(B1, " - ", C1);          # Join B1 and C1 with " - "
```

### LEN

Returns the length of a string.

```tabula
LEN(value:string):number
```

Example:

```tabula
let A1 = LEN("Hello");          # Result: 5
let A1 = LEN(B1);               # Length of text in B1
```

### UPPER

Converts text to uppercase.

```tabula
UPPER(value:string):string
```

Example:

```tabula
let A1 = UPPER("hello");        # Result: "HELLO"
let A1 = UPPER(B1);             # Convert B1 to uppercase
```

### LOWER

Converts text to lowercase.

```tabula
LOWER(value:string):string
```

Example:

```tabula
let A1 = LOWER("HELLO");        # Result: "hello"
let A1 = LOWER(B1);             # Convert B1 to lowercase
```

### TRIM

Removes leading and trailing spaces from text.

```tabula
TRIM(value:string):string
```

Example:

```tabula
let A1 = TRIM("  hello  ");     # Result: "hello"
let A1 = TRIM(B1);              # Remove spaces from B1
```

### EXACT

Tests whether two strings are exactly the same.

```tabula
EXACT(a:string, b:string):boolean
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
FIND(where:string, what:string, [start:int]):number
```

Example:

```tabula
let A1 = FIND("hello", "lo");       # Result: 3
let A1 = FIND("Hello World", "World"); # Result: 6
```

### LEFT

Returns the leftmost characters from a string.

```tabula
LEFT(value:string, [amount:int]):string
```

Example:

```tabula
let A1 = LEFT("Hello World", 5);    # Result: "Hello"
let A1 = LEFT(B1, 3);               # First 3 characters of B1
```

### RIGHT

Returns the rightmost characters from a string.

```tabula
RIGHT(value:string, [amount:int]):string
```

Example:

```tabula
let A1 = RIGHT("Hello World", 5);   # Result: "World"
let A1 = RIGHT(B1, 3);              # Last 3 characters of B1
```

### MID

Returns characters from the middle of a string.

```tabula
MID(value:string, start:int, amount:int):string
```

Example:

```tabula
let A1 = MID("Hello World", 7, 5);  # Result: "World"
let A1 = MID(B1, 2, 3);             # 3 chars starting from position 2
```

### SUBSTITUTE

Replaces occurrences of old text with new text in a string.

```tabula
SUBSTITUTE(text:string, old:string, new:string, [instances:int]):string
```

Example:

```tabula
let A1 = SUBSTITUTE("Hello World", "o", "0"); # Result: "Hell0 W0rld"
let A1 = SUBSTITUTE("Hello World", "o", "0", 1); # Result: "Hell0 World"
```

### VALUE

Converts text that represents a number to a number.

```tabula
VALUE(value:string):number
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
IF(predicate:boolean, positive:any, negative:any):any
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
AND(a:boolean, b:boolean):boolean
```

Example:

```tabula
let A1 = AND(B1 > 0, C1 > 0);      # True if both B1 and C1 are positive
let A1 = AND(B1 == "Yes", C1 > 10); # Multiple condition types
```

### OR

Returns true if any condition is true.

```tabula
OR(a:boolean, b:boolean):boolean
```

Example:

```tabula
let A1 = OR(B1 > 100, C1 > 100);   # True if either B1 or C1 > 100
let A1 = OR(B1 == "", C1 == "");   # True if either is empty
```

### NOT

Returns the opposite of a logical value.

```tabula
NOT(value:boolean):boolean
```

Example:

```tabula
let A1 = NOT(B1 > 10);              # True if B1 is NOT > 10
let A1 = NOT(B1 == "");             # True if B1 is NOT empty
```

### TRUE

Returns the logical value TRUE.

```tabula
TRUE():boolean
```

Example:

```tabula
let A1 = TRUE();                    # Result: true
```

### FALSE

Returns the logical value FALSE.

```tabula
FALSE():boolean
```

Example:

```tabula
let A1 = FALSE();                   # Result: false
```

## Date Functions

### TODATE

Converts a string to a date using a specified format.

```tabula
TODATE(layout:string, value:string):date
```

Example:

```tabula
let A1 = TODATE("2006-01-02", "2023-12-25");     # Parse date string
let A1 = TODATE("01/02/2006", "12/25/2023");     # Different format
```

### FROMDATE

Converts a date to a string using a specified format.

```tabula
FROMDATE(layout:string, source:date):string
```

Example:

```tabula
let A1 = FROMDATE("2006-01-02", B1);             # Format date as string
let A1 = FROMDATE("January 2, 2006", B1);        # Custom format
```

### DAY

Returns the day of the month from a date.

```tabula
DAY(value:date):number
```

Example:

```tabula
let A1 = DAY(TODATE("2006-01-02", "2023-12-25")); # Result: 25
```

### HOUR

Returns the hour from a date/time.

```tabula
HOUR(value:date):number
```

Example:

```tabula
let A1 = HOUR(NOW());                            # Current hour
```

### MINUTE

Returns the minute from a date/time.

```tabula
MINUTE(value:date):number
```

Example:

```tabula
let A1 = MINUTE(NOW());                          # Current minute
```

### MONTH

Returns the month from a date.

```tabula
MONTH(value:date):number
```

Example:

```tabula
let A1 = MONTH(TODATE("2006-01-02", "2023-12-25")); # Result: 12
```

### SECOND

Returns the second from a date/time.

```tabula
SECOND(value:date):number
```

Example:

```tabula
let A1 = SECOND(NOW());                          # Current second
```

### YEAR

Returns the year from a date.

```tabula
YEAR(value:date):number
```

Example:

```tabula
let A1 = YEAR(TODATE("2006-01-02", "2023-12-25")); # Result: 2023
```

### WEEKDAY

Returns the day of the week from a date (1=Sunday, 7=Saturday).

```tabula
WEEKDAY(value:date):number
```

Example:

```tabula
let A1 = WEEKDAY(TODATE("2006-01-02", "2023-12-25")); # Day of week
```

### NOW

Returns the current date and time.

```tabula
NOW():date
```

Example:

```tabula
let A1 = NOW();                                  # Current date/time
```

### DATE

Creates a date from year, month, and day values.

```tabula
DATE(year:number, month:number, day:number):date
```

Example:

```tabula
let A1 = DATE(2023, 12, 25);                    # Create date Dec 25, 2023
let A1 = DATE(B1, C1, D1);                      # Use cell values
```

### DATEDIF

Calculates the difference between two dates in specified units.

```tabula
DATEDIF(start:date, end:date, unit:string):number
```

Example:

```tabula
let A1 = DATEDIF(DATE(2020,1,1), DATE(2023,1,1), "Y"); # Years difference
let A1 = DATEDIF(B1, C1, "M");                  # Months difference
```

### DAYS

Returns the number of days between two dates.

```tabula
DAYS(start:date, end:date):number
```

Example:

```tabula
let A1 = DAYS(DATE(2023,1,1), DATE(2023,12,31)); # Days in 2023
let A1 = DAYS(B1, C1);                          # Days between B1 and C1
```

### DATEVALUE

Converts a date string to a date value.

```tabula
DATEVALUE(value:string):date
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
COUNT(values:any...):number
```

Example:

```tabula
let A1 = COUNT(1, 2, "text", 4);                # Result: 3
let A1 = COUNT(B1:D1);                          # Count numbers in range
```

### COUNTA

Counts the number of non-empty values in a range.

```tabula
COUNTA(values:any...):number
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
ISNUMBER(value:any):boolean
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
ISTEXT(value:any):boolean
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
ISLOGICAL(value:any):boolean
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
ISBLANK(value:any):boolean
```

Example:

```tabula
let A1 = ISBLANK("");                           # Result: true
let A1 = ISBLANK("text");                       # Result: false
let A1 = ISBLANK(B1);                           # Test if B1 is blank
```

## Lookup Functions

### ADDRESS

Returns a cell reference as a string.

```tabula
ADDRESS(row:int, column:int):string
```

Example:

```tabula
let x = ADDRESS(3, 2);                          # Result: "B3"
```

### COLUMN

Returns the column number of a specified cell.

```tabula
COLUMN(cell:string):int
```

Example:

```tabula
let x = COLUMN("B3");                          # Result: 2
```

### ROW

Returns the row number of a specified cell.

```tabula
ROW(cell:string):int
```

Example:

```tabula
let x = ROW("B3");                          # Result: 3
```

### RANGE

Takes two cell references and returns a range

```tabula
RANGE(a:string, b:string):range
```

Example:

```tabula
let x = RANGE(B4, C5);                          # Result: [B4, C4, B5, C5]
```

## Special Functions

### REF (Cell Reference)

The `REF` function takes a cell reference as a string and returns the value of that cell. This is especially useful for dynamically accessing cell values, often in combination with the `REL` function.

```tabula
REF(cell:string):any
```

The `cell_reference` is a string representing a cell, like `"A1"`, `"B2"`, etc. The function returns the value of the specified cell, which can be of any type (number, string, boolean).

Examples:

```tabula
let A1 = REF("B1");              # A1 gets the value of B1
let C1 = REF("D2") + 10;         # C1 gets the value of D2 plus 10
```

The real power of `REF` is unleashed when used with `REL` to create dynamic, relative references. `REL` calculates a new cell address as a string, and `REF` retrieves the value from that address.

```tabula
# In cell B2, REL(-1, 0) resolves to \"A2\". REF then gets the value of A2.
let B2 = REF(REL(-1, 0));

# Sum the values of the cell to the left and the cell above.
let C3 = REF(REL(-1, 0)) + REF(REL(0, -1));
```

This allows you to define formulas that are position-independent and can be applied across a range of cells to perform calculations based on their neighbors.

### REL (Relative Reference)

REL creates relative cell references based on the target cell being assigned to.

```tabula
REL(column_offset:int, row_offset:int): string
```

The REL function calculates a cell reference relative to the target cell.
The offsets specify how many columns (positive = right, negative = left)
and rows (positive = down, negative = up) to move from the target cell.

Examples:

```tabula
let A1 = REL(1, 0);              # "B1" (1 column right, same row)
let B2 = REL(-1, 0);             # "A2" (1 column left, same row)
let C3 = REL(0, -1);             # "C2" (same column, 1 row up)
let D4 = REL(2, 1);              # "F5" (2 columns right, 1 row down)
```

REL can be used in arithmetic expressions and nested function calls when combined with REF:

```tabula
let A1 = REF(REL(1, 0)) + REF(REL(0, 1);           # Sum of B1 and A2
let B1 = SUM(REL(-1, 0)), REF(REL(1, 0)));      # Sum of A1 and C1
let C1 = IF(REF(REL(0, 1)) > 10, REF(REL(-1, 0)), 0); # Conditional with relative refs
```

REL with arithmetic expressions:

```tabula
let A1 = REL(SUM(1 , 1), 2 - 2); # Same as REL(2, 0) - references C1
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
let B1:B4 = REF(REL(-1, 0)) + REF(REL(-2, 0));

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
EXEC(command:string, args:string...):string
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
