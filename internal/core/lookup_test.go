package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
)

func TestADDRESS(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 0 at input:0:0`,
		},
		{
			Name: "single int",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
			},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 1 at input:0:0`,
		},
		{
			Name: "multiple ints",
			Input: []ast.Node{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 3 at input:0:0`,
		},
		{
			Name: "with string column",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
				ast.StringExpression{Value: "A"},
			},
			Error: `ADDRESS(row:int, column:int):string invalid argument "A" at input:0:0`,
		},
		{
			Name: "happy path",
			Input: []ast.Node{
				ast.IntExpression{Value: 4},
				ast.IntExpression{Value: 2},
			},
			Expected: `B4`,
		},
	}

	RunFunctionTest(t, "ADDRESS", testcases, map[string]string{}, [][]string{}, map[string]string{})
}

func TestROW(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ROW(cell:string):int expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `ROW(cell:string):int expects 1 argument, got 2 at input:0:0`,
		},
		{
			Name: "with an int column",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
			},
			Error: `ROW(cell:string):int invalid argument 42 at input:0:0`,
		},
		{
			Name: "with an Identifier",
			Input: []ast.Node{
				ast.IdentifierExpression{Value: "B4"},
			},
			Expected: `4`,
		},
		{
			Name: "with a Range",
			Input: []ast.Node{
				ast.RangeExpression{Value: []string{"B4", "C5"}},
			},
			Expected: `4`,
		},
	}

	RunFunctionTest(t, "ROW", testcases, map[string]string{}, [][]string{}, map[string]string{})
}

func TestCOLUMN(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `COLUMN(cell:string):int expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `COLUMN(cell:string):int expects 1 argument, got 2 at input:0:0`,
		},
		{
			Name: "with an int column",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
			},
			Error: `COLUMN(cell:string):int invalid argument 42 at input:0:0`,
		},
		{
			Name: "with an Identifier",
			Input: []ast.Node{
				ast.IdentifierExpression{Value: "B4"},
			},
			Expected: `2`,
		},
		{
			Name: "with a Range",
			Input: []ast.Node{
				ast.RangeExpression{Value: []string{"B4", "C5"}},
			},
			Expected: `2`,
		},
	}

	RunFunctionTest(t, "COLUMN", testcases, map[string]string{}, [][]string{}, map[string]string{})
}

func TestREF(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `REF(address:string):any expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple arguments",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1"},
				ast.StringExpression{Value: "B2"},
			},
			Error: `REF(address:string):any expects 1 argument, got 2 at input:0:0`,
		},
		{
			Name: "with an int",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
			},
			Error: `REF(address:string):any invalid argument 42 at input:0:0`,
		},
		{
			Name: "single cell",
			Input: []ast.Node{
				ast.StringExpression{Value: "B2"},
			},
			Expected: `3`,
		},
		{
			Name: "context variable",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"world"`,
		},
		{
			Name: "cell range",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:B2"},
			},
			Expected: `[A1, B1, A2, B2]`,
		},
		{
			Name: "open cell range",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:A"},
			},
			Expected: `[A1, A2]`,
		},
		{
			Name: "open row range",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:1"},
			},
			Expected: `[A1, B1]`,
		},
		{
			Name: "open sheet range",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:"},
			},
			Expected: `[A1, B1, A2, B2]`,
		},
		{
			Name: "open start range",
			Input: []ast.Node{
				ast.StringExpression{Value: ":B1"},
			},
			Expected: `[A2, B2, A1, B1]`,
		},
		{
			Name: "range with extra cell",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:B2,C3"},
			},
			Expected: `[A1, B1, A2, B2, C3]`,
		},
		{
			Name: "range with spaces",
			Input: []ast.Node{
				ast.StringExpression{Value: " A1 : B2 , C3 "},
			},
			Expected: `[A1, B1, A2, B2, C3]`,
		},
		{
			Name: "empty string",
			Input: []ast.Node{
				ast.StringExpression{Value: ""},
			},
			Error: `REF(address:string):any invalid argument "" at input:0:0`,
		},
		{
			Name: "invalid single cell",
			Input: []ast.Node{
				ast.StringExpression{Value: "1A"},
			},
			Error: `REF(address:string):any invalid argument "1A" at input:0:0`,
		},
		{
			Name: "invalid range start",
			Input: []ast.Node{
				ast.StringExpression{Value: "1A:B2"},
			},
			Error: `cannot expand: invalid range 1A:B2`,
		},
		{
			Name: "invalid range end",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1:2B"},
			},
			Error: `cannot expand: invalid range A1:2B`,
		},
	}

	context := map[string]string{
		"hello": "world",
	}

	input := [][]string{
		{"0", "$1"},
		{"2", "3"},
	}

	formats := map[string]string{"B1": "$%d"}

	RunFunctionTest(t, "REF", testcases, context, input, formats)
}
