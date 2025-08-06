# CALL Function Example

This example demonstrates how to use the CALL function to execute external
programs and integrate their output into your CSV processing.

The CALL function allows you to run system commands and capture their output
as strings that can be used in your CSV transformations. Output from commands
is processed to replace newlines with spaces for CSV compatibility.

## Security Notes

When using the CALL function with user-provided data, be careful about:

- Command injection vulnerabilities
- Path traversal attacks
- Validating input before passing to commands

The CALL function is powerful but should be used responsibly, especially
when processing untrusted CSV data.

## Example Files

### Input

```csv
Command, Arg, Out
echo,hello,
ls,examples/call-function/,
date,Jun 20 1975,
#csvss:./script.csvs
```

### Script

```csvs
let C2 = CALL(A2, B2);
let C3 = CALL(A3, B3);
let C4 = CALL(A4, "-j", "-f", "%b %d %Y", B4, "+%A");
```

### Output

```csv
Command , Arg                     , Out
echo    , hello                   , hello
ls      , examples/call-function/ , input.csv output.csv README.md script.csvs
date    , Jun 20 1975             , Friday
#csvss:./script.csvs
```
