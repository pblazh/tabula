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
			Error: "NOT(value:boolean):boolean expects 1 argument, got 0 in (NOT), at <: input:0:0>",
		},
		{
			Name: "true input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Expected: "<bool false>",
		},
		{
			Name: "false input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: "<bool true>",
		},
		{
			Name: "multiple inputs",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Error: "NOT(value:boolean):boolean expects 1 argument, got 2 in (NOT <bool true> <bool false>), at <: input:0:0>",
		},
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			Error: "NOT(value:boolean):boolean received invalid argument <str \"true\"> in (NOT <str \"true\">), at <: input:0:0>",
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			Error: "NOT(value:boolean):boolean received invalid argument <int 1> in (NOT <int 1>), at <: input:0:0>",
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 1.0},
			},
			Error: "NOT(value:boolean):boolean received invalid argument <float 1.00> in (NOT <float 1.00>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "NOT", testcases)
}

func TestAND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "AND(a:boolean, b:boolean):boolean expects 2 arguments, got 0 in (AND), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "AND(a:boolean, b:boolean):boolean expects 2 arguments, got 1 in (AND <bool true>), at <: input:0:0>",
		},
		{
			Name: "true and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: "<bool false>",
		},
		{
			Name: "true and true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			Expected: "<bool true>",
		},
		{
			Name: "false and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			Expected: "<bool false>",
		},
		{
			Name: "three arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Error: "AND(a:boolean, b:boolean):boolean expects 2 arguments, got 3 in (AND <bool true> <bool false> <bool true>), at <: input:0:0>",
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: "AND(a:boolean, b:boolean):boolean received invalid argument <str \"false\"> in (AND <bool true> <str \"false\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "AND", testcases)
}

func TestOR(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "OR(a:boolean, b:boolean):boolean expects 2 arguments, got 0 in (OR), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "OR(a:boolean, b:boolean):boolean expects 2 arguments, got 1 in (OR <bool true>), at <: input:0:0>",
		},
		{
			Name: "true and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: "<bool true>",
		},
		{
			Name: "true and true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			Expected: "<bool true>",
		},
		{
			Name: "false and false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			Expected: "<bool false>",
		},
		{
			Name: "three arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Error: "OR(a:boolean, b:boolean):boolean expects 2 arguments, got 3 in (OR <bool true> <bool false> <bool true>), at <: input:0:0>",
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: "OR(a:boolean, b:boolean):boolean received invalid argument <str \"false\"> in (OR <bool true> <str \"false\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "OR", testcases)
}

func TestIF(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 0 in (IF), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 1 in (IF <bool true>), at <: input:0:0>",
		},
		{
			Name: "true condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: "<str \"yes\">",
		},
		{
			Name: "false condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: "<str \"no\">",
		},
		{
			Name: "mixed types true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.IntExpression{Value: 42},
				ast.FloatExpression{Value: 3.14},
			},
			Expected: "<int 42>",
		},
		{
			Name: "mixed types false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.IntExpression{Value: 42},
				ast.FloatExpression{Value: 3.14},
			},
			Expected: "<float 3.14>",
		},
		{
			Name: "boolean results",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			Expected: "<bool false>",
		},
		{
			Name: "non-boolean condition",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Error: "IF(predicate:boolean, positive:any, negative:any):any received invalid argument <str \"true\"> in (IF <str \"true\"> <str \"yes\"> <str \"no\">), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
				ast.StringExpression{Value: "extra"},
			},
			Error: "IF(predicate:boolean, positive:any, negative:any):any expects 3 arguments, got 4 in (IF <bool true> <str \"yes\"> <str \"no\"> <str \"extra\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "IF", testcases)
}

func TestTRUE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "no arguments",
			Input:    []ast.Expression{},
			Expected: "<bool true>",
		},
		{
			Name: "with arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "TRUE():boolean expects 0 arguments, got 1 in (TRUE <bool true>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "TRUE", testcases)
}

func TestFALSE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "no arguments",
			Input:    []ast.Expression{},
			Expected: "<bool false>",
		},
		{
			Name: "with arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "FALSE():boolean expects 0 arguments, got 1 in (FALSE <bool true>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "FALSE", testcases)
}
