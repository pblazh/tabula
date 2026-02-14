# Tabula - Spreadsheet-Inspired CSV Transformation

**Transform CSV files using familiar spreadsheet formulas from the command line.**

## 🎯 What is Tabula?

Tabula is a command-line tool that brings the power of spreadsheet calculations to CSV file processing.

## 💡 Why Tabula?

### **Own Your Data**

Your data should be **yours** - not locked in proprietary formats or cloud services:

- 📁 **Store anywhere** - Local files, git repos, network drives, your choice
- 🔐 **Share when you want** - No forced cloud sync or account requirements
- 📝 **Plain text** - CSV files are readable, editable, and future-proof
- 🔍 **Use any tool** - Works with grep, sed, awk, and all text utilities
- 📊 **Universal format** - Open in Excel, Google Sheets, databases, or any CSV-compatible tool

### **No Bloated Software**

Why install a multi-gigabyte office suite when you only need basic calculations?

- 🪶 **Tiny footprint** - Single ~10MB binary, no installers, no dependencies
- 🚀 **Fast** - Processes large files efficiently from the command line
- 🌍 **Cross-platform** - macOS, Linux, Windows - works everywhere
- 💻 **Scriptable** - Integrates with shell scripts, CI/CD, and automation tools

### **Version Control Everything**

Both your data AND transformations are text files:

- ✅ **Git-friendly** - Track changes to CSV data and `.tbl` scripts
- ✅ **Diff & merge** - See exactly what changed in your data
- ✅ **Collaborate** - Share scripts and data through version control
- ✅ **Reproducible** - Exact same results every time you run a script
- ✅ **Documented** - Scripts serve as documentation for your transformations

### **Familiar & Powerful**

- 📊 **Spreadsheet syntax** - Cell references (A1, B2), functions SUM, IF, etc
- 🔢 **Rich function library** - 50+ built-in functions for numbers, text, dates, logic
- 🎯 **Purpose-built** - Designed specifically for CSV transformation, not generic programming

## 🚀 Quick Start

### Installation

Download the binary for your system:

<https://pblazh.github.io/tabula/>

Or build from source

### Hello World

**Input CSV** (`sales.csv`):

```csv
Product,Price,Quantity
Apple,1.20,10
Cherry,2.50,8
Banana,0.80,15
```

**Script** (`script.tbl`):

```tabula
// Add header for total column
let D1 = "Total";

// Calculate total for each row
let D2:D4 = REF(REL(-2,0)) * REF(REL(-1,0));

// Add a grand total row
let A5 = "TOTAL";
let D5 = SUM(D2:D4);
```

**Run**:

```bash
tabula -a -s script.tbl -i sales.csv
```

**Output**:

```csv
Product , Price , Quantity , Total
Apple   , 1.20  , 10       , 12
Cherry  , 2.50  , 8        , 20
Banana  , 0.80  , 15       , 12
TOTAL   ,       ,          , 44
```

## 🛠️ Real-World Benefits

### **Work with Standard Tools**

```bash
# Use grep to find rows
grep "Alice" data.csv | tabula -s transform.tbl

# Pipe through standard Unix tools
# calculate sales and output sorted by total
head -n1 sales.csv ; tabula -s script.tbl -i sales.csv | tail -n +2 | sort -t, -k3 -nr

# Combine with git
git diff data.csv  # See exactly what changed
git log transform.tbl  # Track transformation history
# etc
```

### **Share & Publish Freely**

Your CSV output works everywhere:

- 📊 **Import** into Excel, Google Sheets, Numbers, etc
- 🗄️ **Load** into databases (PostgreSQL, MySQL, SQLite, etc)
- 📈 **Visualize** with Tableau, Power BI, R, Python, etc
- 🌐 **Publish** to GitHub, static sites, or data portals
- 📧 **Email** as attachments without format issues

### **No Lock-in**

- ✅ Data is yours - readable in any text editor
- ✅ Scripts are portable - run anywhere
- ✅ No subscriptions, no accounts, no cloud requirements
- ✅ Works offline - no internet needed
- ✅ Free & open source - use without restrictions

## 📚 Documentation

For complete documentation, see **[doc/README.md](doc/README.md)**

## 🔌 Editor Plugins

- **[Vim/Neovim](doc/vim-plugin.md)** - Auto-execution on save, syntax highlighting for .tbl files
- **[VS Code](plugins/tabula.vscode/readme.md)** - Auto-execution on save, syntax highlighting for .tbl files

## 🎓 Learn More

- [Full Documentation](doc/README.md)
- [GitHub Repository](https://github.com/pblazh/tabula)
- [Report Issues](https://github.com/pblazh/tabula/issues)
- [Website](https://pblazh.github.io/tabula)

## 📝 License

GNU General Public License v3.0 - See [LICENSE](LICENSE) for details
