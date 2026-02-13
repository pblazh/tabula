" Vim syntax file
" Language: Tabula
" Maintainer: Tabula Team
" Latest Revision: 2026-02-13

if exists("b:current_syntax")
  finish
endif

" Keywords
syn keyword tabulaKeyword let fmt
syn keyword tabulaBoolean true false

" Preprocessor directives
syn match tabulaPreProc "^#include\>"

" Comments
syn match tabulaComment "//.*$"
syn region tabulaComment start="/\*" end="\*/"

" Cell references (case-insensitive)
" Match patterns like A1, B2, AA1, AB23, etc.
syn match tabulaCellRef "\<[A-Za-z]\+[0-9]\+\>"

" Cell ranges (e.g., A1:C3)
syn match tabulaRange "\<[A-Za-z]\+[0-9]\+:[A-Za-z]\+[0-9]\+\>"

" Numbers
syn match tabulaNumber "\<\d\+\>"
syn match tabulaFloat "\<\d\+\.\d\+\>"
syn match tabulaNumber "[-+]\d\+\>"
syn match tabulaFloat "[-+]\d\+\.\d\+\>"

" Strings
syn region tabulaString start=+"+ skip=+\\\\\|\\"+ end=+"+

" Operators
syn match tabulaOperator "[-+*/<>=!]"
syn match tabulaOperator "=="
syn match tabulaOperator "!="
syn match tabulaOperator "<="
syn match tabulaOperator ">="
syn match tabulaOperator "&&"
syn match tabulaOperator "||"

" Built-in functions
" Number functions
syn keyword tabulaFunction SUM AVERAGE MIN MAX ABS ROUND SQRT MOD POWER
syn keyword tabulaFunction FLOOR CEILING INT TRUNC EVEN ODD SIGN
syn keyword tabulaFunction PRODUCT COUNT COUNTA COUNTBLANK
syn keyword tabulaFunction SUMPRODUCT SUMSQ GCD LCM FACT

" String functions
syn keyword tabulaFunction CONCATENATE LEFT RIGHT MID UPPER LOWER
syn keyword tabulaFunction TRIM LEN FIND SEARCH SUBSTITUTE REPLACE
syn keyword tabulaFunction TEXT VALUE REPT EXACT

" Date functions
syn keyword tabulaFunction DATE DATEVALUE YEAR MONTH DAY
syn keyword tabulaFunction HOUR MINUTE SECOND NOW TODAY
syn keyword tabulaFunction DATEDIF WEEKDAY WORKDAY NETWORKDAYS
syn keyword tabulaFunction EDATE EOMONTH TIME TIMEVALUE

" Logical functions
syn keyword tabulaFunction IF AND OR NOT IFERROR IFNA

" Lookup functions
syn keyword tabulaFunction COLUMN ROW COLUMNS ROWS ADDRESS REF RANGE

" Info functions
syn keyword tabulaFunction ISNUMBER ISTEXT ISLOGICAL ISBLANK ISERROR

" Special function
syn keyword tabulaFunction EXEC

" Delimiters
syn match tabulaDelimiter ";"
syn match tabulaDelimiter ","
syn match tabulaDelimiter "[()]"

" Format specifiers in strings (for fmt statements)
syn match tabulaFormatSpec "%[-+0 #]*\d*\.\?\d*[diouxXeEfFgGaAcspv%]" contained containedin=tabulaString

" Highlight definitions
hi def link tabulaKeyword Keyword
hi def link tabulaBoolean Boolean
hi def link tabulaPreProc PreProc
hi def link tabulaComment Comment
hi def link tabulaCellRef Identifier
hi def link tabulaRange Special
hi def link tabulaNumber Number
hi def link tabulaFloat Float
hi def link tabulaString String
hi def link tabulaOperator Operator
hi def link tabulaFunction Function
hi def link tabulaDelimiter Delimiter
hi def link tabulaFormatSpec SpecialChar

let b:current_syntax = "tabula"
