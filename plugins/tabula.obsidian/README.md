[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner2-direct.svg)](https://stand-with-ukraine.pp.ua)

# Tabula for Obsidian

Obsidian plugin for [Tabula](https://github.com/pblazh/tabula) - a spreadsheet-inspired CSV transformation tool.

## Features

- 🔄 **Auto-execution** - Automatically runs Tabula when you save your markdown files
- ⚡ **Instant Updates** - See transformations applied immediately after save
- 🎛️ **Toggle Control** - Enable/disable auto-execution with a command
- 📝 **Manual Execution** - Execute Tabula on demand via command palette

## Prerequisites

- **Tabula CLI** must be installed
  See https://pblazh.github.io/tabula for details

## Installation

### From Release

1. Download the latest release from the [GitHub releases page](https://github.com/pblazh/tabula/releases)
2. Extract the files to your vault's plugins folder: `<vault>/.obsidian/plugins/tabula/`
3. Reload Obsidian
4. Enable the plugin in Settings → Community Plugins

### Manual Installation

1. Clone this repository or download the source code
2. Run `npm install` to install dependencies
3. Run `npm run build` to build the plugin
4. Copy all files from `out/` folder to your vault's plugins folder: `<vault>/.obsidian/plugins/tabula/`
   ```bash
   mkdir -p <vault>/.obsidian/plugins/tabula
   cp out/* <vault>/.obsidian/plugins/tabula/
   ```
5. Reload Obsidian
6. Enable the plugin in Settings → Community Plugins

## Usage

### Auto-Execution on Save

1. Create or open a markdown file in Obsidian
2. Add csv directive:

   ```csv
   A,B,C
   1,2,3
   4,5,6
   ```

3. Add a your Tabula script directive immediately after the csv:

   ```tabula
   // Calculate sum
   let D1 = "Total";
   let D2 = A2 + B2 + C2;
   let D3 = A3 + B3 + C3;
   ```

4. Save the CSV file (Ctrl+S / Cmd+S)
5. Tabula runs automatically and updates the file!

PS. Tabula also support include directive

```tabula
#include "caculate-sum.tbl"
```

For complete documentation visit [Tabula Website](https://pblazh.github.io/tabula)

### Commands

Access commands via Command Palette (Ctrl+P / Cmd+P):

- **Tabula: execute** - Manually run Tabula on the active markdown file
- **Tablua: toggle auto execution: Toggle auto-execution on Save** - Enable/disable automatic execution

### Settings

Configure the plugin in Settings → Tabula:

#### Auto-execute on save

Enable/disable automatic execution of Tabula scripts when saving CSV files.

- Default: `true`

#### Tabula executable path

Path to the tabula executable. Use `tabula` to use the version in your PATH, or specify an absolute path.

- Default: `tabula`
- Examples:
  - macOS/Linux: `/usr/local/bin/tabula`
  - Custom location: `/Users/yourname/bin/tabula`
  - Windows: `C:\Program Files\tabula\tabula.exe`

#### Auto format output

Enable/disable the `-a` flag passed to tabula. When enabled, tabula will automatically format the output CSV.

- Default: `true`

## How It Works

1. **File Save Detection** - Plugin listens for CSV file saves
2. **Execute Tabula** - Runs `tabula [-a] -u <file>` on the saved file (with optional `-a` flag based on settings)
3. **Reload File** - Updates the editor with transformed content

## Troubleshooting

### "Tabula command not found"

Make sure Tabula is installed and accessible:

**Option 1: Add to PATH**

```bash
which tabula
tabula -v
```

**Option 2: Set custom path in settings**

1. Open Settings → Tabula
2. Set "Tabula executable path" to the full path to your tabula binary:
   - macOS/Linux: `/usr/local/bin/tabula`
   - Windows: `C:\path\to\tabula.exe`

### Auto-execution not working

1. Check if auto-execute is enabled in Settings → Tabula
2. Verify the file has `.md` extension
3. Try manually executing via Command Palette: "Tabula: Execute Tabula on current file"

### Changes not appearing

Try manually reloading the file:

- Close and reopen the CSV file
- Or execute the command again manually

### Permission errors

On macOS/Linux, ensure the tabula executable has execute permissions:

```bash
chmod +x /path/to/tabula
```

## Development

### Building

```bash
cd plugins/tabula.obsidian
npm install
npm run build
```

The build outputs to the `out/` folder:

- `out/main.js` - Bundled plugin code
- `out/manifest.json` - Plugin metadata
- `out/styles.css` - Plugin styles

### Development Mode

```bash
npm run dev
```

This will watch for changes and rebuild automatically.

### Installation from Build

Simply copy the entire `out/` folder contents to your vault:

```bash
cp -r out/* <vault>/.obsidian/plugins/tabula/
```

## Links

- [Tabula Website](https://pblazh.github.io/tabula)
- [Tabula Documentation](https://github.com/pblazh/tabula/tree/main/doc)
- [GitHub Repository](https://github.com/pblazh/tabula)
- [Report Issues](https://github.com/pblazh/tabula/issues)

## License

[GNU General Public License v3.0](./LICENSE.txt)

## Support

If you find this plugin useful, consider:

- ⭐ Starring the [GitHub repository](https://github.com/pblazh/tabula)
- 🐛 Reporting issues or suggesting features
- 📖 Contributing to the documentation
