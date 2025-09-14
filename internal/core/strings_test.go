package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
)

func TestCONCATENATE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "empty input",
			Input:    []ast.Expression{},
			Expected: `""`,
		},
		{
			Name: "single string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: " "},
				ast.StringExpression{Value: "world"},
			},
			Expected: `"hello world"`,
		},
		{
			Name: "with empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "start"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "end"},
			},
			Expected: `"startend"`,
		},
		{
			Name: "all empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: `""`,
		},
		{
			Name: "special characters",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Line1\n"},
				ast.StringExpression{Value: "Line2\t"},
				ast.StringExpression{Value: "Line3"},
			},
			Expected: `"Line1
Line2	Line3"`,
		},
		{
			Name: "string and integer",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
			},
			Error: `CONCATENATE(values:string...):string received invalid argument 42 in CONCATENATE("hello", 42), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "CONCATENATE", testcases)
}

func TestLEN(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `LEN(value:string):number expects 1 argument, got 0 in LEN(), at <: input:0:0>`,
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `5`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `0`,
		},
		{
			Name: "string with spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello world  "},
			},
			Expected: `15`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `LEN(value:string):number expects 1 argument, got 2 in LEN("hello", "world"), at <: input:0:0>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			Error: `LEN(value:string):number received invalid argument 123 in LEN(123), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "LEN", testcases)
}

func TestLOWER(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `LOWER(value:string):string expects 1 argument, got 0 in LOWER(), at <: input:0:0>`,
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "mixed case string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
			},
			Expected: `"hello world"`,
		},
		{
			Name: "uppercase string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "HELLO"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "string with numbers",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello123"},
			},
			Expected: `"hello123"`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "World"},
			},
			Error: `LOWER(value:string):string expects 1 argument, got 2 in LOWER("Hello", "World"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "LOWER", testcases)
}

func TestUPPER(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `UPPER(value:string):string expects 1 argument, got 0 in UPPER(), at <: input:0:0>`,
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"HELLO"`,
		},
		{
			Name: "mixed case string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
			},
			Expected: `"HELLO WORLD"`,
		},
		{
			Name: "lowercase string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"HELLO"`,
		},
		{
			Name: "string with numbers",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello123"},
			},
			Expected: `"HELLO123"`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `UPPER(value:string):string expects 1 argument, got 2 in UPPER("hello", "world"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "UPPER", testcases)
}

func TestTRIM(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `TRIM(value:string):string expects 1 argument, got 0 in TRIM(), at <: input:0:0>`,
		},
		{
			Name: "string without spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "string with leading and trailing spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello world  "},
			},
			Expected: `"hello world"`,
		},
		{
			Name: "string with only leading spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "string with only trailing spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello  "},
			},
			Expected: `"hello"`,
		},
		{
			Name: "string with tabs and newlines",
			Input: []ast.Expression{
				ast.StringExpression{Value: "\t\n  hello  \n\t"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello  "},
				ast.StringExpression{Value: "  world  "},
			},
			Error: `TRIM(value:string):string expects 1 argument, got 2 in TRIM("  hello  ", "  world  "), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "TRIM", testcases)
}

func TestEXACT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `EXACT(a:string, b:string):boolean expects 2 arguments, got 0 in EXACT(), at <: input:0:0>`,
		},
		{
			Name: "single argument",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Error: `EXACT(a:string, b:string):boolean expects 2 arguments, got 1 in EXACT("hello"), at <: input:0:0>`,
		},
		{
			Name: "identical strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: `true`,
		},
		{
			Name: "different strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Expected: `false`,
		},
		{
			Name: "case sensitive test",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: `false`,
		},
		{
			Name: "empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: `true`,
		},
		{
			Name: "spaces matter",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello "},
				ast.StringExpression{Value: "hello"},
			},
			Expected: `false`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "extra"},
			},
			Error: `EXACT(a:string, b:string):boolean expects 2 arguments, got 3 in EXACT("hello", "world", "extra"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "EXACT", testcases)
}

func TestFIND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `FIND(what:string, where:string, [start:int]):number expects 3 arguments, got 0 in FIND(), at <: input:0:0>`,
		},
		{
			Name: "basic substring search",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
			},
			Expected: `6`,
		},
		{
			Name: "substring at beginning",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: `0`,
		},
		{
			Name: "substring not found",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "xyz"},
			},
			Expected: `-1`,
		},
		{
			Name: "empty search string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
			},
			Expected: `0`,
		},
		{
			Name: "search in empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "hello"},
			},
			Expected: `-1`,
		},
		{
			Name: "both strings empty",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: `0`,
		},
		{
			Name: "case sensitive search",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
			},
			Expected: `-1`,
		},
		{
			Name: "with start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello hello world"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			Expected: `6`,
		},
		{
			Name: "start position beyond string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: `-1`,
		},
		{
			Name: "negative start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: -1},
			},
			Expected: `-1`,
		},
		{
			Name: "wrong type for third argument",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "not_int"},
			},
			Error: `FIND(what:string, where:string, [start:int]):number received invalid argument "not_int" in FIND("hello", "world", "not_int"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "FIND", testcases)
}

