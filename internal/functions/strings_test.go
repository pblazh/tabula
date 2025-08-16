package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestStringFunctions(t *testing.T) {
	testcases := []struct {
		name  string
		input []ast.Expression
		cases []inputCase
	}{
		// Empty input
		{
			name:  "empty input",
			input: []ast.Expression{},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "">`,
				},
				{
					f:     "LEN",
					error: "LEN(string) expected 1 argument, but got 0 in (LEN), at <: input:0:0>",
				},
				{
					f:     "LOWER",
					error: "LOWER(string) expected 1 argument, but got 0 in (LOWER), at <: input:0:0>",
				},
				{
					f:     "UPPER",
					error: "UPPER(string) expected 1 argument, but got 0 in (UPPER), at <: input:0:0>",
				},
				{
					f:     "TRIM",
					error: "TRIM(string) expected 1 argument, but got 0 in (TRIM), at <: input:0:0>",
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 0 in (EXACT), at <: input:0:0>",
				},
				{
					f:     "FIND",
					error: "FIND(string, string, [int]) expected 3 arguments, but got 0 in (FIND), at <: input:0:0>",
				},
				{
					f:     "LEFT",
					error: "LEFT(string, [int]) expected 2 arguments, but got 0 in (LEFT), at <: input:0:0>",
				},
				{
					f:     "RIGHT",
					error: "RIGHT(string, [int]) expected 2 arguments, but got 0 in (RIGHT), at <: input:0:0>",
				},
				{
					f:     "MID",
					error: "MID(string, int, int) expected 3 arguments, but got 0 in (MID), at <: input:0:0>",
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 0 in (SUBSTITUTE), at <: input:0:0>",
				},
			},
		},
		// Single string
		{
			name: "single string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "hello">`,
				},
				{
					f:        "LEN",
					expected: "<int 5>",
				},
				{
					f:        "LOWER",
					expected: `<str "hello">`,
				},
				{
					f:        "UPPER",
					expected: `<str "HELLO">`,
				},
				{
					f:        "TRIM",
					expected: `<str "hello">`,
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "FIND",
					error: "FIND(string, string, [int]) expected 3 arguments, but got 1 in (FIND <str \"hello\">), at <: input:0:0>",
				},
				{
					f:        "LEFT",
					expected: `<str "h">`,
				},
				{
					f:        "RIGHT",
					expected: `<str "o">`,
				},
				{
					f:     "MID",
					error: "MID(string, int, int) expected 3 arguments, but got 1 in (MID <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 1 in (SUBSTITUTE <str \"hello\">), at <: input:0:0>",
				},
			},
		},
		// Mixed case string
		{
			name: "mixed case string",
			input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "Hello World">`,
				},
				{
					f:        "LEN",
					expected: "<int 11>",
				},
				{
					f:        "LOWER",
					expected: `<str "hello world">`,
				},
				{
					f:        "UPPER",
					expected: `<str "HELLO WORLD">`,
				},
				{
					f:        "TRIM",
					expected: `<str "Hello World">`,
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <str \"Hello World\">), at <: input:0:0>",
				},
				{
					f:        "LEFT",
					expected: `<str "H">`,
				},
				{
					f:        "RIGHT",
					expected: `<str "d">`,
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 1 in (SUBSTITUTE <str \"Hello World\">), at <: input:0:0>",
				},
			},
		},
		// String with whitespace
		{
			name: "string with whitespace",
			input: []ast.Expression{
				ast.StringExpression{Value: "  hello world  "},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "  hello world  ">`,
				},
				{
					f:        "LEN",
					expected: "<int 15>",
				},
				{
					f:        "LOWER",
					expected: `<str "  hello world  ">`,
				},
				{
					f:        "UPPER",
					expected: `<str "  HELLO WORLD  ">`,
				},
				{
					f:        "TRIM",
					expected: `<str "hello world">`,
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <str \"  hello world  \">), at <: input:0:0>",
				},
				{
					f:        "LEFT",
					expected: `<str " ">`,
				},
				{
					f:        "RIGHT",
					expected: `<str " ">`,
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 1 in (SUBSTITUTE <str \"  hello world  \">), at <: input:0:0>",
				},
			},
		},
		// Multiple strings
		{
			name: "multiple strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: " "},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "hello world">`,
				},
				{
					f:     "LEN",
					error: `LEN(string) expected 1 argument, but got 3 in (LEN <str "hello"> <str " "> <str "world">), at <: input:0:0>`,
				},
				{
					f:     "LOWER",
					error: `LOWER(string) expected 1 argument, but got 3 in (LOWER <str "hello"> <str " "> <str "world">), at <: input:0:0>`,
				},
				{
					f:     "UPPER",
					error: `UPPER(string) expected 1 argument, but got 3 in (UPPER <str "hello"> <str " "> <str "world">), at <: input:0:0>`,
				},
				{
					f:     "TRIM",
					error: `TRIM(string) expected 1 argument, but got 3 in (TRIM <str "hello"> <str " "> <str "world">), at <: input:0:0>`,
				},
				{
					f:     "EXACT",
					error: `EXACT(string, string) expected 2 arguments, but got 3 in (EXACT <str "hello"> <str " "> <str "world">), at <: input:0:0>`,
				},
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello">`,
				},
			},
		},
		// Empty strings
		{
			name: "with empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "start"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "end"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "startend">`,
				},
				{
					f:        "SUBSTITUTE",
					expected: `<str "start">`,
				},
			},
		},
		// All empty strings
		{
			name: "all empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "">`,
				},
				{
					f:        "SUBSTITUTE",
					expected: `<str "">`,
				},
			},
		},
		// Special characters
		{
			name: "special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\n"},
				ast.StringExpression{Value: "Line2\t"},
				ast.StringExpression{Value: "Line3"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: "<str \"Line1\nLine2\tLine3\">",
				},
				{
					f:        "SUBSTITUTE",
					expected: "<str \"Line1\n\">",
				},
			},
		},
		// Two strings for EXACT testing
		{
			name: "two identical strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "EXACT",
					expected: "<bool true>",
				},
				{
					f:        "FIND",
					expected: "<int 0>",
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str \"hello\"> <str \"hello\">), at <: input:0:0>",
				},
			},
		},
		{
			name: "two different strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "EXACT",
					expected: "<bool false>",
				},
				{
					f:        "FIND",
					expected: "<int -1>",
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str \"hello\"> <str \"world\">), at <: input:0:0>",
				},
			},
		},
		{
			name: "case sensitivity test",
			input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "EXACT",
					expected: "<bool false>",
				},
				{
					f:        "FIND",
					expected: "<int -1>",
				},
				{
					f:     "SUBSTITUTE",
					error: "SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str \"Hello\"> <str \"hello\">), at <: input:0:0>",
				},
			},
		},

		// Numbers and other types (should error)
		{
			name: "string and integer",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <int 42> in (CONCATENATE <str "hello"> <int 42>), at <: input:0:0>`,
				},
				{
					f:     "EXACT",
					error: `EXACT(string, string) got a wrong argument <int 42> in (EXACT <str "hello"> <int 42>), at <: input:0:0>`,
				},
				{
					f:     "FIND",
					error: `FIND(string, string, [int]) got a wrong argument <int 42> in (FIND <str "hello"> <int 42>), at <: input:0:0>`,
				},
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str "hello"> <int 42>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "string and float",
			input: []ast.Expression{
				ast.StringExpression{Value: "value: "},
				ast.FloatExpression{Value: 3.14},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <float 3.14> in (CONCATENATE <str "value: "> <float 3.14>), at <: input:0:0>`,
				},
				{
					f:     "EXACT",
					error: `EXACT(string, string) got a wrong argument <float 3.14> in (EXACT <str "value: "> <float 3.14>), at <: input:0:0>`,
				},
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str "value: "> <float 3.14>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "string and boolean",
			input: []ast.Expression{
				ast.StringExpression{Value: "result: "},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <bool true> in (CONCATENATE <str "result: "> <bool true>), at <: input:0:0>`,
				},
				{
					f:     "EXACT",
					error: `EXACT(string, string) got a wrong argument <bool true> in (EXACT <str "result: "> <bool true>), at <: input:0:0>`,
				},
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str "result: "> <bool true>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "integer only",
			input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <int 123> in (CONCATENATE <int 123>), at <: input:0:0>`,
				},
				{
					f:     "LEN",
					error: "LEN(string) got a wrong argument <int 123> in (LEN <int 123>), at <: input:0:0>",
				},
				{
					f:     "LOWER",
					error: "LOWER(string) got a wrong argument <int 123> in (LOWER <int 123>), at <: input:0:0>",
				},
				{
					f:     "UPPER",
					error: "UPPER(string) got a wrong argument <int 123> in (UPPER <int 123>), at <: input:0:0>",
				},
				{
					f:     "TRIM",
					error: "TRIM(string) got a wrong argument <int 123> in (TRIM <int 123>), at <: input:0:0>",
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <int 123>), at <: input:0:0>",
				},
				{
					f:     "LEFT",
					error: "LEFT(string, [int]) got a wrong argument <int 123> in (LEFT <int 123>), at <: input:0:0>",
				},
				{
					f:     "RIGHT",
					error: "RIGHT(string, [int]) got a wrong argument <int 123> in (RIGHT <int 123>), at <: input:0:0>",
				},
				{
					f:     "MID",
					error: "MID(string, int, int) expected 3 arguments, but got 1 in (MID <int 123>), at <: input:0:0>",
				},
			},
		},
		{
			name: "boolean only",
			input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <bool false> in (CONCATENATE <bool false>), at <: input:0:0>`,
				},
				{
					f:     "LEN",
					error: "LEN(string) got a wrong argument <bool false> in (LEN <bool false>), at <: input:0:0>",
				},
				{
					f:     "LOWER",
					error: "LOWER(string) got a wrong argument <bool false> in (LOWER <bool false>), at <: input:0:0>",
				},
				{
					f:     "UPPER",
					error: "UPPER(string) got a wrong argument <bool false> in (UPPER <bool false>), at <: input:0:0>",
				},
				{
					f:     "TRIM",
					error: "TRIM(string) got a wrong argument <bool false> in (TRIM <bool false>), at <: input:0:0>",
				},
				{
					f:     "EXACT",
					error: "EXACT(string, string) expected 2 arguments, but got 1 in (EXACT <bool false>), at <: input:0:0>",
				},
			},
		},

		// FIND function specific tests
		{
			name: "FIND: basic substring search",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 6>",
				},
			},
		},
		{
			name: "FIND: substring at beginning",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 0>",
				},
			},
		},
		{
			name: "FIND: substring not found",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "xyz"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		{
			name: "FIND: empty search string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 0>",
				},
			},
		},
		{
			name: "FIND: search in empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		{
			name: "FIND: both strings empty",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 0>",
				},
			},
		},
		{
			name: "FIND: case sensitive search",
			input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		{
			name: "FIND: multiple occurrences",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello hello world"},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 0>",
				},
			},
		},
		{
			name: "FIND: single character",
			input: []ast.Expression{
				ast.StringExpression{Value: "abcdef"},
				ast.StringExpression{Value: "d"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 3>",
				},
			},
		},
		{
			name: "FIND: special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\nLine2\tEnd"},
				ast.StringExpression{Value: "\n"},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 5>",
				},
			},
		},
		// FIND with start position (3 arguments)
		{
			name: "FIND: with start position",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello hello world"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 6>",
				},
			},
		},
		{
			name: "FIND: start position at exact match",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: 6},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 6>",
				},
			},
		},
		{
			name: "FIND: start position past match",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		{
			name: "FIND: start position zero",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int 0>",
				},
			},
		},
		{
			name: "FIND: start position beyond string length",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		{
			name: "FIND: negative start position",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: -1},
			},
			cases: []inputCase{
				{
					f:        "FIND",
					expected: "<int -1>",
				},
			},
		},
		// FIND error cases
		{
			name: "FIND: too many arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: 0},
				ast.StringExpression{Value: "extra"},
			},
			cases: []inputCase{
				{
					f:     "FIND",
					error: `FIND(string, string, [int]) expected 3 arguments, but got 4 in (FIND <str "hello"> <str "world"> <int 0> <str "extra">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "FIND: wrong type for third argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "not_int"},
			},
			cases: []inputCase{
				{
					f:     "FIND",
					error: `FIND(string, string, [int]) got a wrong argument <str "not_int"> in (FIND <str "hello"> <str "world"> <str "not_int">), at <: input:0:0>`,
				},
			},
		},

		// LEFT function specific tests
		{
			name: "LEFT: basic usage with count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "LEFT: count larger than string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "LEFT: count equal to string length",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "LEFT: zero count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "LEFT: negative count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "LEFT: single character string",
			input: []ast.Expression{
				ast.StringExpression{Value: "a"},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "a">`,
				},
			},
		},
		{
			name: "LEFT: empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "LEFT: empty string with count",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "LEFT: string with spaces",
			input: []ast.Expression{
				ast.StringExpression{Value: "  hello  "},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "  h">`,
				},
			},
		},
		{
			name: "LEFT: string with special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\nLine2\tEnd"},
				ast.IntExpression{Value: 7},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: "<str \"Line1\nL\">",
				},
			},
		},
		{
			name: "LEFT: count one",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "h">`,
				},
			},
		},
		{
			name: "LEFT: two character string without count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hi"},
			},
			cases: []inputCase{
				{
					f:        "LEFT",
					expected: `<str "h">`,
				},
			},
		},
		// LEFT error cases
		{
			name: "LEFT: too many arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			cases: []inputCase{
				{
					f:     "LEFT",
					error: `LEFT(string, [int]) expected 2 arguments, but got 3 in (LEFT <str "hello"> <int 3> <str "extra">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "LEFT: wrong type for second argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "not_int"},
			},
			cases: []inputCase{
				{
					f:     "LEFT",
					error: `LEFT(string, [int]) got a wrong argument <str "not_int"> in (LEFT <str "hello"> <str "not_int">), at <: input:0:0>`,
				},
			},
		},

		// RIGHT function specific tests
		{
			name: "RIGHT: basic usage with count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "world">`,
				},
			},
		},
		{
			name: "RIGHT: count larger than string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "RIGHT: count equal to string length",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "RIGHT: zero count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "RIGHT: negative count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -5},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "RIGHT: single character string",
			input: []ast.Expression{
				ast.StringExpression{Value: "a"},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "a">`,
				},
			},
		},
		{
			name: "RIGHT: empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "RIGHT: empty string with count",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "RIGHT: string with spaces",
			input: []ast.Expression{
				ast.StringExpression{Value: "  hello  "},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "o  ">`,
				},
			},
		},
		{
			name: "RIGHT: string with special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\nLine2\tEnd"},
				ast.IntExpression{Value: 7},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: "<str \"ne2\tEnd\">",
				},
			},
		},
		{
			name: "RIGHT: count one",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "o">`,
				},
			},
		},
		{
			name: "RIGHT: two character string without count",
			input: []ast.Expression{
				ast.StringExpression{Value: "hi"},
			},
			cases: []inputCase{
				{
					f:        "RIGHT",
					expected: `<str "i">`,
				},
			},
		},
		// RIGHT error cases
		{
			name: "RIGHT: too many arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			cases: []inputCase{
				{
					f:     "RIGHT",
					error: `RIGHT(string, [int]) expected 2 arguments, but got 3 in (RIGHT <str "hello"> <int 3> <str "extra">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "RIGHT: wrong type for second argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "not_int"},
			},
			cases: []inputCase{
				{
					f:     "RIGHT",
					error: `RIGHT(string, [int]) got a wrong argument <str "not_int"> in (RIGHT <str "hello"> <str "not_int">), at <: input:0:0>`,
				},
			},
		},

		// MID function specific tests
		{
			name: "MID: basic usage",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 7},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "world">`,
				},
			},
		},
		{
			name: "MID: start at beginning",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "MID: single character",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "e">`,
				},
			},
		},
		{
			name: "MID: length larger than remaining string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "llo">`,
				},
			},
		},
		{
			name: "MID: start at string boundary",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 4},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "lo">`,
				},
			},
		},
		{
			name: "MID: zero length",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "MID: negative start position",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: -2},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "hel">`,
				},
			},
		},
		{
			name: "MID: empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "MID: single character string",
			input: []ast.Expression{
				ast.StringExpression{Value: "a"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "a">`,
				},
			},
		},
		{
			name: "MID: extract entire string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "MID: extract with spaces",
			input: []ast.Expression{
				ast.StringExpression{Value: "  hello  world  "},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 7},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "hello  ">`,
				},
			},
		},
		{
			name: "MID: extract with special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\nLine2\tEnd"},
				ast.IntExpression{Value: 6},
				ast.IntExpression{Value: 6},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: "<str \"\nLine2\">",
				},
			},
		},
		{
			name: "MID: start near string end",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 5},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "o">`,
				},
			},
		},
		{
			name: "MID: start at last character",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 5},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "MID",
					expected: `<str "o">`,
				},
			},
		},
		// MID error cases
		{
			name: "MID: missing arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:     "MID",
					error: `MID(string, int, int) expected 3 arguments, but got 2 in (MID <str "hello"> <int 2>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "MID: too many arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.StringExpression{Value: "extra"},
			},
			cases: []inputCase{
				{
					f:     "MID",
					error: `MID(string, int, int) expected 3 arguments, but got 4 in (MID <str "hello"> <int 2> <int 3> <str "extra">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "MID: wrong type for start position",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "not_int"},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:     "MID",
					error: `MID(string, int, int) got a wrong argument <str "not_int"> in (MID <str "hello"> <str "not_int"> <int 3>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "MID: wrong type for length",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 2},
				ast.StringExpression{Value: "not_int"},
			},
			cases: []inputCase{
				{
					f:     "MID",
					error: `MID(string, int, int) got a wrong argument <str "not_int"> in (MID <str "hello"> <int 2> <str "not_int">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "MID: wrong type for string",
			input: []ast.Expression{
				ast.IntExpression{Value: 123},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:     "MID",
					error: `MID(string, int, int) got a wrong argument <int 123> in (MID <int 123> <int 1> <int 2>), at <: input:0:0>`,
				},
			},
		},

		// SUBSTITUTE function specific tests
		{
			name: "SUBSTITUTE: basic replacement",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello universe">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: replace first occurrence",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello universe world">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: replace second occurrence",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello world universe">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: replace all occurrences",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello universe universe universe">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: no match found",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "universe"},
				ast.StringExpression{Value: "galaxy"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello world">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: empty old value",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "x"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: empty new value",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello ">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "x"},
				ast.StringExpression{Value: "y"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: case sensitive",
			input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "Hello World">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: replace with same value",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello world">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: single character replacement",
			input: []ast.Expression{
				ast.StringExpression{Value: "a,b,c,d"},
				ast.StringExpression{Value: ","},
				ast.StringExpression{Value: ";"},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "a;b;c;d">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: zero occurrence index",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello universe universe">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: occurrence index beyond matches",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:        "SUBSTITUTE",
					expected: `<str "hello world">`,
				},
			},
		},
		{
			name: "SUBSTITUTE: negative occurrence index",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello world world"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: -1},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <int -1> in (SUBSTITUTE <str "hello world world"> <str "world"> <str "universe"> <int -1>), at <: input:0:0>`,
				},
			},
		},
		// SUBSTITUTE error cases
		{
			name: "SUBSTITUTE: missing arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 2 in (SUBSTITUTE <str "hello"> <str "world">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "SUBSTITUTE: too many arguments",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.IntExpression{Value: 1},
				ast.StringExpression{Value: "extra"},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) expected 4 arguments, but got 5 in (SUBSTITUTE <str "hello"> <str "world"> <str "universe"> <int 1> <str "extra">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "SUBSTITUTE: wrong type for source string",
			input: []ast.Expression{
				ast.IntExpression{Value: 123},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <int 123> in (SUBSTITUTE <int 123> <str "world"> <str "universe">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "SUBSTITUTE: wrong type for old value",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
				ast.StringExpression{Value: "universe"},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <int 42> in (SUBSTITUTE <str "hello"> <int 42> <str "universe">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "SUBSTITUTE: wrong type for new value",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.IntExpression{Value: 123},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <int 123> in (SUBSTITUTE <str "hello"> <str "world"> <int 123>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "SUBSTITUTE: wrong type for occurrence index",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "universe"},
				ast.StringExpression{Value: "not_int"},
			},
			cases: []inputCase{
				{
					f:     "SUBSTITUTE",
					error: `SUBSTITUTE(string, string, string, [int]) got a wrong argument <str "not_int"> in (SUBSTITUTE <str "hello"> <str "world"> <str "universe"> <str "not_int">), at <: input:0:0>`,
				},
			},
		},
		{
			name: "VALUE of string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "VALUE",
					expected: "<str \"hello\">",
				},
			},
		},
		{
			name: "VALUE of boolean",
			input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			cases: []inputCase{
				{
					f:        "VALUE",
					expected: "<bool true>",
				},
			},
		},
		{
			name: "VALUE of float",
			input: []ast.Expression{
				ast.StringExpression{Value: "1.55"},
			},
			cases: []inputCase{
				{
					f:        "VALUE",
					expected: "<float 1.55>",
				},
			},
		},
		{
			name: "VALUE of int",
			input: []ast.Expression{
				ast.StringExpression{Value: "7"},
			},
			cases: []inputCase{
				{
					f:        "VALUE",
					expected: "<int 7>",
				},
			},
		},
		{
			name: "VALUE of date",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 15:39"},
			},
			cases: []inputCase{
				{
					f:        "VALUE",
					expected: "<date 2025-08-17 15:39:00>",
				},
			},
		},
	}

	for _, tc := range testcases {
		for _, c := range tc.cases {
			t.Run(tc.name+":"+c.f, func(t *testing.T) {
				result, err := DispatchMap[c.f](ast.CallExpression{
					Identifier: ast.IdentifierExpression{
						Value: c.f,
						Token: lexer.Token{Literal: c.f},
					}, Arguments: tc.input,
				}, tc.input...)

				if c.error != "" {
					if err == nil {
						t.Errorf("Expected error %q but got result: %v", c.error, result)
						return
					}
					if err.Error() != c.error {
						t.Errorf("Expected error %q, got %q", c.error, err.Error())
					}
					return
				}

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				if result.String() != c.expected {
					t.Errorf("Expected %s, got %s", c.expected, result.String())
				}
			})
		}
	}
}
