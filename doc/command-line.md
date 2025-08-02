# Command Line Reference

## Basic Usage

```bash
csvss [options] [input.csv]
```

## Options

### `-s <file>` - Script File

Specify the script file containing CSVSS commands.

```bash
csvss -s script.csvs data.csv
```

If not specified, CSVSS reads the script from standard input:

```bash
echo "let A1 = 'Hello';" | csvss data.csv
```

### `-o <file>` - Output File

Write the processed CSV to a specific file instead of standard output.

```bash
csvss -s script.csvs data.csv -o output.csv
```

### `-u` - Update In Place

Update the input CSV file directly instead of creating a new output.

```bash
csvss -s script.csvs -u data.csv
```

**Warning:** This overwrites the original file. Make backups of important data.

### `-a` - Align Output

Align CSV columns for better readability when viewing the output.

```bash
csvss -s script.csvs data.csv -a
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
csvss -s script.csvs data.csv -t
```

**When to use `-t`:**

- When your script has interdependent statements that need to execute in dependency order
- To automatically resolve statement ordering issues
- When you want CSVSS to determine the optimal execution sequence

**Example without `-t` (execution order matters):**

```csvss
let B1 = A1 + 1;  // A1 is undefined (0) at this point
let A1 = 10;      // A1 gets set after B1 calculation
```

Result: A1=10, B1=1

**Example with `-t` (dependencies resolved):**

```csvss
let B1 = A1 + 1;  // Will execute second
let A1 = 10;      // Will execute first due to sorting
```

Result: A1=10, B1=11

### `-h` - Help

Display help information and available options.

```bash
csvss -h
```

## Input Sources

### File Input

Specify the CSV file as the last argument:

```bash
csvss -s script.csvs input.csv
```

### Standard Input

If no CSV file is specified, read from standard input:

```bash
cat data.csv | csvss -s script.csvs
```

### Embedded Scripts

CSV files can contain embedded script references in comments:

```csv
Name,Age,Score
Alice,25,85
Bob,30,92
#csvss:./script.csvs
```

Then run without specifying a script file:

```bash
csvss data.csv
```

## Output Destinations

### Standard Output (Default)

By default, output goes to standard output:

```bash
csvss -s script.csvs data.csv > output.csv
```

### File Output

Use `-o` to write directly to a file:

```bash
csvss -s script.csvs data.csv -o output.csv
```

### In-Place Update

Use `-u` to modify the input file directly:

```bash
csvss -s script.csvs -u data.csv
```

## Common Usage Patterns

### Basic Processing

```bash
# Script from file, CSV from file → stdout
csvss -s transform.csvs data.csv

# Script from stdin, CSV from file → stdout
csvss data.csv < script.csvs

# Script from file, CSV from stdin → stdout
cat data.csv | csvss -s transform.csvs

# With topological sorting for dependency resolution
csvss -s complex_script.csvs data.csv -t

# Aligned output with dependency sorting
csvss -s script.csvs data.csv -t -a
```

### File Operations

```bash
# Save to new file
csvss -s script.csvs input.csv -o output.csv

# Update original file
csvss -s script.csvs -u data.csv

# Process multiple files
csvss -s script.csvs file1.csv -o processed1.csv
csvss -s script.csvs file2.csv -o processed2.csv
```

### Pipeline Integration

```bash
# Part of a data pipeline
curl -s "https://api.example.com/data.csv" | \
  csvss -s transform.csvs | \
  sort -t',' -k2 > final.csv

# Multiple processing steps
csvss -s step1.csvs data.csv | \
  csvss -s step2.csvs | \
  csvss -s step3.csvs -o result.csv
```

### Batch Processing

```bash
# Process all CSV files in directory
for file in *.csv; do
  csvss -s common_script.csvs "$file" -o "processed_$file"
done

# Update all files in place
for file in *.csv; do
  csvss -s cleanup.csvs -u "$file"
done
```

## Error Handling

### Script Errors

If there are syntax errors in your script, CSVSS will display the error location:

```bash
$ csvss -s bad_script.csvs data.csv
Error: unexpected token at line 3, column 15
```

### Runtime Errors

Errors during script execution show the context:

```bash
$ csvss -s script.csvs data.csv
Error: division by zero in expression at line 5: let A1 = B1 / C1;
```

## Exit Codes

- `0` - Success
- `1` - General error (file not found, syntax error, etc.)
- `2` - Invalid command line arguments