func TestLEFT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `LEFT(value:string, [amount:int]):string expects 2 arguments, got 0 in LEFT(), at <: input:0:0>`,
		},
		{
			Name: "single character default",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"h"`,
		},
		{
			Name: "with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			Expected: `"hello"`,
		},
		{
			Name: "count larger than string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: `"hello"`,
		},
		{
			Name: "zero count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			Expected: `""`,
		},
		{
			Name: "negative count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			Expected: `""`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `""`,
		},
		{
			Name: "empty string with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			Expected: `""`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `LEFT(value:string, [amount:int]):string expects 2 arguments, got 3 in LEFT("hello", 3, "extra"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "LEFT", testcases)
}

func TestRIGHT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `RIGHT(value:string, [amount:int]):string expects 2 arguments, got 0 in RIGHT(), at <: input:0:0>`,
		},
		{
			Name: "single character default",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"o"`,
		},
		{
			Name: "with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			Expected: `"world"`,
		},
		{
			Name: "count larger than string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: `"hello"`,
		},
		{
			Name: "zero count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			Expected: `""`,
		},
		{
			Name: "negative count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			Expected: `""`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `""`,
		},
		{
			Name: "empty string with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			Expected: `""`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `RIGHT(value:string, [amount:int]):string expects 2 arguments, got 3 in RIGHT("hello", 3, "extra"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "RIGHT", testcases)
}

func TestMID(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `MID(value:string, start:int, amount:int):string expects 3 arguments, got 0 in MID(), at <: input:0:0>`,
		},
		{
			Name: "basic usage",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 7},
				ast.IntExpression{Value: 5},
			},
			Expected: `"world"`,
		},
		{
			Name: "start at beginning",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `"hello"`,
		},
		{
			Name: "single character",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 1},
			},
			Expected: `"e"`,
		},
		{
			Name: "length larger than remaining string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 10},
			},
			Expected: `"llo"`,
		},
		{
			Name: "zero length",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 0},
			},
			Expected: `""`,
		},
		{
			Name: "negative start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -2},
				ast.IntExpression{Value: 3},
			},
			Expected: `"hel"`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `""`,
		},
		{
			Name: "extract entire string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `"hello"`,
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
			},
			Error: `MID(value:string, start:int, amount:int):string expects 3 arguments, got 2 in MID("hello", 2), at <: input:0:0>`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `MID(value:string, start:int, amount:int):string expects 3 arguments, got 4 in MID("hello", 2, 3, "extra"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "MID", testcases)
}

func TestSUBSTITUTE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `SUBSTITUTE(text:string, old:string, new:string, [instances:int]):string expects 4 arguments, got 0 in SUBSTITUTE(), at <: input:0:0>`,
		},
		{
			Name: "basic replacement",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			Expected: `"hello universe"`,
		},
		{
			Name: "replace first occurrence",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 1},
			},
			Expected: `"hello universe world"`,
		},
		{
			Name: "replace second occurrence",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 2},
			},
			Expected: `"hello world universe"`,
		},
		{
			Name: "replace all occurrences (zero index)",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 0},
			},
			Expected: `"hello universe universe universe"`,
		},
		{
			Name: "no match found",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "universe"},
				ast.StringExpression{Value: "galaxy"},
			},
			Expected: `"hello world"`,
		},
		{
			Name: "empty old value",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "x"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "empty new value",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: ""},
			},
			Expected: `"hello "`,
		},
		{
			Name: "case sensitive",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			Expected: `"Hello World"`,
		},
		{
			Name: "occurrence index beyond matches",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 3},
			},
			Expected: `"hello world"`,
		},
		{
			Name: "negative occurrence index",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: -1},
			},
			Error: `SUBSTITUTE(text:string, old:string, new:string, [instances:int]):string received invalid argument -1 in SUBSTITUTE("hello world world", "world", "universe", -1), at <: input:0:0>`,
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `SUBSTITUTE(text:string, old:string, new:string, [instances:int]):string expects 4 arguments, got 2 in SUBSTITUTE("hello", "world"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "SUBSTITUTE", testcases)
}

func TestVALUE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `VALUE(value:string):number expects 1 argument, got 0 in VALUE(), at <: input:0:0>`,
		},
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "boolean string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			Expected: `true`,
		},
		{
			Name: "false boolean string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "false"},
			},
			Expected: `false`,
		},
		{
			Name: "integer string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "7"},
			},
			Expected: `7`,
		},
		{
			Name: "float string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "1.55"},
			},
			Expected: `1.55`,
		},
		{
			Name: "negative integer string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "-42"},
			},
			Expected: `-42`,
		},
		{
			Name: "date string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 15:39"},
			},
			Expected: `<2025-08-17 15:39:00>`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "7"},
				ast.StringExpression{Value: "42"},
			},
			Error: `VALUE(value:string):number expects 1 argument, got 2 in VALUE("7", "42"), at <: input:0:0>`,
		},
		{
			Name: "non-string input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `VALUE(value:string):number received invalid argument 42 in VALUE(42), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "VALUE", testcases)
}
