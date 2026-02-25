" Tabula plugin for Vim/Neovim
" Maintainer: Pawel Blazejewski
" Latest Revision: 2026-02-24

" Prevent loading the plugin twice
if exists('g:loaded_tabula')
  finish
endif
let g:loaded_tabula = 1

" Save compatibility options
let s:save_cpo = &cpo
set cpo&vim

" Configuration: enable auto-format with -a flag (default: enabled)
if !exists('g:tabula_auto_format')
  let g:tabula_auto_format = 1
endif

" Configuration: path to tabula executable (default: 'tabula')
if !exists('g:tabula_command')
  let g:tabula_command = 'tabula'
endif

" Configuration: enable folding (default: enabled)
if !exists('g:tabula_enable_folding')
  let g:tabula_enable_folding = 1
endif

" Setup autocommands for CSV and Markdown files
augroup tabula
  autocmd!
  autocmd FileType csv,markdown call s:SetupTabula()
augroup END

" Async handlers (Neovim)
function! s:TabulaOnStdout(job_id, data, event) abort
  if !exists('b:tabula_stdout')
    let b:tabula_stdout = []
  endif
  call extend(b:tabula_stdout, a:data)
endfunction

function! s:TabulaOnStderr(job_id, data, event) abort
  if !exists('b:tabula_stderr')
    let b:tabula_stderr = []
  endif
  call extend(b:tabula_stderr, a:data)
endfunction

" Helper function to run after job finishes
function! s:TabulaJobExit(job_id, code, event) abort
  " Handle errors
  if a:code != 0
    echohl ErrorMsg

    if exists('b:tabula_stderr') && !empty(b:tabula_stderr)
      for l:line in b:tabula_stderr
        if !empty(l:line)
          echom 'Tabula error: ' . l:line
        endif
      endfor
    elseif exists('b:tabula_stdout') && !empty(b:tabula_stdout)
      for l:line in b:tabula_stdout
        if !empty(l:line)
          echom 'Tabula error: ' . l:line
        endif
      endfor
    else
      echom 'Tabula error: exit code ' . a:code
    endif

    echohl None
  else
    " Reload file silently
    silent! checktime
  endif

  " Cleanup buffers
  unlet! b:tabula_stdout
  unlet! b:tabula_stderr

  " Release debounce lock
  let b:tabula_job_running = 0
endfunction

" Main function to execute Tabula on the current file
function! s:ExecuteTabula() abort
  " Debounce: skip if job already running
  if exists('b:tabula_job_running') && b:tabula_job_running
    return
  endif
  let b:tabula_job_running = 1

  " Save the file first
  silent! write!

  " Get the current file path
  let l:filepath = expand('%:p')

  " Build command as list (single source of truth)
  let l:cmd_list = [g:tabula_command]

  if &filetype ==# 'markdown'
    call add(l:cmd_list, '-m')
  endif

  if g:tabula_auto_format
    call add(l:cmd_list, '-a')
  endif

  call extend(l:cmd_list, ['-u', l:filepath])

  " Run asynchronously in Neovim
  if has('nvim')
    call jobstart(l:cmd_list, {
          \ 'on_stdout': function('s:TabulaOnStdout'),
          \ 'on_stderr': function('s:TabulaOnStderr'),
          \ 'on_exit': function('s:TabulaJobExit')
          \ })
  else
    " string version for Vim
    let l:cmd_string = join(map(copy(l:cmd_list), 'shellescape(v:val)'), ' ')
    let l:output = system(l:cmd_string)

    " Check for errors
    if v:shell_error != 0
      echohl ErrorMsg
      for l:line in split(l:output, "\n")
        if !empty(l:line)
          echom 'Tabula error: ' . l:line
        endif
      endfor
      echohl None
    else
      " Reload the file and refresh CSVView
      silent! checktime
    endif

    " Release debounce lock (sync case)
    let b:tabula_job_running = 0
  endif
endfunction

" Setup function for CSV files
function! s:SetupTabula() abort
  " Don't setup if already done for this buffer
  if exists('b:tabula_is_loaded')
    return
  endif
  let b:tabula_is_loaded = 1

  " Enable folding for #tabula markers (optional)
  if g:tabula_enable_folding
    setlocal foldmethod=marker
    setlocal foldlevel=0
  endif

  " Enable auto-read for external changes
  setlocal autoread

  " Setup autocommands for this buffer
  " Auto-execute on write
  augroup tabula_save
    autocmd! * <buffer>
    autocmd BufWritePost <buffer> call s:ExecuteTabula()
  augroup END

  " Mark auto-execution as enabled by default
  let b:tabula_auto_enabled = 1

  " Initialize debounce state
  let b:tabula_job_running = 0
endfunction

" Command to manually execute Tabula
command! Tabula call s:ExecuteTabula()

" Command to toggle auto-execution
command! TabulaToggle call s:ToggleTabula()

" Command to show status
command! TabulaStatus call s:TabulaStatus()

function! s:ToggleTabula() abort
  if exists('b:tabula_auto_enabled') && b:tabula_auto_enabled
    augroup tabula_save
      autocmd! * <buffer>
    augroup END
    let b:tabula_auto_enabled = 0
    echom 'Tabula auto-execution disabled'
  else
    augroup tabula_save
      autocmd! * <buffer>
      autocmd BufWritePost <buffer> call s:ExecuteTabula()
    augroup END
    let b:tabula_auto_enabled = 1
    echom 'Tabula auto-execution enabled'
  endif
endfunction

function! s:TabulaStatus() abort
  if exists('b:tabula_auto_enabled') && b:tabula_auto_enabled
    echom 'Tabula auto-execution is ENABLED'
  else
    echom 'Tabula auto-execution is DISABLED'
  endif
endfunction

" Restore compatibility options
let &cpo = s:save_cpo
unlet s:save_cpo
