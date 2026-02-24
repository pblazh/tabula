" Tabula plugin for Vim/Neovim
" Maintainer: Tabula Team
" Latest Revision: 2026-02-13

" Prevent loading the plugin twice
if exists('g:loaded_tabula')
  finish
endif
let g:loaded_tabula = 1

" Configuration: enable csvview integration (default: enabled)
if !exists('g:tabula_enable_csvview')
  let g:tabula_enable_csvview = 1
endif

" Configuration: enable auto-format with -a flag (default: enabled)
if !exists('g:tabula_auto_format')
  let g:tabula_auto_format = 1
endif

" Configuration: path to tabula executable (default: 'tabula')
if !exists('g:tabula_command')
  let g:tabula_command = 'tabula'
endif

" Save compatibility options
let s:save_cpo = &cpo
set cpo&vim

" Main function to execute Tabula on the current file
function! s:ExecuteTabula() abort
  " Save the file first
  silent! write!

  " Get the current file path
  let l:filepath = expand('%:p')

  " optional -a flag
  if g:tabula_auto_format
    let l:a_flg = ' -a '
  else
    let l:a_flg = ''
  endif


  " optional -m flag if file is a markdown
  if &filetype ==# 'markdown'
    let l:m_flg = ' -m '
  else
    let l:m_flg = ''
  endif

  echohl ErrorMsg
  echohm g:tabula_command . l:m_flg . l:a_flg . " -u " . shellescape(l:filepath)
  echohl None

  let l:output = system(g:tabula_command . l:m_flg . l:a_flg . " -u " . shellescape(l:filepath))

  " Check for errors
  if v:shell_error != 0
    echohl ErrorMsg
    echom 'Tabula error: ' . l:output
    echohl None
  endif

  " Reload the file to show changes
  silent! edit
endfunction

" Function to check if csvview is available and enable it
function! s:EnableCsvView() abort
  " Try to enable csvview if available (Neovim with csvview.nvim)
  if exists(':CsvViewEnable')
    CsvViewEnable
  elseif has('nvim')
    " Try to call csvview via Lua if available
    lua << EOF
      local ok, csvview = pcall(require, "csvview")
      if ok and not csvview.is_enabled(0) then
        csvview.enable(0)
      end
EOF
  endif
endfunction

" Setup autocommands for CSV and Markdown files
augroup tabula
  autocmd!

  " When a CSV or Markdown file is opened
  autocmd FileType csv,markdown call s:SetupTabula()
augroup END

" Setup function for CSV files
function! s:SetupTabula() abort
  " Don't setup if already done for this buffer
  if exists('b:tabula_is_loaded')
    return
  endif
  let b:tabula_is_loaded = 1

  " Enable folding for #tabula markers
  setlocal foldmethod=marker
  setlocal foldlevel=0

  " Enable auto-read for external changes
  setlocal autoread

  " Enable csvview if configured
  if g:tabula_enable_csvview
    call s:EnableCsvView()
  endif

  " Setup autocommands for this buffer
  augroup tabula_save
    autocmd! * <buffer>
    " Auto-execute on write
    autocmd BufWritePost <buffer> call s:ExecuteTabula()
    " Auto-execute on leaving insert mode (auto-save enabled)
    " autocmd InsertLeave <buffer> call s:ExecuteTabula()
  augroup END
endfunction

" Command to manually execute Tabula
command! Tabula call s:ExecuteTabula()

" Command to toggle auto-execution
command! TabulaToggle call s:ToggleTabula()

function! s:ToggleTabula() abort
  if exists('#tabula_save#BufWritePost')
    augroup tabula_save
      autocmd!
    augroup END
    echom 'Tabula auto-execution disabled'
  else
    augroup tabula_save
      autocmd! * <buffer>
      autocmd BufWritePost <buffer> call s:ExecuteTabula()
      autocmd InsertLeave <buffer> call s:ExecuteTabula()
    augroup END
    echom 'Tabula auto-execution enabled'
  endif
endfunction

" Restore compatibility options
let &cpo = s:save_cpo
unlet s:save_cpo
