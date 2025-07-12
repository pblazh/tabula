# csv-spreadsheet

Command line tool for evaluating spreadsheets stored in the CSV file

## Usage

```bash
./csvss -s script.csvs input.csv > output.csv
```

### Command Line Options

- `-s <script_file>`: Path to the CSVSS script file (required)
- `-h`: Show help message
- `-i`: Update CSV file in place (not yet implemented)

### Example

Given a CSV file `data.csv`:

```csv
Name,Age,Score
Alice,25,85
Bob,30,92
```

And a script file `transform.csvs`:

```
let A1 = "Full Name";
let D1 = "Grade";
let D2 = A2 + " has score " + C2;
let D3 = A3 + " has score " + C3;
```

Run:

```bash
./csvss -s transform.csvs data.csv
```

Output:

```csv
Full Name,Age,Score,Grade
Alice,25,85,"Alice has score 85"
Bob,30,92,"Bob has score 92"
```

## Script Language Features

- **Let statements**: `let A1 = 42;` - Assign values to cells or variables
- **Format statements**: `fmt A1 = "%d";` - Set formatting for cells
- **Cell references**: Use A1, B2, etc. to reference CSV cells
- **Arithmetic**: Basic operations (+, -, \*, /)
- **Variables**: Use named variables in calculations
- **String concatenation**: Combine strings with +

## Cell Reference System

- A1 = Row 1, Column 1 (first cell)
- B2 = Row 2, Column 2
- Z26 = Row 26, Column 26
- AA1 = Row 1, Column 27 (after Z)

Cell references are zero-based internally but use 1-based notation (like Excel).

#### Math Functions

- **ABS**: Returns the absolute value of a number
- **AVERAGE**: Returns the numerical average value in a dataset
- **AVERAGEIF**: Returns the average of a range that meets criteria
- **CEILING**: Rounds a number up to the nearest integer or specified factor
- **COUNT**: Returns the count of numeric values in a dataset
- **COUNTIF**: Returns a conditional count across a range
- **FLOOR**: Rounds a number down to the nearest integer or specified factor
- **MAX**: Returns the maximum value in a numeric dataset
- **MIN**: Returns the minimum value in a numeric dataset
- **POWER**: Returns a number raised to a power
- **PRODUCT**: Returns the product of a series of numbers
- **ROUND**: Rounds a number to a certain number of decimal places
- **SQRT**: Returns the positive square root of a number
- **SUM**: Returns the sum of a series of numbers and/or cells
- **SUMIF**: Returns a conditional sum across a range
- **SUMIFS**: Returns a sum based on multiple criteria

#### Text Functions

- **CONCATENATE**: Joins several text strings into one string
- **EXACT**: Tests whether two strings are exactly the same
- **FIND**: Returns the position at which a string is first found within text
- **LEFT**: Returns the leftmost characters from a text value
- **LEN**: Returns the length of a string
- **LOWER**: Converts text to lowercase
- **MID**: Returns a substring from the middle of a text string
- **REPLACE**: Replaces part of a text string with another text string
- **RIGHT**: Returns the rightmost characters from a text value
- **SUBSTITUTE**: Replaces existing text with new text in a string
- **TRIM**: Removes extra spaces from text
- **UPPER**: Converts text to uppercase

#### Date and Time Functions

- **DATE**: Converts year, month, and day into a date
- **DATEDIF**: Calculates the difference between two dates
- **DAY**: Returns the day of a date
- **HOUR**: Returns the hour of a time value
- **MINUTE**: Returns the minute of a time value
- **MONTH**: Returns the month of a date
- **NOW**: Returns the current date and time
- **SECOND**: Returns the second of a time value
- **TIME**: Returns the decimal number for a particular time
- **TODAY**: Returns the current date
- **WEEKDAY**: Returns the day of the week corresponding to a date
- **YEAR**: Returns the year of a date

#### Logical Functions

- **AND**: Returns TRUE if all logical statements are TRUE
- **FALSE**: Returns the logical value FALSE
- **IF**: Returns one value if condition is TRUE, another if FALSE
- **IFERROR**: Returns a value if expression is an error, otherwise returns expression
- **NOT**: Returns the opposite of a logical value
- **OR**: Returns TRUE if any logical statement is TRUE
- **TRUE**: Returns the logical value TRUE

#### Lookup and Reference Functions

- **CHOOSE**: Returns an element from a list of choices based on index
- **COLUMN**: Returns the column number of a specified cell
- **HLOOKUP**: Looks up values horizontally across the top row
- **INDEX**: Returns the content of a cell at row and column intersection
- **INDIRECT**: Returns a reference specified by a text string
- **LOOKUP**: Looks up values in a vector or array
- **MATCH**: Returns the relative position of an item in a range
- **OFFSET**: Returns a reference offset from a starting cell
- **ROW**: Returns the row number of a specified cell
- **VLOOKUP**: Searches down the first column for a key and returns a value

