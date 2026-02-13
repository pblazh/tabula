# Tabula for VS Code

VS Code extension for [Tabula](https://github.com/pblazh/tabula) - a spreadsheet-inspired CSV transformation tool.

## Features

- 🔄 **Auto-execution** - Automatically runs Tabula when you save CSV files
- ⚡ **Instant Updates** - See transformations applied immediately after save
- 🎛️ **Toggle Control** - Enable/disable auto-execution with a command
- 🔄 **Smart Reload** - Automatically reloads file from disk after transformation
- 🎨 **Syntax Highlighting** - Beautiful syntax coloring for `.tbl` script files
- 📝 **Language Support** - Auto-completion brackets, comments, and code folding for Tabula scripts

## Prerequisites

- **Tabula CLI** must be installed and in your `$PATH`

  ```bash
  # Download from GitHub Pages
  curl -LO https://pblazh.github.io/tabula/bin/darwin/arm64/tabula  # macOS M1/M2
  chmod +x tabula
  sudo mv tabula /usr/local/bin/

  # Or build from source
  go install github.com/pblazh/tabula/cmd/cli@latest
  ```

## Recommended Extensions

For a better CSV editing experience, we recommend installing a CSV formatting extension:

**[CSV Extension by ReprEng](https://marketplace.visualstudio.com/items?itemName=ReprEng.csv)**

This extension provides:

- 📊 **Table view** - View CSV files in a formatted table
- 🎨 **Column highlighting** - Color-coded columns for better readability
- 🔍 **Filtering & sorting** - Interactive data manipulation
- ✏️ **Cell editing** - Edit CSV data directly in table view

**To install:**

```bash
# Via command line
code --install-extension ReprEng.csv

# Or search "CSV" in VS Code Extensions marketplace
```

**Why use both?**

- **CSV Extension**: For viewing and editing CSV data in a nice table format
- **Tabula Extension**: For running transformations and scripts on CSV files

These extensions work great together! View your CSV in table mode, make changes, save, and watch Tabula automatically process it.

## Usage

### Auto-Execution on Save

1. Open a CSV file in VS Code
2. Add Tabula script directive at the top:

   ```csv
   #tabula:#include "process.tbl"
   A,B,C
   1,2,3
   4,5,6
   ```

3. Create your Tabula script (`process.tbl`):

   ```tabula
   // Calculate sum
   let D1 = "Total";
   let D2 = A2 + B2 + C2;
   let D3 = A3 + B3 + C3;
   ```

4. Save the CSV file (Ctrl+S / Cmd+S)
5. Tabula runs automatically and updates the file!

### Commands

Access commands via Command Palette (Ctrl+Shift+P / Cmd+Shift+P):

- **Tabula: Toggle Auto-Execute on Save** - Enable/disable automatic execution

### Configuration

You can configure the extension behavior in VS Code settings:

```json
{
  "tabula.autoExecute": true, // Enable/disable auto-execution on save
  "tabula.executablePath": "tabula", // Path to tabula executable
  "tabula.autoFormat": true // Enable/disable auto-format output (-a flag)
}
```

**Setting the Tabula Path:**

By default, the extension uses `tabula` from your system PATH. If you need to specify a different location:

1. Open VS Code Settings (Ctrl+, / Cmd+,)
2. Search for "tabula"
3. Set **Tabula: Executable Path** to your custom path

Examples:

- Default (uses PATH): `tabula`
- macOS/Linux: `/usr/local/bin/tabula`
- Custom location: `/Users/yourname/bin/tabula`
- Windows: `C:\Program Files\tabula\tabula.exe`

**Auto Format Option:**

The `tabula.autoFormat` setting controls the `-a` flag passed to tabula:

- **Enabled (default)**: Runs `tabula -a -u <file>` - Auto-formats the output CSV
- **Disabled**: Runs `tabula -u <file>` - No automatic formatting

This is useful if you want to control formatting manually or have custom formatting requirements.

## Syntax Highlighting for .tbl Files

The extension provides rich syntax highlighting for Tabula script files (`.tbl`):

### **Supported Elements:**

- **Keywords**: `let`, `fmt`
- **Functions**: `SUM`, `AVERAGE`, `IF`, `CONCATENATE`, etc. (50+ functions)
- **Cell References**: `A1`, `B2`, `AA10`
- **Cell Ranges**: `A1:C10`, `B2:D5`
- **Operators**: `+`, `-`, `*`, `/`, `==`, `!=`, `<`, `>`, `&&`, `||`
- **Numbers**: `42`, `3.14`
- **Strings**: `"text"`, `'text'`
- **Comments**: `// line comment`, `/* block comment */`

### **Language Features:**

- ✅ Auto-closing brackets, quotes, and parentheses
- ✅ Comment toggling (Ctrl+/ or Cmd+/)
- ✅ Block comment support (Shift+Alt+A)
- ✅ Code folding with regions
- ✅ Bracket matching

### **Example:**

```tabula
// Calculate totals with formatting
let D1 = "Total";
let D2 = B2 * C2;
let D3 = SUM(D2:D10);

// Format as currency
fmt D2:D3 = "$%.2f";

// Conditional logic
let E2 = IF(D2 > 100, "High", "Low");
```

## How It Works

1. **File Save Detection** - Extension listens for CSV file saves
2. **Execute Tabula** - Runs `tabula [-a] -u <file>` on the saved file (with optional `-a` flag based on settings)
3. **Reload File** - Updates the editor with transformed content
4. **Show Errors** - Displays any errors in VS Code notifications

## Examples

### Example 1: Calculate Grades

**data.csv:**

```csv
#tabula:#include "grades.tbl"
Name,Score
Alice,85
Bob,92
```

**grades.tbl:**

```tabula
let C1 = "Grade";
let C2 = IF(B2 >= 90, "A", IF(B2 >= 80, "B", "C"));
let C3 = IF(B3 >= 90, "A", IF(B3 >= 80, "B", "C"));
```

**Result after save:**

```csv
#tabula:#include "grades.tbl"
Name,Score,Grade
Alice,85,B
Bob,92,A
```

### Example 2: Calculate Totals

**sales.csv:**

```csv
#tabula:let D1 = "Total"
Product,Price,Quantity
Apple,1.50,10
Banana,0.80,20
```

**With inline script:**

```tabula
let D1 = "Total";
let D2 = B2 * C2;
let D3 = B3 * C3;
fmt D2:D3 = "%.2f";
```

## Troubleshooting

### "Tabula command not found"

Make sure Tabula is installed and accessible:

**Option 1: Add to PATH**

```bash
which tabula
tabula -v
```

**Option 2: Set custom path in settings**

1. Open VS Code Settings (Ctrl+, / Cmd+,)
2. Search for "tabula executable"
3. Set the full path to your tabula binary:
   - macOS/Linux: `/usr/local/bin/tabula`
   - Windows: `C:\path\to\tabula.exe`

### Auto-execution not working

1. Check if auto-execute is enabled:
   - Open Command Palette
   - Run "Tabula: Toggle Auto-Execute on Save"
   - Ensure it says "enabled"

2. Check VS Code settings:
   ```json
   {
     "tabula.autoExecute": true
   }
   ```

### Changes not appearing

Try manually reloading the file:

- Close and reopen the CSV file
- Or use "File: Revert File" command

## Development

### Building

```bash
cd plugins/tabula.vscode
npm install
npm run compile
```

## Links

- [Tabula Website](https://pblazh.github.io/tabula)
- [Tabula Documentation](https://github.com/pblazh/tabula/tree/main/doc)
- [GitHub Repository](https://github.com/pblazh/tabula)
- [Report Issues](https://github.com/pblazh/tabula/issues)

## License

GNU General Public License v3.0
