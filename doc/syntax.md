# Tabula Syntax Guide

## Basic Syntax

Tabula uses a simple, spreadsheet-inspired syntax. Each line in your script is a statement that performs an operation.

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

Use `#include` to include other script files in your Tabula scripts. This allows you to organize your code across multiple files and reuse common definitions.

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
- If you include from a CSV file using `#tabula:#include`, paths are relative to the CSV file's directory
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

**Circular Dependency Detection**: Tabula detects and reports circular includes:

```
// a.tbl
#include "b.tbl";  // Error: circular include detected

// b.tbl
#include "a.tbl";
```

#### CSV-Embedded Includes

You can use `#include` in CSV files by prefixing with `#tabula:`:

```csv
#tabula:#include "script.tbl"
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

Use `fmt` to control how values are displayed. Format statements only accept string values or expressions that evaluate to strings.

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

Variables follow the same naming rules as programming languages:

- Start with a letter or underscore
- Can contain letters, numbers, and underscores
- Case sensitive

## Expressions and Precedence

Tabula follows standard mathematical precedence:

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

Tabula automatically analyzes dependencies and executes statements in the correct order:

```
let B1 = A1 * 2;    # This depends on A1
let A1 = 10;        # This will be executed first
```

This means you can write statements in any order, and Tabula will figure out the correct execution sequence.
