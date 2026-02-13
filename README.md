# Tabula.nvim

Vim/Neovim plugin for [Tabula](https://github.com/pblazh/tabula) - a spreadsheet-inspired CSV transformation tool.

## Features

- **Syntax highlighting** for `.tbl` script files
- **Auto-execution** of Tabula scripts on CSV files
- **Fold support** for `#tabula` markers in CSV files
- **Auto-save and execute** on buffer write or leaving insert mode
- **CSV integration** with [csvview.nvim](https://github.com/hat0uma/csvview.nvim) (optional)

### Syntax Highlighting

Automatic highlighting for `.tbl` files including:

- Keywords: `let`, `fmt`
- Preprocessor directive: `#include`
- Cell references: `A1`, `B2`, `AA1`, etc.
- Cell ranges: `A1:C3`
- Comments: `//` single-line and `/* */` multi-line
- Built-in functions: `SUM`, `AVERAGE`, `IF`, `DATE`, etc.
- Operators, strings, numbers, and booleans

### Auto-execution

For CSV files with embedded Tabula scripts (marked with `#tabula:`):

- Automatically runs Tabula when you save the file
- Automatically runs Tabula when you leave insert mode
- Displays errors in the command line
- Reloads the file to show changes

## Prerequisites

- **Tabula** CLI tool must be installed and in your `$PATH`

  ```bash
  # Download from GitHub Pages
  curl -LO https://pblazh.github.io/tabula/bin/darwin/arm64/tabula  # macOS M1/M2
  chmod +x tabula
  sudo mv tabula /usr/local/bin/

  # Or build from source
  go install github.com/pblazh/tabula/cmd/cli@latest
  ```

- **Optional**: [csvview.nvim](https://github.com/hat0uma/csvview.nvim) for enhanced CSV viewing in Neovim

## Installation

### Neovim with Lazy.nvim (Recommended)

Add to your `~/.config/nvim/lua/plugins/tabula.lua`:

```lua
return {
  "pblazh/tabula",
  branch = "vim-plugin",
  dependencies = {
    "hat0uma/csvview.nvim", -- Optional but recommended
  },
  ft = { "csv", "tabula" }, -- Lazy load on filetype
}
```

Or add directly to your lazy setup:

```lua
require("lazy").setup({
  -- ... other plugins
  {
    "pblazh/tabula",
    branch = "vim-plugin",
    dependencies = { "hat0uma/csvview.nvim" },
    ft = { "csv", "tabula" },
  },
})
```

### Neovim with Packer

Add to your `~/.config/nvim/lua/plugins.lua`:

```lua
use {
  'pblazh/tabula',
  branch = 'vim-plugin',
  requires = { 'hat0uma/csvview.nvim' },
  ft = { 'csv', 'tabula' },
}
```

### Vim or Neovim with vim-plug

Add to your `~/.vimrc` or `~/.config/nvim/init.vim`:

```vim
Plug 'pblazh/tabula', {'branch': 'vim-plugin'}
Plug 'hat0uma/csvview.nvim'  " Optional
```

Then run `:PlugInstall`

### Plain Vim - Manual Installation

1. Download the plugin from the official website:
   ```bash
   # Visit https://pblazh.github.io/tabula to download
   # Or download directly:
   curl -L https://pblazh.github.io/tabula/releases/tabula.nvim-latest.tar.gz -o tabula.nvim.tar.gz
   ```

2. Extract and install:
   ```bash
   mkdir -p ~/.vim/pack/plugins/start
   tar -xzf tabula.nvim.tar.gz -C ~/.vim/pack/plugins/start/
   ```

3. Generate help tags (optional):
   ```bash
   vim -u NONE -c "helptags ~/.vim/pack/plugins/start/tabula.nvim/doc" -c q
   ```

### Neovim - Manual Installation

1. Download the plugin from the official website:
   ```bash
   # Visit https://pblazh.github.io/tabula to download
   # Or download directly:
   curl -L https://pblazh.github.io/tabula/releases/tabula.nvim-latest.tar.gz -o tabula.nvim.tar.gz
   ```

2. Extract and install:
   ```bash
   mkdir -p ~/.config/nvim/pack/plugins/start
   tar -xzf tabula.nvim.tar.gz -C ~/.config/nvim/pack/plugins/start/
   ```

3. Generate help tags (optional):
   ```bash
   nvim -u NONE -c "helptags ~/.config/nvim/pack/plugins/start/tabula.nvim/doc" -c q
   ```

## Usage

### Commands

- `:Tabula` - Manually execute Tabula on the current CSV file
- `:TabulaToggle` - Toggle auto-execution on/off for the current buffer

### Filetype Support

The plugin automatically activates for:

- **`.tbl` files** - Get syntax highlighting
- **`.csv` files** - Get auto-execution and fold support

### Working with CSV Files

1. Open a CSV file:

   ```bash
   vim data.csv
   ```

2. Add a Tabula script directive at the top:

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

4. Save the CSV file - Tabula runs automatically!

### Using Folds

The plugin sets up fold markers for Tabula sections in CSV files:

```csv
#tabula ---{{{
#tabula:#include "script.tbl"
#tabula ---}}}
Name,Age,Score
John,25,85
Jane,30,92
```

Use Vim fold commands:

- `zo` - Open fold
- `zc` - Close fold
- `za` - Toggle fold

### Working with .tbl Files

Simply create and edit `.tbl` files to get automatic syntax highlighting:

```bash
vim script.tbl
```

## Configuration

The plugin works out of the box with no configuration needed. However, you can customize behavior:

### Disable Auto-execution

To disable automatic execution, use `:TabulaToggle` or add to your config:

```vim
" In your vimrc/init.vim
let g:tabula_auto_execute = 0
```

### Custom Tabula Command

If `tabula` is not in your PATH, specify the full path:

```vim
let g:tabula_command = '/usr/local/bin/tabula'
```

(Note: This requires modifying the plugin source currently - feature request for customization)

## Troubleshooting

### "Tabula not found" error

Make sure Tabula is installed and in your PATH:

```bash
which tabula
tabula -v
```

### CSV not reloading after execution

Try manually reloading: `:e`

If autoread isn't working, add to your config:

```vim
set autoread
```

### Syntax highlighting not working

1. Check filetype detection: `:set filetype?`
   - Should show `filetype=tabula` for `.tbl` files
   - Should show `filetype=csv` for `.csv` files

2. Force reload: `:e` or `:syntax on`

### Plugin not loading

Check if the plugin loaded:

```vim
:scriptnames
```

Look for `tabula.vim` in the output.

## Documentation

For detailed help, see:

```vim
:help tabula
```

## Contributing

Issues and pull requests welcome at [https://github.com/pblazh/tabula](https://github.com/pblazh/tabula)

## License

GNU General Public License v3.0

## See Also

- [Tabula](https://github.com/pblazh/tabula) - The main Tabula CLI tool
- [Tabula Website](https://pblazh.github.io/tabula) - Download binaries and plugin
- [csvview.nvim](https://github.com/hat0uma/csvview.nvim) - Enhanced CSV viewing for Neovim
- [Tabula Documentation](https://github.com/pblazh/tabula/tree/main/doc) - Full language reference
