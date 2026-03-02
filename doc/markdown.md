[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner2-direct.svg)](https://stand-with-ukraine.pp.ua)

# **Tabula** Markdown Support

**Tabula** can process Markdown files containing data tables and CSV blocks, making it ideal for use with note-taking applications like Obsidian, Notion, or any Markdown editor.

## Overview

Markdown mode allows you to:

- Process Markdown tables directly in your notes
- Embed CSV data blocks in Markdown files
- Attach **Tabula** scripts to data blocks
- Maintain formatted notes with live calculations
- Use Markdown files as both input and output

## Usage

Enable Markdown mode with the `-m` flag:

```bash
tabula -m -s script.tbl -i data.md
```

When using Markdown mode:

- Input: Reads Markdown file with tables and CSV blocks
- Output: Writes updated Markdown file with computed values
- Scripts: Can be embedded inline or referenced via `#include`

## Supported Data Block Types

### 1. Markdown Tables

Standard Markdown table syntax is supported:

```markdown
| Title  | Price | Amount | Total |
| ------ | ----- | ------ | ----- |
| Apples | $10   | 3      | $0.00 |
| Pears  | $15   | 7      | $0.00 |
|        |       | Total  | $0.00 |
```

**Notes:**

- Header row defines column labels (row #1)
- Separator row (`| --- |`) is required by Markdown spec
- Data starts from the row #2
- Cells can contain any text; type conversion applies automatically

### 2. CSV Code Blocks

CSV data can be embedded in code blocks with the `csv` language tag:

````markdown
```csv
10, 30, 0
20, 40, 0
```
````

**Notes:**

- Standard comma-separated values
- Whitespace around commas is trimmed
- Empty cells are represented as empty strings
- No header row assumed (all rows are data)

## Associating Scripts with Data Blocks

There are three ways to associate a **Tabula** script with a data block:

### Method 1: Tabula Code Block After Data

Place a `tabula` code block immediately after the data block:

````markdown
```csv
10 , 30 , 40
20 , 40 , 60
```

```tabula
let C1 = A1 + B1;
let C2 = A2 + B2;
```
````

Or with Markdown tables:

````markdown
| A   | B   | AB  |
| --- | --- | --- |
| 10  | 30  | 40  |
| 20  | 40  | 0   |

```tabula
let C1 = A1 + B1;
let C2 = A2 + B2;
```
````

### Method 2: Inline **Tabula** script

Add `#tabula` script inside a CSV block:

````markdown
```csv
10, 30, 0
20, 40, 0
#tabula #include "script.tbl"
```
````

**Notes:**

- Can combine inline code with includes
- Include paths are relative to the Markdown file
- Supports nested includes (included files can include other files)

## Complete Example

Here's a complete Markdown document with **Tabula** processing:
[Markdown](../examples/markdown/README.md)

**Error behavior:**

- Errors are written as HTML comments (`<!-- ... -->`)
- Markdown renderers ignore these comments
- Re-running Tabula with a fixed script removes the error message
- Original data is preserved even when errors occur

## Cell Addressing in Markdown Tables

When processing Markdown tables, **Tabula** treats them as CSV data:

**Example table:**

```markdown
| Name  | Score |
| ----- | ----- |
| Alice | 95    |
| Bob   | 87    |
```

**Cell addresses:**

- Row 1 contains headers (`Name`, `Score`) - typically not processed
- Row with separator (`-----`) - ignored
- Row 3 onwards contains data:
  - `A2` = "Alice", `B2` = 95
  - `A3` = "Bob", `B3` = 87

## Integration with Note-Taking Apps

- [Obsidian](../plugins/tabula.obsidian/README.md)
- [Visual Studio Code](../plugins/tabula.vscode/README.md)
- [Vim/NeoVim](../plugins/tabula.vim/README.md)
- [Web Storm](plugins/tabula.webstorm/README.md)

## Limitations

- **Table detection:** Only standard Markdown table syntax is supported
- **Formatting:** Complex Markdown formatting in cells may interfere with parsing
- **Performance:** Very large tables may be slow in Markdown mode

## See Also

- [Markdown Example](../examples/markdown/README.md) - Complete working example

## Links

- [Tabula Website](https://pblazh.github.io/tabula)
- [GitHub Repository](https://github.com/pblazh/tabula)
- [Report Issues](https://github.com/pblazh/tabula/issues)

## License

[GNU General Public License v3.0](./LICENSE.txt)

## Support

If you find this project useful, consider:

- ⭐ Starring the [GitHub repository](https://github.com/pblazh/tabula)
- 🐛 Reporting issues or suggesting features
- 📖 Contributing to the documentation