#### Statistical Functions

- **CORREL**: Returns the correlation coefficient between two data sets
- **MEDIAN**: Returns the median of a numeric dataset
- **MODE**: Returns the most common value in a dataset
- **PERCENTILE**: Returns the k-th percentile of values in a range
- **STDEV**: Returns the standard deviation based on a sample
- **VAR**: Returns the variance based on a sample

#### Array Functions

- **FILTER**: Returns a filtered version of a source range
- **SORT**: Sorts the rows of a given array by the values in columns
- **TRANSPOSE**: Returns a transposed version of an array
- **UNIQUE**: Returns unique rows in the provided source range

### Cell Formatting

<!-- TODO: CRITICAL - Cell formatting is NOT implemented yet -->
<!-- TODO: Add support for all Google Sheets format types -->
<!-- TODO: Implement custom format string parsing and application -->

#### Standard Format Types

- **Number**: Default numerical format with customizable decimal places
- **Currency**: Displays numbers with currency symbols ($, €, ¥) and decimal places
- **Percentage**: Multiplies by 100 and adds percentage sign (%)
- **Date**: Various date formats (MM/DD/YYYY, DD/MM/YYYY, etc.)
- **Time**: 12-hour or 24-hour time formats
- **Scientific**: Scientific notation (1.23E+4)
- **Accounting**: Accounting-style number formatting
- **Financial**: Financial number formatting
- **Plain Text**: Forces numbers to be treated as text

#### Custom Format Codes

##### Basic Number Format Characters

- **0** (Zero): Forces display of digit or zero (padding with leading zeros)
- **#** (Hash): Placeholder for optional digits (no leading zeros)
- **?** (Question): Used for fraction formatting (e.g., # ?/?)
- **@** (At): Placeholder for text (preserves text as-is)
- **,** (Comma): Thousand separator
- **.** (Period): Decimal point
- **/** (Slash): Fraction separator

##### Format Structure

Custom formats can have up to four sections separated by semicolons:

```
[POSITIVE FORMAT];[NEGATIVE FORMAT];[ZERO FORMAT];[TEXT FORMAT]
```

Examples:

- `#,##0.00` - Number with thousands separator and 2 decimal places
- `$#,##0.00;($#,##0.00)` - Currency with negative values in parentheses
- `0.00%` - Percentage with 2 decimal places
- `"$"#,##0.00` - Currency with custom dollar sign

##### Date and Time Format Codes

- **m**: Month as 1-2 digits or minutes in time
- **mm**: Month as 2 digits or minutes in time
- **mmm**: Month abbreviation (Jan, Feb, etc.)
- **mmmm**: Full month name (January, February, etc.)
- **d**: Day as 1-2 digits
- **dd**: Day as 2 digits
- **ddd**: Day abbreviation (Sun, Mon, etc.)
- **dddd**: Full day name (Sunday, Monday, etc.)
- **yy**: Year as 2 digits
- **yyyy**: Year as 4 digits
- **h**: Hour as 1-2 digits (12-hour)
- **hh**: Hour as 2 digits (12-hour)
- **H**: Hour as 1-2 digits (24-hour)
- **HH**: Hour as 2 digits (24-hour)
- **s**: Seconds as 1-2 digits
- **ss**: Seconds as 2 digits
- **AM/PM**: 12-hour time indicator

##### Conditional Formatting

Format based on value conditions:

- `[<10]"Low";[>99]"High";00` - Display "Low" for values < 10, "High" for values > 99
- `[=0]"Zero";[>0]"Positive";[<0]"Negative"` - Different formats based on value

##### Text and Symbols

- Use quotes to include literal text: `"Units: "#,##0`
- Use backslash to escape special characters: `\#` displays #
- Use underscore for spacing: `_` adds space equivalent to character width

#### Format Examples

##### Number Formats

- `#,##0` - Integer with thousands separator
- `#,##0.00` - Two decimal places with thousands separator
- `0.000` - Three decimal places with leading zeros
- `#.##` - Up to two decimal places, no leading zeros

##### Currency Formats

- `"$"#,##0.00` - Dollar currency format
- `"€"#,##0.00` - Euro currency format
- `[$$-409]#,##0.00` - Localized dollar format

##### Percentage Formats

- `0%` - Percentage as integer
- `0.00%` - Percentage with two decimal places

##### Date Formats

- `MM/DD/YYYY` - US date format (01/15/2024)
- `DD/MM/YYYY` - European date format (15/01/2024)
- `MMMM DD, YYYY` - Long date format (January 15, 2024)
- `MMM DD` - Short date format (Jan 15)

##### Time Formats

- `HH:MM` - 24-hour time (14:30)
- `H:MM AM/PM` - 12-hour time (2:30 PM)
- `HH:MM:SS` - Time with seconds (14:30:45)

<https://clickup.com/blog/google-sheets-cheat-sheet/>
