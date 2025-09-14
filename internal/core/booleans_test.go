package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
)

func TestNOT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `NOT(value:boolean):boolean expects 1 argument, got 0 in NOT(), at <: input:0:0>`,
		},
		{
			Name: "true input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Expected: `false`,
		},
		{
			Name: "false input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: `true`,
		},
		{
			Name: "multiple inputs",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Error: `NOT(value:boolean):boolean expects 1 argument, got 2 in NOT(true, false), at <: input:0:0>`,
		},
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			Error: `NOT(value:boolean):boolean received invalid argument "true" in NOT("true"), at <: input:0:0>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			Error: `NOT(value:boolean):boolean received invalid argument 1 in NOT(1), at <: input:0:0>`,
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 1.0},
			},
			Error: `NOT(value:boolean):boolean received invalid argument 1.00 in NOT(1.00), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "NOT", testcases)
}

func TestAND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `AND(a:boolean, b:boolean):boolean expects 2 arguments, got 0 in AND(), at <: input:0:0>`,
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: `AND(a:boolean, b:boolean):boolean expects 2 arguments, got 1 in AND(true), at <: input:0:0>`,
		},
		{
			Name: "true and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "true and true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			Expected: `true`,
		},
		{
			Name: "false and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "three arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Error: `AND(a:boolean, b:boolean):boolean expects 2 arguments, got 3 in AND(true, false, true), at <: input:0:0>`,
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: `AND(a:boolean, b:boolean):boolean received invalid argument "false" in AND(true, "false"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "AND", testcases)
}

func TestOR(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `OR(a:boolean, b:boolean):boolean expects 2 arguments, got 0 in OR(), at <: input:0:0>`,
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: `OR(a:boolean, b:boolean):boolean expects 2 arguments, got 1 in OR(true), at <: input:0:0>`,
		},
		{
			Name: "true and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: `true`,
		},
		{
			Name: "true and true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			Expected: `true`,
		},
		{
			Name: "false and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "three arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Error: `OR(a:boolean, b:boolean):boolean expects 2 arguments, got 3 in OR(true, false, true), at <: input:0:0>`,
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: `OR(a:boolean, b:boolean):boolean received invalid argument "false" in OR(true, "false"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "OR", testcases)
}

func TestIF(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: `IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 0 in IF(), at <: input:0:0>`,
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: `IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 1 in IF(true), at <: input:0:0>`,
		},
		{
			Name: "true condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: `"yes"`,
		},
		{
			Name: "false condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: `"no"`,
		},
		{
			Name: "mixed types true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.IntExpression{Value: 42},
				ast.FloatExpression{Value: 3.14},
			},
			Expected: `42`,
		},
		{
			Name: "mixed types false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.IntExpression{Value: 42},
				ast.FloatExpression{Value: 3.14},
			},
			Expected: `3.14`,
		},
		{
			Name: "boolean results",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Expected: `false`,
		},
		{
			Name: "non-boolean condition",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Error: `IF(predicate:boolean, positive:any, negative:any):any received invalid argument "true" in IF("true", "yes", "no"), at <: input:0:0>`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
				ast.StringExpression{Value: "extra"},
			},
			Error: `IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 4 in IF(true, "yes", "no", "extra"), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "IF", testcases)
}

func TestTRUE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "no arguments",
			Input:    []ast.Expression{},
			Expected: `true`,
		},
		{
			Name: "with arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: `TRUE():boolean expects 0 arguments, got 1 in TRUE(true), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "TRUE", testcases)
}

func TestFALSE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "no arguments",
			Input:    []ast.Expression{},
			Expected: `false`,
		},
		{
			Name: "with arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: `FALSE():boolean expects 0 arguments, got 1 in FALSE(true), at <: input:0:0>`,
		},
	}

	RunFunctionTest(t, "FALSE", testcases)
}
