# Tabula - CSV Spreadsheet Language

## What is Tabula?

Tabula is a command-line tool that lets you transform and manipulate CSV files using a simple scripting language. Think of it as a way to apply Excel-like formulas and calculations to your CSV data from the command line.

## Why Use Tabula?

- **Automate CSV processing** - Perfect for data pipelines and batch operations
- **Familiar syntax** - Uses spreadsheet-style cell references (A1, B2, etc.)
- **Powerful functions** - Built-in mathematical, text, and logical functions
- **Flexible** - Works with files or standard input/output
- **Lightweight** - Single binary, no dependencies

## Installation

Download a binary for your system and add it to your PATH

<https://pblazh.github.io/tabula/>

## Quick Start

1. **Download** the binary for your system
2. **Create a CSV file** (data.csv):

   ```csv
   Name,Age,Score
   Alice,25,85
   Bob,30,92
   ```

3. **Create a script** (transform.tbl):

   ```
   let D1 = "Grade";
   let D2 = IF(C2 > 90, "A", "B");
   let D3 = IF(C3 > 90, "A", "B");
   ```

4. **Run the command**:

   ```bash
   tabula -s transform.tbl data.csv
   ```

5. **See the result**:

   ```csv
   Name,Age,Score,Grade
   Alice,25,85,B
   Bob,30,92,A
   ```

## How It Works

Tabula reads your CSV file and applies the transformations defined in your script. You can:

- **Reference cells** using spreadsheet notation (A1, B2, C3, etc.)
- **Assign values** to cells or variables using `let` statements
- **Use functions** like SUM, IF, CONCATENATE for calculations
- **Format output** with `fmt` statements

The tool processes your script and outputs the transformed CSV data.

## Editor Integration

### Vim/Neovim Plugin

Get syntax highlighting and auto-execution for Tabula scripts in your editor!

**Features:**

- Syntax highlighting for `.tbl` files
- Auto-execution when saving CSV files
- Built-in commands (`:Tabula`, `:TabulaToggle`)
- Fold support for organizing scripts

**See:** [plugins/vim/tabula.nvim/README.md](plugins/vim/tabula.nvim/README.md) for full installation and usage guide.

## Further Reading

- [Command Line Options](doc/command-line.md)
- [Syntax Guide](doc/syntax.md)
- [Functions](doc/functions.md)
- [Examples](doc/examples.md)
