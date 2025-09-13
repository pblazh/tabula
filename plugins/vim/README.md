# Tabula plugin for Vim

Install plugin with a plugin manager or manually

## Lazy

```lua
return {
  "pblazh/tabula/plugins/vim/tabula.nvim",
  {
    depends = {"hat0uma/csvview.nvim"},
    config = function()
      require("tabula").setup()
    end,
  },
}

```

In legacy vim just create a file `after/ftplugin/csv.vim`

```Vim
" Tabula CSV
if !exists("did_load_filetypes")
  finish
endif

if exists('b:tabula_is_loaded')
  finish
endif

let b:tabula_is_loaded = 1

" Auto-close comments like:
" #tabula ---{{{
" #tabulafile: ./script.tbl
" #tabula ---}}}
set foldmethod=marker
set foldlevel=0 "Close folds

" Enable CSV plugin of choice
:CsvViewEnable

augroup tabula
  autocmd!
  " Auto-save on exit insert mode
  au! InsertLeave *.csv call SaveAndTabula()
  " Run Tabula over a fresh saved file
  au! BufWritePost *.csv :!tabula -q -a -u %
augroup END

function SaveAndTabula()
    :w
    :!tabula -q -a -u %
endfunction
```
