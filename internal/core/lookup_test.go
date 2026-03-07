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
			Error: `REF(cell:string):any expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Node{
				ast.StringExpression{Value: "A1"},
				ast.StringExpression{Value: "B2"},
			},
			Error: `REF(cell:string):any expects 1 argument, got 2 at input:0:0`,
		},
		{
			Name: "with an int",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
			},
			Error: `REF(cell:string):any invalid argument 42 at input:0:0`,
		},
		{
			Name: "with a cell Identifier",
			Input: []ast.Node{
				ast.StringExpression{Value: "B2"},
			},
			Expected: `3`,
		},
		{
			Name: "with a formated cell Identifier",
			Input: []ast.Node{
				ast.StringExpression{Value: "B1"},
			},
			Expected: `1`,
		},
		{
			Name: "with a variable Identifier",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"world"`,
		},
		{
			Name: "with a wrong identifier",
			Input: []ast.Node{
				ast.StringExpression{Value: "2B"},
			},
			Error: `REF(cell:string):any invalid argument "2B" at input:0:0`,
		},
	}

	input := [][]string{
		{"0", "$1"},
		{"2", "3"},
	}
	context := map[string]string{
		"hello": "world",
	}
	formats := map[string]string{
		"B1": "$%d",
	}

	RunFunctionTest(t, "REF", testcases, context, input, formats)
}

func TestRANGE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `RANGE(a:string, b:string):range expects 2 arguments, got 0 at input:0:0`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Error: `RANGE(a:string, b:string):range expects 2 arguments, got 1 at input:0:0`,
		},
		{
			Name: "with an int column",
			Input: []ast.Node{
				ast.IntExpression{Value: 42},
				ast.IntExpression{Value: 24},
			},
			Error: `RANGE(a:string, b:string):range invalid argument 42 at input:0:0`,
		},
		{
			Name: "with a variable",
			Input: []ast.Node{
				ast.StringExpression{Value: "x"},
				ast.StringExpression{Value: "C5"},
			},
			Error: `RANGE(a:string, b:string):range invalid argument "x" at input:0:0`,
		},
		{
			Name: "with a Range",
			Input: []ast.Node{
				ast.StringExpression{Value: "B4"},
				ast.StringExpression{Value: "C5"},
			},
			Expected: `[B4, C4, B5, C5]`,
		},
	}

	RunFunctionTest(t, "RANGE", testcases, map[string]string{}, [][]string{}, map[string]string{})
}
