package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
)

func TestADDRESS(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 0 in ADDRESS(), at <: input:0:0>`,
		},
		{
			Name: "single int",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 1 in ADDRESS(42), at <: input:0:0>`,
			// ADDRESS(row:int, column:int):string expects 2 arguments, got 1 in ADDRESS(42), at <: input:0:0>
		},
		{
			Name: "multiple ints",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			Error: `ADDRESS(row:int, column:int):string expects 2 arguments, got 3 in ADDRESS(1, 2, 3), at <: input:0:0>`,
		},
		{
			Name: "with string column",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
				ast.StringExpression{Value: "A"},
			},
			Error: `ADDRESS(row:int, column:int):string received an invalid argument "A" in ADDRESS(42, "A"), at <: input:0:0>`,
		},
		{
			Name: "happy path",
			Input: []ast.Expression{
				ast.IntExpression{Value: 4},
				ast.IntExpression{Value: 2},
			},
			Expected: `B4`,
		},
	}

	RunFunctionTest(t, "ADDRESS", testcases)
}

func TestROW(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `ROW(cell:string):int expects 1 argument, got 0 in ROW(), at <: input:0:0>`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `ROW(cell:string):int expects 1 argument, got 2 in ROW("hello", "world"), at <: input:0:0>`,
		},
		{
			Name: "with an int column",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `ROW(cell:string):int received an invalid argument 42 in ROW(42), at <: input:0:0>`,
		},
		{
			Name: "with an Identifier",
			Input: []ast.Expression{
				ast.IdentifierExpression{Value: "B4"},
			},
			Expected: `4`,
		},
		{
			Name: "with a Range",
			Input: []ast.Expression{
				ast.RangeExpression{Value: []string{"B4", "C5"}},
			},
			Expected: `4`,
		},
	}

	RunFunctionTest(t, "ROW", testcases)
}

func TestCOLUMN(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `COLUMN(cell:string):int expects 1 argument, got 0 in COLUMN(), at <: input:0:0>`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `COLUMN(cell:string):int expects 1 argument, got 2 in COLUMN("hello", "world"), at <: input:0:0>`,
		},
		{
			Name: "with an int column",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `COLUMN(cell:string):int received an invalid argument 42 in COLUMN(42), at <: input:0:0>`,
		},
		{
			Name: "with an Identifier",
			Input: []ast.Expression{
				ast.IdentifierExpression{Value: "B4"},
			},
			Expected: `2`,
		},
		{
			Name: "with a Range",
			Input: []ast.Expression{
				ast.RangeExpression{Value: []string{"B4", "C5"}},
			},
			Expected: `2`,
		},
	}

	RunFunctionTest(t, "COLUMN", testcases)
}
