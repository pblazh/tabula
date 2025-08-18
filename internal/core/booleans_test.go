package core

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
)

func TestNOT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "NOT(boolean) expected 1 argument, but got 0 in (NOT), at <: input:0:0>",
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
			Error: "NOT(boolean) expected 1 argument, but got 2 in (NOT <bool true> <bool false>), at <: input:0:0>",
		},
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			Error: "NOT(boolean) got a wrong argument <str \"true\"> in (NOT <str \"true\">), at <: input:0:0>",
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			Error: "NOT(boolean) got a wrong argument <int 1> in (NOT <int 1>), at <: input:0:0>",
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 1.0},
			},
			Error: "NOT(boolean) got a wrong argument <float 1.00> in (NOT <float 1.00>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "NOT", testcases)
}

func TestAND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "AND(boolean, boolean) expected 2 arguments, but got 0 in (AND), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <bool true>), at <: input:0:0>",
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
			Error: "AND(boolean, boolean) expected 2 arguments, but got 3 in (AND <bool true> <bool false> <bool true>), at <: input:0:0>",
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: "AND(boolean, boolean) got a wrong argument <str \"false\"> in (AND <bool true> <str \"false\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "AND", testcases)
}

func TestOR(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "OR(boolean, boolean) expected 2 arguments, but got 0 in (OR), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <bool true>), at <: input:0:0>",
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
			Error: "OR(boolean, boolean) expected 2 arguments, but got 3 in (OR <bool true> <bool false> <bool true>), at <: input:0:0>",
		},
		{
			Name: "boolean and string",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			Error: "OR(boolean, boolean) got a wrong argument <str \"false\"> in (OR <bool true> <str \"false\">), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "OR", testcases)
}

func TestIF(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "IF(boolean, any, any) expected 3 arguments, but got 0 in (IF), at <: input:0:0>",
		},
		{
			Name: "single input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Error: "IF(boolean, any, any) expected 3 arguments, but got 1 in (IF <bool true>), at <: input:0:0>",
		},
		{
			Name: "true condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: `<str "yes">`,
		},
		{
			Name: "false condition",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
			},
			Expected: `<str "no">`,
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
			Error: `IF(boolean, any, any) got a wrong argument <str "true"> in (IF <str "true"> <str "yes"> <str "no">), at <: input:0:0>`,
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "yes"},
				ast.StringExpression{Value: "no"},
				ast.StringExpression{Value: "extra"},
			},
			Error: `IF(boolean, any, any) expected 3 arguments, but got 4 in (IF <bool true> <str "yes"> <str "no"> <str "extra">), at <: input:0:0>`,
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
			Error: "TRUE() expected 0 arguments, but got 1 in (TRUE <bool true>), at <: input:0:0>",
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
			Error: "FALSE() expected 0 arguments, but got 1 in (FALSE <bool true>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "FALSE", testcases)
}