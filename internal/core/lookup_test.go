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
			Error: `ADDRESS(row:int, column:int) expected 2 arguments, but got 0 in (ADDRESS), at <: input:0:0>`,
		},
		{
			Name: "single int",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `ADDRESS(row:int, column:int) expected 2 arguments, but got 1 in (ADDRESS <int 42>), at <: input:0:0>`,
		},
		{
			Name: "multiple ints",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			Error: `ADDRESS(row:int, column:int) expected 2 arguments, but got 3 in (ADDRESS <int 1> <int 2> <int 3>), at <: input:0:0>`,
		},
		{
			Name: "with string column",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
				ast.StringExpression{Value: "A"},
			},
			Error: `ADDRESS(row:int, column:int) got a wrong argument <str "A"> in (ADDRESS <int 42> <str "A">), at <: input:0:0>`,
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

func TestCOLUMN(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `COLUMN(cell:string) expected 1 argument, but got 0 in (COLUMN), at <: input:0:0>`,
		},
		{
			Name: "multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
			},
			Error: `COLUMN(cell:string) expected 1 argument, but got 2 in (COLUMN <str "hello"> <str "world">), at <: input:0:0>`,
		},
		{
			Name: "with an int column",
			Input: []ast.Expression{
				ast.IntExpression{Value: 42},
			},
			Error: `COLUMN(cell:string) got a wrong argument <int 42> in (COLUMN <int 42>), at <: input:0:0>`,
		},
		{
			Name: "with an Identifier",
			Input: []ast.Expression{
				ast.IdentifierExpression{Value: "B4"},
			},
			Expected: `<int 2>`,
		},
		{
			Name: "with a Range",
			Input: []ast.Expression{
				ast.RangeExpression{Value: []string{"B4", "C5"}},
			},
			Expected: `<int 2>`,
		},
	}

	RunFunctionTest(t, "COLUMN", testcases)
}
