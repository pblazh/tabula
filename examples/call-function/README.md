# CALL Function Example

This example demonstrates how to use the CALL function to execute external programs and integrate their output into your CSV processing.

The CALL function allows you to run system commands and capture their output as strings that can be used in your CSV transformations.

## Files

- `input.csv` - Sample data with system information to gather
- `script.csvs` - Script demonstrating various CALL function uses
- `output.csv` - Expected result after running the script

## Running the Example

```bash
csvss -i input.csv -s script.csvs -a
```

## What It Does

The script demonstrates:

1. **Basic system commands** - Getting current user, date, and directory
2. **Commands with arguments** - Using formatted date output
3. **Integration with CSV data** - Using cell values as command arguments
4. **Error handling** - Dealing with commands that might fail

## Key Features Shown

- Execute shell commands from within CSVSS
- Capture command output as string values
- Pass CSV cell data as command arguments
- Combine CALL with other functions for complex transformations

## Security Notes

When using the CALL function with user-provided data, be careful about:

- Command injection vulnerabilities
- Path traversal attacks
- Validating input before passing to commands

The CALL function is powerful but should be used responsibly, especially when processing untrusted CSV data.