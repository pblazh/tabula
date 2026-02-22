# Markdown

This example demonstrates conditional logic using nested IF statements to assign
letter grades based on numeric scores.

## What it does

1. Assigns letter grades (A, B, C, F) for each subject based on score thresholds:
   - A: 90+
   - B: 80-89
   - C: 70-79
   - F: Below 70
2. Calculates overall average and assigns an overall grade

### Data Sources

#### Markdown Table

A Markdown table defines structured data using standard table syntax.

| Title  | Price | amount | Total |
| ------ | ----- | ------ | ----- |
| Apples | $10   | 3      |       |
| Pears  | $15   | 7      |       |
|        |       | Total  |       |

### CSV Block

A CSV block defines rows of comma-separated values.

```csv
10 , 30 , 0
d , e , f
```

### Script Association

A tabula script is applied to a data block using one of the following methods:

#### 1. Inline Script Block

```csv
10 , 30 , 0
d , e , f
#tabula #include "script.tbl"
```

#### 3. External Script Block

A tabula block placed directly after a data block:

```csv
10 , 30 , 0
d , e , f
```

```tabula
...
#include "script.tbl"

let A1 = 10;
```

### Handling errors

If an error occurs during parsing or executing an output message placed after the data block as a comment.
Markdown parser ignores them, so rerunning with fixed script removes error messages

## Example Files

- [input.md]
