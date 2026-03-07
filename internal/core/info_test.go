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
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `false`,
		},
		{
			Name: "empty string input",
			Input: []ast.Node{
				ast.StringExpression{Value: ""},
			},
			Expected: `false`,
		},
		{
			Name: "integer input",
			Input: []ast.Node{
				ast.IntExpression{Value: 7},
			},
			Expected: `true`,
		},
		{
			Name: "float input",
			Input: []ast.Node{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `true`,
		},
		{
			Name: "boolean input",
			Input: []ast.Node{
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "date input",
			Input: []ast.Node{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `false`,
		},
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ISNUMBER(value:any):boolean expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple values",
			Input: []ast.Node{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: `ISNUMBER(value:any):boolean expects 1 argument, got 2 at input:0:0`,
		},
	}

	RunFunctionTest(
		t,
		"ISNUMBER",
		testcases,
		map[string]string{},
		[][]string{},
		map[string]string{},
	)
}

func TestISTEXT(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `true`,
		},
		{
			Name: "empty string input",
			Input: []ast.Node{
				ast.StringExpression{Value: ""},
			},
			Expected: `true`,
		},
		{
			Name: "integer input",
			Input: []ast.Node{
				ast.IntExpression{Value: 7},
			},
			Expected: `false`,
		},
		{
			Name: "float input",
			Input: []ast.Node{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `false`,
		},
		{
			Name: "boolean input",
			Input: []ast.Node{
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "date input",
			Input: []ast.Node{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `false`,
		},
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ISTEXT(value:any):boolean expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple values",
			Input: []ast.Node{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: `ISTEXT(value:any):boolean expects 1 argument, got 2 at input:0:0`,
		},
	}

	RunFunctionTest(t, "ISTEXT", testcases, map[string]string{}, [][]string{}, map[string]string{})
}

func TestISLOGICAL(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `false`,
		},
		{
			Name: "empty string input",
			Input: []ast.Node{
				ast.StringExpression{Value: ""},
			},
			Expected: `false`,
		},
		{
			Name: "integer input",
			Input: []ast.Node{
				ast.IntExpression{Value: 7},
			},
			Expected: `false`,
		},
		{
			Name: "float input",
			Input: []ast.Node{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `false`,
		},
		{
			Name: "boolean input",
			Input: []ast.Node{
				ast.BooleanExpression{Value: false},
			},
			Expected: `true`,
		},
		{
			Name: "date input",
			Input: []ast.Node{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `false`,
		},
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ISLOGICAL(value:any):boolean expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple values",
			Input: []ast.Node{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: `ISLOGICAL(value:any):boolean expects 1 argument, got 2 at input:0:0`,
		},
	}

	RunFunctionTest(
		t,
		"ISLOGICAL",
		testcases,
		map[string]string{},
		[][]string{},
		map[string]string{},
	)
}

func TestISBLANK(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "string input",
			Input: []ast.Node{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `false`,
		},
		{
			Name: "empty string input",
			Input: []ast.Node{
				ast.StringExpression{Value: ""},
			},
			Expected: `true`,
		},
		{
			Name: "integer input",
			Input: []ast.Node{
				ast.IntExpression{Value: 7},
			},
			Expected: `false`,
		},
		{
			Name: "float input",
			Input: []ast.Node{
				ast.FloatExpression{Value: 7.4},
			},
			Expected: `false`,
		},
		{
			Name: "boolean input",
			Input: []ast.Node{
				ast.BooleanExpression{Value: false},
			},
			Expected: `false`,
		},
		{
			Name: "date input",
			Input: []ast.Node{
				ast.DateExpression{Value: time.Now()},
			},
			Expected: `false`,
		},
		{
			Name:  "empty input",
			Input: []ast.Node{},
			Error: `ISBLANK(value:any):boolean expects 1 argument, got 0 at input:0:0`,
		},
		{
			Name: "multiple values",
			Input: []ast.Node{
				ast.StringExpression{Value: "test"},
				ast.IntExpression{Value: 39},
			},
			Error: `ISBLANK(value:any):boolean expects 1 argument, got 2 at input:0:0`,
		},
	}

	RunFunctionTest(t, "ISBLANK", testcases, map[string]string{}, [][]string{}, map[string]string{})
}
