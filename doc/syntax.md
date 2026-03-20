[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner2-direct.svg)](https://stand-with-ukraine.pp.ua)

# **Tabula** Syntax Guide

## Basic Syntax

**Tabula** uses a simple, spreadsheet-inspired syntax. Each line in your script is a statement that performs an operation.

### Statements

Every statement ends with a semicolon (`;`).

```
let A1 = 42;
fmt B1 = "%d";
```

### Comments

Single and multyline comments are supported:

```
// comment
/* here is a comment
that spans multiple lines */
```

### Include Directives

Use `#include` to include other script files in your **Tabula** scripts. This allows you to organize your code across multiple files and reuse common definitions.

```
#include "utilities.tbl";
#include "lib/functions.tbl";
```

The semicolon after `#include` is optional:

```
#include "utilities.tbl"
```

#### File Path Resolution

Include paths are resolved relative to the file containing the `#include` directive:

- If you include from `main.tbl`, paths are relative to `main.tbl`'s directory
- If you include from a CSV file using `#tabula #include`, paths are relative to the CSV file's directory
- Subdirectories are supported: `#include "lib/utils.tbl"`
- Parent directories are supported: `#include "../shared/common.tbl"`

#### Include Features

**Duplicate Prevention**: Files are only included once, even if referenced multiple times:

```
#include "common.tbl";
#include "common.tbl";  // Ignored - already included
```

**Nested Includes**: Included files can include other files:

```
// main.tbl
#include "a.tbl";

// a.tbl
#include "b.tbl";

// b.tbl
let A1 = 42;
```

**Circular Dependency Detection**: **Tabula** detects and reports circular includes:

```
// a.tbl
#include "b.tbl";  // Error: circular include detected

// b.tbl
#include "a.tbl";
```

#### CSV-Embedded Includes

You can use `#include` in CSV files by prefixing with `#tabula`:

```csv
#tabula #include "script.tbl"
A,B,C
1,2,3
```

This allows you to reference external script files from within CSV data files.

## Cell References

Use spreadsheet-style cell references to access CSV data:

- `A1` - First row, first column
- `B2` - Second row, second column
- `Z26` - 26th row, 26th column
- `AA1` - First row, 27th column (after Z)

### Cell Range References

Reference multiple cells at once:

- `A1:C1` - Cells A1, B1, C1 (horizontal range)
- `A1:A3` - Cells A1, A2, A3 (vertical range)
- `A1:C3` - All cells from A1 to C3 (rectangular range)

## Let Statements

Use `let` to assign values to cells or variables.

### Basic Assignment

```
let A1 = 42;           # Assign number to cell A1
let B1 = "Hello";      # Assign string to cell B1
let total = 100;       # Assign to variable named 'total'
```

### Multiple Assignment

Assign the same value to multiple cells:

```
let A1, B1, C1 = 0;           # Set multiple cells
let A1:C1 = "Header";         # Set range of cells
let A1, B2:D2, E3 = 100;      # Mix individual cells and ranges
```

### Using Expressions

```
let A1 = B1 + C1;            # Add two cells
let A1 = B1 * 2;             # Multiply by constant
let A1 = "Hello " + B1;      # Concatenate strings
let A1 = SUM(B1:D1);         # Use function
```

## Format Statements

Use `fmt` to control how values are displayed in the output CSV. Format statements only accept string values or expressions that evaluate to strings.

```
fmt A1 = "%d";               # Format as integer
fmt B1 = "%.2f";             # Format as float with 2 decimals
fmt C1 = "%s";               # Format as string
```

Multiple cells can be formatted with the same format:

```
fmt A1:C1 = "%.2f";          # Format range
fmt A1, B1, C1 = "%d";       # Format multiple cells
```

### Format String Syntax

**Tabula** uses Go's `fmt.Sprintf` format specifiers. The format string determines how values are converted to text when writing the output CSV.

#### Common Format Specifiers

**Integers:**

```
fmt A1 = "%d";               # Decimal integer: 42
fmt A1 = "%5d";              # Width 5, right-aligned: "   42"
fmt A1 = "%-5d";             # Width 5, left-aligned: "42   "
fmt A1 = "%05d";             # Width 5, zero-padded: "00042"
fmt A1 = "%x";               # Hexadecimal lowercase: "2a"
fmt A1 = "%X";               # Hexadecimal uppercase: "2A"
fmt A1 = "%b";               # Binary: "101010"
```

**Floats:**

```
fmt A1 = "%f";               # Default precision (6 decimals): 3.140000
fmt A1 = "%.2f";             # 2 decimal places: 3.14
fmt A1 = "%.0f";             # No decimals (rounds): 3
fmt A1 = "%8.2f";            # Width 8, 2 decimals: "    3.14"
fmt A1 = "%e";               # Scientific notation: 3.140000e+00
fmt A1 = "%.2e";             # Scientific with precision: 3.14e+00
fmt A1 = "%g";               # Compact format (auto %f or %e)
```

**Strings:**

```
fmt A1 = "%s";               # String as-is: "hello"
fmt A1 = "%10s";             # Width 10, right-aligned: "     hello"
fmt A1 = "%-10s";            # Width 10, left-aligned: "hello     "
fmt A1 = "%.5s";             # Max 5 characters: "hello"
fmt A1 = "%q";               # Quoted string: "\"hello\""
```

**Currency and Custom:**

```
fmt A1 = "$%.2f";            # Currency: $3.14
fmt A1 = "%.2f%%";           # Percentage: 3.14%
fmt A1 = "%s units";         # With suffix: "42 units"
fmt A1 = "Item: %s";         # With prefix: "Item: Apple"
```

#### Format Application

Format statements are applied when writing the output CSV:

```
let A1 = 3.14159;
fmt A1 = "%.2f";             # A1 stored as 3.14159, written as "3.14"
```

**Important notes:**

- Format strings don't change the stored value, only the output representation
- Calculations use the stored value, not the formatted string
- Invalid format strings for the value type will cause errors

#### Practical Examples

**Financial data:**

```
fmt B2:B10 = "$%.2f";        # Currency with 2 decimals
let B2 = 1234.5;             # Output: "$1234.50"
```

**Percentages:**

```
fmt C1 = "%.1f%%";           # One decimal with % sign
let C1 = 95.5;               # Output: "95.5%"
```

**Scientific data:**

```
fmt D1 = "%.3e";             # Scientific notation, 3 decimals
let D1 = 0.00012345;         # Output: "1.235e-04"
```

**Aligned tables:**

```
fmt A1:A10 = "%8d";          # Right-align integers, width 8
fmt B1:B10 = "%10.2f";       # Right-align floats, width 10
```

## Data Types

### Numbers

```
let A1 = 42;        # Integer
let A1 = 3.14;      # Float
let A1 = -10;       # Negative number
```

### Strings

Strings must be enclosed in double quotes:

```
let A1 = "Hello World";
let A1 = "Value: 42";
let A1 = "";        # Empty string
```

### Booleans

```
let A1 = true;
let A1 = false;
```

### Dates

**Tabula** supports dates as a first-class data type using Go's `time.Time` under the hood. Date values can be created from strings, date components, or the current time.

#### Go's Date Format Approach

Date parsing and formatting in **Tabula** uses Go's reference time layout approach. Instead of using format codes like `%Y-%m-%d`, you specify the format by showing what the reference time `Mon Jan 2 15:04:05 MST 2006` would look like in your desired format.

For example, to parse ISO format dates (YYYY-MM-DD), use `"2006-01-02"` as the layout:

- `2006` represents the year (4 digits)
- `01` represents the month (2 digits)
- `02` represents the day (2 digits)

The key insight: the layout string positions define what to parse or format. The reference time itself is: `Mon Jan 2 15:04:05 MST 2006` (January 2, 2006, at 3:04:05 PM MST).

#### Creating Dates

```
let A1 = TODATE("2006-01-02", "2023-12-25");        # Parse with explicit format
let A1 = DATE(2023, 12, 25);                        # Create from year, month, day
let A1 = DATEVALUE("2023-12-25");                   # Auto-detect format
let A1 = NOW();                                     # Current date and time
```

#### Date Output

Dates are written to CSV output in ISO 8601 format (YYYY-MM-DD HH:MM:SS). Use `FROMDATE()` to control output formatting explicitly.

#### Working with Dates

See the [Date Functions](#date-functions) section in functions.md for the complete set of date manipulation functions, including parsing, formatting, extracting components, and calculating differences.

### CSV Type Conversion

When **Tabula** reads CSV files, cell values are automatically converted to the appropriate data type. This allows you to perform numeric operations on CSV data without explicit conversion.

#### Conversion Rules

**Tabula** attempts to detect the type of each CSV cell value in the following priority order:

1. **Quoted strings** - Values enclosed in double quotes are always treated as strings

   ```csv
   "42"         # String, not number
   "true"       # String, not boolean
   ```

2. **Dates** - Values matching common date formats are converted to date type

   ```csv
   2023-12-25              # ISO date
   2023-12-25 14:30:00     # ISO datetime
   25.12.2023              # European format
   12/25/2023              # US format
   ```

3. **Floats** - Numbers with decimal points

   ```csv
   3.14        # Float
   -2.5        # Negative float
   0.001       # Float
   ```

4. **Integers** - Whole numbers

   ```csv
   42          # Integer
   -10         # Negative integer
   ```

5. **Booleans** - Exact matches for "true" or "false" (case-sensitive)

   ```csv
   true        # Boolean true
   false       # Boolean false
   TRUE        # String (wrong case)
   ```

6. **Strings** - Everything else defaults to string type
   ```csv
   Hello       # String
   123abc      # String (mixed content)
   ```

#### Important Notes

- **Whitespace is trimmed** before type detection
- **Empty cells** are treated as empty strings
- Once converted, values behave according to their detected type in formulas and operations

#### Examples

Given this CSV file:

```csv
42,3.14,"100",true,2023-12-25
```

The values are interpreted as:

```
A1 = 42         # Integer
B1 = 3.14       # Float
C1 = "100"      # String (quoted)
D1 = true       # Boolean
E1 = 2023-12-25 # Date
```

This means:

```
let F1 = A1 + 10;        # 52 (numeric addition)
let F1 = C1 + 10;        # "10010" (string concatenation, not 110)
let F1 = VALUE(C1) + 10; # 110 (convert string to number first)
```

## Operators

### Arithmetic Operators

```
let A1 = B1 + C1;    # Addition
let A1 = B1 - C1;    # Subtraction
let A1 = B1 * C1;    # Multiplication
let A1 = B1 / C1;    # Division
```

### Comparison Operators

```
let A1 = B1 == C1;   # Equal
let A1 = B1 != C1;   # Not equal
let A1 = B1 > C1;    # Greater than
let A1 = B1 < C1;    # Less than
let A1 = B1 >= C1;   # Greater than or equal
let A1 = B1 <= C1;   # Less than or equal
```

### Logical Operators

```
let A1 = !B1;        # NOT (negation)
let A1 = B1 && C1;   # AND (both true)
let A1 = B1 || C1;   # OR (either true)
```

### String Concatenation

```
let A1 = B1 + C1;           # Concatenate strings
let A1 = "Name: " + B1;     # Concatenate with literal
```

## Variables

You can create variables to store intermediate values:

```
let tax_rate = 0.08;
let subtotal = A1 + B1 + C1;
let total = subtotal * (1 + tax_rate);
let D1 = total;
```

### Variable Naming Rules

Variables follow standard programming language naming conventions:

- **Must start with:** A letter (a-z, A-Z) or underscore (\_)
- **Can contain:** Letters, numbers (0-9), and underscores
- **Case sensitive:** `total`, `Total`, and `TOTAL` are different variables
- **Cannot be:** Reserved keywords (`let`, `fmt`, `true`, `false`)

**Valid variable names:**

```
let count = 10;
let total_sum = 0;
let _temp = 5;
let user1 = "Alice";
let MAX_VALUE = 100;
```

**Invalid variable names:**

```
let 1count = 10;        # Cannot start with number
let total-sum = 0;      # Cannot contain hyphens
let total sum = 0;      # Cannot contain spaces
let let = 5;            # Cannot use reserved keyword
```

### Variable Scope and Lifetime

**Global scope:**

- All variables are globally scoped across the entire script
- A variable defined anywhere can be used anywhere (after execution order is resolved)
- Variables persist for the entire script execution

**Example:**

```
let result = calculate();    # 'result' is global
let A1 = result;             # Can use 'result' here

let calculate() = B1 * 2;    # Error: functions not supported
                             # Variables only, not functions
```

### Variable Assignment and Reassignment

Variables can be assigned multiple times:

```
let total = 0;               # Initial assignment
let total = total + A1;      # Reassignment using previous value
let total = total + B1;      # Another reassignment
let C1 = total;              # Final value used
```

**Important notes:**

- The last assignment wins (if no dependencies prevent it)
- **Tabula** uses topological sorting to determine execution order based on dependencies
- Circular dependencies will cause an error

**Example with dependencies:**

```
let b = a + 1;               # Depends on 'a'
let a = 10;                  # Will execute first (dependency resolution)
let c = a + b;               # Depends on both 'a' and 'b'
# Execution order: a → b → c
```

### Variables vs Cell References

**Variables:**

- Stored in memory only
- Not written to CSV output
- Useful for intermediate calculations

**Cell references:**

- Represent cells in the CSV data
- Written to output CSV
- Persist in the final result

**Example:**

```
# Variables (not in output)
let tax_rate = 0.08;
let subtotal = A1 + B1;

# Cell reference (in output)
let C1 = subtotal * (1 + tax_rate);

# Output CSV will contain C1, but not tax_rate or subtotal
```

### Best Practices

**Use descriptive names:**

```
# Good
let tax_rate = 0.08;
let total_revenue = SUM(A1:A10);

# Less clear
let x = 0.08;
let t = SUM(A1:A10);
```

**Group related calculations:**

```
# Calculate subtotals
let subtotal = SUM(B2:B10);
let tax = subtotal * 0.08;
let shipping = 15.00;

# Calculate final total
let total = subtotal + tax + shipping;
let D11 = total;
```

**Avoid name collisions with cell references:**

```
# Potentially confusing
let A1 = 10;                 # Is this variable A1 or cell A1?

# Better
let value_a1 = 10;
let A1 = value_a1;           # Clear: variable used to set cell
```

## Expressions and Precedence

**Tabula** follows standard mathematical precedence:

1. Parentheses `()`
2. Multiplication `*` and Division `/`
3. Addition `+` and Subtraction `-`
4. Comparison operators `>`, `<`, `>=`, `<=`
5. Equality operators `==`, `!=`
6. Logical AND `&&`
7. Logical OR `||`

Examples:

```
let A1 = 2 + 3 * 4;        # Result: 14 (not 20)
let A1 = (2 + 3) * 4;      # Result: 20
let A1 = A1 > 10 && B1 < 5; # Logical expression
```

## Function Calls

Functions are called with parentheses and comma-separated arguments:

```
let A1 = SUM(B1, C1, D1);           # Sum specific cells
let A1 = SUM(B1:D1);                # Sum a range
let A1 = AVERAGE(B1:D1);            # Calculate average
let A1 = IF(B1 > 10, "High", "Low"); # Conditional value
let A1 = CONCATENATE(B1, " ", C1);  # Join strings
```

## Script Processing Order

**Tabula** automatically analyzes dependencies and executes statements in the correct order:

```
let B1 = A1 * 2;    # This depends on A1
let A1 = 10;        # This will be executed first
```

This means you can write statements in any order, and **Tabula** will figure out the correct execution sequence.

## Links

- [Tabula Website](https://pblazh.github.io/tabula)
- [Tabula Documentation](https://github.com/pblazh/tabula/tree/main/doc)
- [GitHub Repository](https://github.com/pblazh/tabula)
- [Report Issues](https://github.com/pblazh/tabula/issues)

## License

[GNU General Public License v3.0](./LICENSE.txt)

## Support

If you find this project useful, consider:

- ⭐ Starring the [GitHub repository](https://github.com/pblazh/tabula)
- 🐛 Reporting issues or suggesting features
- 📖 Contributing to the documentation
