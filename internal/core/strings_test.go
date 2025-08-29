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
			Expected: `<str "">`,
		},
		{
			Name: "single string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: " "},
				ast.StringExpression{Value: "world"},
			},
			Expected: `<str "hello world">`,
		},
		{
			Name: "with empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "start"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "end"},
			},
			Expected: `<str "startend">`,
		},
		{
			Name: "all empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: `<str "">`,
		},
		{
			Name: "special characters",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Line1\n"},
				ast.StringExpression{Value: "Line2\t"},
				ast.StringExpression{Value: "Line3"},
			},
			Expected: "<str \"Line1\nLine2\tLine3\">",
		},
		{
			Name: "string and integer",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
			},
			Error: `CONCATENATE(string...) got a wrong argument <int 42> in (CONCATENATE <str "hello"> <int 42>), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "CONCATENATE", testcases)
}

func TestLEN(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "LEN(string) expected 1 argument, but got 0 in (LEN), at <: input:0:0>",
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<int 5>",
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: "<int 0>",
		},
		{
			Name: "string with spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello world  "},
			},
			Expected: "<int 15>",
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `LEN(string) expected 1 argument, but got 2 in (LEN <str "hello"> <str "world">), at <: input:0:0>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			Error: "LEN(string) got a wrong argument <int 123> in (LEN <int 123>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "LEN", testcases)
}

func TestLOWER(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "LOWER(string) expected 1 argument, but got 0 in (LOWER), at <: input:0:0>",
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "mixed case string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
			},
			Expected: `<str "hello world">`,
		},
		{
			Name: "uppercase string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "HELLO"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "string with numbers",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello123"},
			},
			Expected: `<str "hello123">`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "World"},
			},
			Error: `LOWER(string) expected 1 argument, but got 2 in (LOWER <str "Hello"> <str "World">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "LOWER", testcases)
}

func TestUPPER(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "UPPER(string) expected 1 argument, but got 0 in (UPPER), at <: input:0:0>",
		},
		{
			Name: "simple string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "HELLO">`,
		},
		{
			Name: "mixed case string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
			},
			Expected: `<str "HELLO WORLD">`,
		},
		{
			Name: "lowercase string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "HELLO">`,
		},
		{
			Name: "string with numbers",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello123"},
			},
			Expected: `<str "HELLO123">`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `UPPER(string) expected 1 argument, but got 2 in (UPPER <str "hello"> <str "world">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "UPPER", testcases)
}

func TestTRIM(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "TRIM(string) expected 1 argument, but got 0 in (TRIM), at <: input:0:0>",
		},
		{
			Name: "string without spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "string with leading and trailing spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello world  "},
			},
			Expected: `<str "hello world">`,
		},
		{
			Name: "string with only leading spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "string with only trailing spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello  "},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "string with tabs and newlines",
			Input: []ast.Expression{
				ast.StringExpression{Value: "\t\n  hello  \n\t"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello  "},
				ast.StringExpression{Value: "  world  "},
			},
			Error: `TRIM(string) expected 1 argument, but got 2 in (TRIM <str "  hello  "> <str "  world  ">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "TRIM", testcases)
}

func TestEXACT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "EXACT(string, string) expected 2 arguments, but got 0 in (EXACT), at <: input:0:0>",
		},
		{
			Name: "single argument",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <str \"hello\">), at <: input:0:0>",
		},
		{
			Name: "identical strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<bool true>",
		},
		{
			Name: "different strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Expected: "<bool false>",
		},
		{
			Name: "case sensitive test",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<bool false>",
		},
		{
			Name: "empty strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: "<bool true>",
		},
		{
			Name: "spaces matter",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello "},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<bool false>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "extra"},
			},
			Error: "EXACT(string, string) expected 2 arguments, but got 3 in (EXACT <str \"hello\"> <str \"world\"> <str \"extra\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "EXACT", testcases)
}

func TestFIND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "FIND(string, string, [int]) expected 3 arguments, but got 0 in (FIND), at <: input:0:0>",
		},
		{
			Name: "basic substring search",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
			},
			Expected: "<int 6>",
		},
		{
			Name: "substring at beginning",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<int 0>",
		},
		{
			Name: "substring not found",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "xyz"},
			},
			Expected: "<int -1>",
		},
		{
			Name: "empty search string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
			},
			Expected: "<int 0>",
		},
		{
			Name: "search in empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<int -1>",
		},
		{
			Name: "both strings empty",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			Expected: "<int 0>",
		},
		{
			Name: "case sensitive search",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
			},
			Expected: "<int -1>",
		},
		{
			Name: "with start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello hello world"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			Expected: "<int 6>",
		},
		{
			Name: "start position beyond string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: "<int -1>",
		},
		{
			Name: "negative start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: -1},
			},
			Expected: "<int -1>",
		},
		{
			Name: "wrong type for third argument",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "not_int"},
			},
			Error: `FIND(string, string, [int]) got a wrong argument <str "not_int"> in (FIND <str "hello"> <str "world"> <str "not_int">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "FIND", testcases)
}

func TestLEFT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "LEFT(string, [int]) expected 2 arguments, but got 0 in (LEFT), at <: input:0:0>",
		},
		{
			Name: "single character default",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "h">`,
		},
		{
			Name: "with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "count larger than string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "zero count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			Expected: `<str "">`,
		},
		{
			Name: "negative count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			Expected: `<str "">`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<str "">`,
		},
		{
			Name: "empty string with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "">`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `LEFT(string, [int]) expected 2 arguments, but got 3 in (LEFT <str "hello"> <int 3> <str "extra">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "LEFT", testcases)
}

