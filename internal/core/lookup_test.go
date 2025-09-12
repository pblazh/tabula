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
