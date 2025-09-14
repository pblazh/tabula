# EXEC Function Example

This example demonstrates how to use the EXEC function to execute external
programs and integrate their output into your CSV processing.

The EXEC function allows you to run system commands and capture their output
as strings that can be used in your CSV transformations. Output from commands
is processed to replace newlines with spaces for CSV compatibility.

## Security Notes

When using the EXEC function with user-provided data, be careful about:

- Command injection vulnerabilities
- Path traversal attacks
- Validating input before passing to commands

The EXEC function is powerful but should be used responsibly, especially
when processing untrusted CSV data.

## Example Files

### Input

```csv
Command, Arg, Out
echo,hello,
ls,examples/call-function/,
date,Jun 20 1975,
#tabulafile:./script.tbl
```

### Script

```tbl
let C2 = EXEC(A2, B2);
let C3 = EXEC(A3, B3);
let C4 = EXEC(A4, "-j", "-f", "%b %d %Y", B4, "+%A");
```

### Output

```csv
Command , Arg                     , Out
echo    , hello                   , hello
ls      , examples/call-function/ , input.csv output.csv README.md script.csvs
date    , Jun 20 1975             , Friday
#tabulafile:./script.tbl
```