func TestRIGHT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "RIGHT(string, [int]) expected 2 arguments, but got 0 in (RIGHT), at <: input:0:0>",
		},
		{
			Name: "single character default",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<str "o">`,
		},
		{
			Name: "with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "world">`,
		},
		{
			Name: "count larger than string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "zero count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			Expected: `<str "">`,
		},
		{
			Name: "negative count",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			Expected: `<str "">`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<str "">`,
		},
		{
			Name: "empty string with count",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "">`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `RIGHT(string, [int]) expected 2 arguments, but got 3 in (RIGHT <str "hello"> <int 3> <str "extra">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "RIGHT", testcases)
}

func TestMID(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "MID(string, int, int) expected 3 arguments, but got 0 in (MID), at <: input:0:0>",
		},
		{
			Name: "basic usage",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 7},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "world">`,
		},
		{
			Name: "start at beginning",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "single character",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 1},
			},
			Expected: `<str "e">`,
		},
		{
			Name: "length larger than remaining string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 10},
			},
			Expected: `<str "llo">`,
		},
		{
			Name: "zero length",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 0},
			},
			Expected: `<str "">`,
		},
		{
			Name: "negative start position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -2},
				ast.IntExpression{Value: 3},
			},
			Expected: `<str "hel">`,
		},
		{
			Name: "empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "">`,
		},
		{
			Name: "extract entire string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
			},
			Error: `MID(string, int, int) expected 3 arguments, but got 2 in (MID <str "hello"> <int 2>), at <: input:0:0>`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			Error: `MID(string, int, int) expected 3 arguments, but got 4 in (MID <str "hello"> <int 2> <int 3> <str "extra">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "MID", testcases)
}

func TestSUBSTITUTE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 0 in (SUBSTITUTE), at <: input:0:0>",
		},
		{
			Name: "basic replacement",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			Expected: `<str "hello universe">`,
		},
		{
			Name: "replace first occurrence",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 1},
			},
			Expected: `<str "hello universe world">`,
		},
		{
			Name: "replace second occurrence",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 2},
			},
			Expected: `<str "hello world universe">`,
		},
		{
			Name: "replace all occurrences (zero index)",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 0},
			},
			Expected: `<str "hello universe universe universe">`,
		},
		{
			Name: "no match found",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "universe"},
				ast.StringExpression{Value: "galaxy"},
			},
			Expected: `<str "hello world">`,
		},
		{
			Name: "empty old value",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "x"},
			},
			Expected: `<str "hello">`,
		},
		{
			Name: "empty new value",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: ""},
			},
			Expected: `<str "hello ">`,
		},
		{
			Name: "case sensitive",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			Expected: `<str "Hello World">`,
		},
		{
			Name: "occurrence index beyond matches",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 3},
			},
			Expected: `<str "hello world">`,
		},
		{
			Name: "negative occurrence index",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: -1},
			},
			Error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <int -1> in (SUBSTITUTE <str "hello world world"> <str "world"> <str "universe"> <int -1>), at <: input:0:0>`,
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str "hello"> <str "world">), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "SUBSTITUTE", testcases)
}

func TestVALUE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "VALUE(string) expected 1 argument, but got 0 in (VALUE), at <: input:0:0>",
		},
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: "<str \"hello\">",
		},
		{
			Name: "boolean string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			Expected: "<bool true>",
		},
		{
			Name: "false boolean string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "false"},
			},
			Expected: "<bool false>",
		},
		{
			Name: "integer string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "7"},
			},
			Expected: "<int 7>",
		},
		{
			Name: "float string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "1.55"},
			},
			Expected: "<float 1.55>",
		},
		{
			Name: "negative integer string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "-42"},
			},
			Expected: "<int -42>",
		},
		{
			Name: "date string",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 15:39"},
			},
			Expected: "<date 2025-08-17 15:39:00>",
		},
		{
			Name: "multiple arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "7"},
				ast.StringExpression{Value: "42"},
			},
			Error: "VALUE(string) expected 1 argument, but got 2 in (VALUE <str \"7\"> <str \"42\">), at <: input:0:0>",
		},
		{
			Name: "non-string input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: "VALUE(string) got a wrong argument <int 42> in (VALUE <int 42>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "VALUE", testcases)
}