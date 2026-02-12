# Command Line Reference

## Basic Usage

```bash
tabula [options] [input.csv]
```

## Options

### `-s <file>` - Script File

Specify the script file containing Tabula commands.

```bash
tabula -s script.tbl data.csv
```

If not specified, Tabula reads the script from standard input:

```bash
echo "let A1 = 'Hello';" | tabula data.csv
```

### `-o <file>` - Output File

Write the processed CSV to a specific file instead of standard output.

```bash
tabula -s script.tbl data.csv -o output.csv
```

### `-u` - Update In Place

Update the input CSV file directly instead of creating a new output.

```bash
tabula -s script.tbl -u data.csv
```

**Warning:** This overwrites the original file. Make backups of important data.

### `-a` - Align Output

Align CSV columns for better readability when viewing the output.

```bash
tabula -s script.tbl data.csv -a
```

This formats the output with consistent column spacing:

```csv
Name     , Age , Score
Alice    , 25  , 85
Bob      , 30  , 92
```

### `-t` - Topological Sort

Sort statements in the script based on their dependencies before execution. This ensures that variables are defined before they are used.

```bash
tabula -s script.tbl data.csv -t
```

**When to use `-t`:**

- When your script has interdependent statements that need to execute in dependency order
- To automatically resolve statement ordering issues
- When you want Tabula to determine the optimal execution sequence

**Example without `-t` (execution order matters):**

```tabula
let B1 = A1 + 1;  // A1 is undefined (0) at this point
let A1 = 10;      // A1 gets set after B1 calculation
```

Result: A1=10, B1=1

**Example with `-t` (dependencies resolved):**

```tabula
let B1 = A1 + 1;  // Will execute second
let A1 = 10;      // Will execute first due to sorting
```

Result: A1=10, B1=11

### `-h` - Help

Display help information and available options.

```bash
tabula -h
```

## Input Sources

### File Input

Specify the CSV file as the last argument:

```bash
tabula -s script.tbl input.csv
```

### Standard Input

If no CSV file is specified, read from standard input:

```bash
cat data.csv | tabula -s script.tbl
```

### Embedded Scripts

CSV files can contain embedded script with #include references in comments. The script path is resolved relative to the CSV file's location, not the current working directory.

```csv
Name,Age,Score
Alice,25,85
Bob,30,92
#tabula:#include "script.tbl"
```

Then run without specifying a script file:

```bash
tabula data.csv
```

## Output Destinations

### Standard Output (Default)

By default, output goes to standard output:

```bash
tabula -s script.tbl data.csv > output.csv
```

### File Output

Use `-o` to write directly to a file:

```bash
tabula -s script.tbl data.csv -o output.csv
```

### In-Place Update

Use `-u` to modify the input file directly:

```bash
tabula -s script.tbl -u data.csv
```

## Common Usage Patterns

### Basic Processing

```bash
# Script from file, CSV from file → stdout
tabula -s transform.tbl data.csv

# Script from stdin, CSV from file → stdout
tabula data.csv < script.tbl

# Script from file, CSV from stdin → stdout
cat data.csv | tabula -s transform.tbl

# With topological sorting for dependency resolution
tabula -s complex_script.tbl data.csv -t

# Aligned output with dependency sorting
tabula -s script.tbl data.csv -t -a
```

### File Operations

```bash
# Save to new file
tabula -s script.tbl input.csv -o output.csv

# Update original file
tabula -s script.tbl -u data.csv

# Process multiple files
tabula -s script.tbl file1.csv -o processed1.csv
tabula -s script.tbl file2.csv -o processed2.csv
```

### Pipeline Integration

```bash
# Part of a data pipeline
curl -s "https://api.example.com/data.csv" | \
  tabula -s transform.tbl | \
  sort -t',' -k2 > final.csv

# Multiple processing steps
tabula -s step1.tbl data.csv | \
  tabula -s step2.tbl | \
  tabula -s step3.tbl -o result.csv
```

### Batch Processing

```bash
# Process all CSV files in directory
for file in *.csv; do
  tabula -s common_script.tbl "$file" -o "processed_$file"
done

# Update all files in place
for file in *.csv; do
  tabula -s cleanup.tbl -u "$file"
done
```

## Error Handling

### Script Errors

If there are syntax errors in your script, Tabula will display the error location:

```bash
$ tabula -s bad_script.tbl data.csv
Error: unexpected token at line 3, column 15
```

### Runtime Errors

Errors during script execution show the context:

```bash
$ tabula -s script.tbl data.csv
Error: division by zero in expression at line 5: let A1 = B1 / C1;
```

## Exit Codes

- `0` - Success
- `1` - General error (file not found, syntax error, etc.)
- `2` - Invalid command line arguments
