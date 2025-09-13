package core

import (
	"testing"
	"time"

	"github.com/pblazh/tabula/internal/ast"
)

func TestISNUMBER(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "empty string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 7},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "boolean input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "date input",
			Input: []ast.Expression{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `<bool false>`,
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "ISNUMBER(value:any):boolean expects 1 argument, got 0 in (ISNUMBER), at <: input:0:0>",
		},
		{
			Name: "multiple values",
			Input: []ast.Expression{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: "ISNUMBER(value:any):boolean expects 1 argument, got 2 in (ISNUMBER <str \"test\"> <int 39>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "ISNUMBER", testcases)
}

func TestISTEXT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "empty string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 7},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "boolean input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "date input",
			Input: []ast.Expression{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `<bool false>`,
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "ISTEXT(value:any):boolean expects 1 argument, got 0 in (ISTEXT), at <: input:0:0>",
		},
		{
			Name: "multiple values",
			Input: []ast.Expression{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: "ISTEXT(value:any):boolean expects 1 argument, got 2 in (ISTEXT <str \"test\"> <int 39>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "ISTEXT", testcases)
}

func TestISLOGICAL(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "empty string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 7},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "boolean input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "date input",
			Input: []ast.Expression{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `<bool false>`,
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "ISLOGICAL(value:any):boolean expects 1 argument, got 0 in (ISLOGICAL), at <: input:0:0>",
		},
		{
			Name: "multiple values",
			Input: []ast.Expression{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: "ISLOGICAL(value:any):boolean expects 1 argument, got 2 in (ISLOGICAL <str \"test\"> <int 39>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "ISLOGICAL", testcases)
}

func TestISBLANK(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "empty string input",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: `<bool true>`,
		},
		{
			Name: "integer input",
			Input: []ast.Expression{
				ast.IntExpression{Value: 7},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "float input",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "boolean input",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: `<bool false>`,
		},
		{
			Name: "date input",
			Input: []ast.Expression{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `<bool false>`,
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "ISBLANK(value:any):boolean expects 1 argument, got 0 in (ISBLANK), at <: input:0:0>",
		},
		{
			Name: "multiple values",
			Input: []ast.Expression{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: "ISBLANK(value:any):boolean expects 1 argument, got 2 in (ISBLANK <str \"test\"> <int 39>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "ISBLANK", testcases)
}
