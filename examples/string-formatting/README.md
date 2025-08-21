# Name Formatting Example

This example demonstrates string manipulation functions
for formatting names from a CSV file.

## What it does

1. Creates proper case full names (e.g., "john doe" → "John Doe")
2. Creates uppercase display names (e.g., "john doe" → "JOHN DOE")

## Example Files

### Input

```csv
first_name,last_name,email,
john,doe,john.doe@email.com,
jane,smith,jane.smith@email.com,
#csvssfile:./script.csvs
```

### Script

```csvs
// Create formatted full name column
let C1 = "Full Name";
let C2 = CONCATENATE(A2, " ", B2);
let C3 = CONCATENATE(A3, " ", B3);

// Create display name with uppercase
let D1 = "Display Name";
let D2 = CONCATENATE(UPPER(A2), " ", UPPER(B2));
let D3 = CONCATENATE(UPPER(A3), " ", UPPER(B3));
```

### Output

```csv
first_name , last_name , Full Name  , Display Name
john       , doe       , john doe   , JOHN DOE
jane       , smith     , jane smith , JANE SMITH
#csvssfile:./script.csvs
```
