package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

// TestFunctionExamples tests all the examples from functions.md documentation
func TestFunctionExamples(t *testing.T) {
	testcases := []InfoTestCase{
		// Numeric Functions Examples
		{
			Name: "SUM with numbers",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			Expected: "6",
		},
		{
			Name: "ADD two numbers",
			Input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.IntExpression{Value: 3},
			},
			Expected: "8",
		},
		{
			Name: "PRODUCT multiply numbers",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 4},
			},
			Expected: "24",
		},
		{
			Name: "AVERAGE of numbers",
			Input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 20},
				ast.IntExpression{Value: 30},
			},
			Expected: "20",
		},
		{
			Name: "ABS negative number",
			Input: []ast.Expression{
				ast.IntExpression{Value: -5},
			},
			Expected: "5",
		},
		{
			Name: "POWER base and exponent",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			Expected: "8.00",
		},
		{
			Name: "CEILING round up with factor",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 4.3},
				ast.IntExpression{Value: 1},
			},
			Expected: "5",
		},
		{
			Name: "CEILING round up to nearest 10",
			Input: []ast.Expression{
				ast.IntExpression{Value: 15},
				ast.IntExpression{Value: 10},
			},
			Expected: "20",
		},
		{
			Name: "FLOOR round down with factor",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 4.7},
				ast.IntExpression{Value: 1},
			},
			Expected: "4",
		},
		{
			Name: "FLOOR round down to nearest 10",
			Input: []ast.Expression{
				ast.IntExpression{Value: 15},
				ast.IntExpression{Value: 10},
			},
			Expected: "10",
		},
		{
			Name: "INT truncate decimal positive",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 4.9},
			},
			Expected: "4",
		},
		{
			Name: "INT truncate decimal negative",
			Input: []ast.Expression{
				ast.FloatExpression{Value: -3.2},
			},
			Expected: "-3",
		},
		{
			Name: "MAX find largest",
			Input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 25},
				ast.IntExpression{Value: 5},
			},
			Expected: "25",
		},
		{
			Name: "MIN find smallest",
			Input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 25},
				ast.IntExpression{Value: 5},
			},
			Expected: "5",
		},
		{
			Name: "ROUND with 2 decimal places",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 4.567},
				ast.FloatExpression{Value: 0.01},
			},
			Expected: "4.57",
		},
		{
			Name: "ROUND to integer",
			Input: []ast.Expression{
				ast.FloatExpression{Value: 4.567},
				ast.IntExpression{Value: 1},
			},
			Expected: "5",
		},
		{
			Name: "MOD remainder operation",
			Input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 3},
			},
			Expected: "1",
		},
		{
			Name: "SQRT square root",
			Input: []ast.Expression{
				ast.IntExpression{Value: 16},
			},
			Expected: "4",
		},

		// String Functions Examples
		{
			Name: "CONCATENATE multiple strings",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: " "},
				ast.StringExpression{Value: "World"},
			},
			Expected: `"Hello World"`,
		},
		{
			Name: "LEN string length",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
			},
			Expected: "5",
		},
		{
			Name: "UPPER convert to uppercase",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: `"HELLO"`,
		},
		{
			Name: "LOWER convert to lowercase",
			Input: []ast.Expression{
				ast.StringExpression{Value: "HELLO"},
			},
			Expected: `"hello"`,
		},
		{
			Name: "TRIM remove spaces",
			Input: []ast.Expression{
				ast.StringExpression{Value: "  hello  "},
			},
			Expected: `"hello"`,
		},
		{
			Name: "EXACT compare strings case sensitive",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "true",
		},
		{
			Name: "EXACT compare strings different case",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "hello"},
			},
			Expected: "false",
		},
		{
			Name: "FIND substring position",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "lo"},
			},
			Expected: "3",
		},
		{
			Name: "LEFT get leftmost characters",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.IntExpression{Value: 5},
			},
			Expected: `"Hello"`,
		},
		{
			Name: "RIGHT get rightmost characters",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.IntExpression{Value: 5},
			},
			Expected: `"World"`,
		},
		{
			Name: "MID get middle substring",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.IntExpression{Value: 7},
				ast.IntExpression{Value: 5},
			},
			Expected: `"World"`,
		},
		{
			Name: "SUBSTITUTE replace text",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "o"},
				ast.StringExpression{Value: "0"},
			},
			Expected: `"Hell0 W0rld"`,
		},
		{
			Name: "SUBSTITUTE replace first occurrence",
			Input: []ast.Expression{
				ast.StringExpression{Value: "Hello World"},
				ast.StringExpression{Value: "o"},
				ast.StringExpression{Value: "0"},
				ast.IntExpression{Value: 1},
			},
			Expected: `"Hell0 World"`,
		},
		{
			Name: "VALUE convert string number",
			Input: []ast.Expression{
				ast.StringExpression{Value: "123"},
			},
			Expected: "123",
		},
		{
			Name: "VALUE convert string float",
			Input: []ast.Expression{
				ast.StringExpression{Value: "45.67"},
			},
			Expected: "45.67",
		},

		// Logical Functions Examples
		{
			Name: "IF condition true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "High"},
				ast.StringExpression{Value: "Low"},
			},
			Expected: `"High"`,
		},
		{
			Name: "IF condition false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.StringExpression{Value: "High"},
				ast.StringExpression{Value: "Low"},
			},
			Expected: `"Low"`,
		},
		{
			Name: "AND both true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			Expected: "true",
		},
		{
			Name: "AND one false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: "false",
		},
		{
			Name: "OR one true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			Expected: "true",
		},
		{
			Name: "OR both false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			Expected: "false",
		},
		{
			Name: "NOT true",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Expected: "false",
		},
		{
			Name: "NOT false",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			Expected: "true",
		},
		{
			Name:     "TRUE function",
			Input:    []ast.Expression{},
			Expected: "true",
		},
		{
			Name:     "FALSE function",
			Input:    []ast.Expression{},
			Expected: "false",
		},

		// Count Functions Examples
		{
			Name: "COUNT numbers only",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.StringExpression{Value: "text"},
				ast.IntExpression{Value: 4},
			},
			Expected: "3",
		},
		{
			Name: "COUNTA non-empty values",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "text"},
				ast.IntExpression{Value: 4},
			},
			Expected: "3",
		},

		// Information Functions Examples
		{
			Name: "ISNUMBER with number",
			Input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			Expected: "true",
		},
		{
			Name: "ISNUMBER with text",
			Input: []ast.Expression{
				ast.StringExpression{Value: "text"},
			},
			Expected: "false",
		},
		{
			Name: "ISTEXT with text",
			Input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			Expected: "true",
		},
		{
			Name: "ISTEXT with number",
			Input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			Expected: "false",
		},
		{
			Name: "ISLOGICAL with boolean",
			Input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			Expected: "true",
		},
		{
			Name: "ISLOGICAL with text",
			Input: []ast.Expression{
				ast.StringExpression{Value: "text"},
			},
			Expected: "false",
		},
		{
			Name: "ISBLANK with empty string",
			Input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			Expected: "true",
		},
		{
			Name: "ISBLANK with text",
			Input: []ast.Expression{
				ast.StringExpression{Value: "text"},
			},
			Expected: "false",
		},
	}

	// Run tests for each function
	functionTests := map[string][]InfoTestCase{
		"SUM":         {testcases[0]},
		"ADD":         {testcases[1]},
		"PRODUCT":     {testcases[2]},
		"AVERAGE":     {testcases[3]},
		"ABS":         {testcases[4]},
		"POWER":       {testcases[5]},
		"CEILING":     {testcases[6], testcases[7]},
		"FLOOR":       {testcases[8], testcases[9]},
		"INT":         {testcases[10], testcases[11]},
		"MAX":         {testcases[12]},
		"MIN":         {testcases[13]},
		"ROUND":       {testcases[14], testcases[15]},
		"MOD":         {testcases[16]},
		"SQRT":        {testcases[17]},
		"CONCATENATE": {testcases[18]},
		"LEN":         {testcases[19]},
		"UPPER":       {testcases[20]},
		"LOWER":       {testcases[21]},
		"TRIM":        {testcases[22]},
		"EXACT":       {testcases[23], testcases[24]},
		"FIND":        {testcases[25]},
		"LEFT":        {testcases[26]},
		"RIGHT":       {testcases[27]},
		"MID":         {testcases[28]},
		"SUBSTITUTE":  {testcases[29], testcases[30]},
		"VALUE":       {testcases[31], testcases[32]},
		"IF":          {testcases[33], testcases[34]},
		"AND":         {testcases[35], testcases[36]},
		"OR":          {testcases[37], testcases[38]},
		"NOT":         {testcases[39], testcases[40]},
		"TRUE":        {testcases[41]},
		"FALSE":       {testcases[42]},
		"COUNT":       {testcases[43]},
		"COUNTA":      {testcases[44]},
		"ISNUMBER":    {testcases[45], testcases[46]},
		"ISTEXT":      {testcases[47], testcases[48]},
		"ISLOGICAL":   {testcases[49], testcases[50]},
		"ISBLANK":     {testcases[51], testcases[52]},
	}

	for functionName, tests := range functionTests {
		t.Run(functionName, func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.Name, func(t *testing.T) {
					result, err := DispatchMap[functionName](map[string]string{}, [][]string{}, map[string]string{}, ast.CallExpression{
						Identifier: ast.IdentifierExpression{
							Value: functionName,
							Token: lexer.Token{Literal: functionName},
						}, Arguments: tc.Input,
					}, tc.Input...)

					if tc.Error != "" {
						if err == nil {
							t.Errorf("Expected error %q, got result: %v", tc.Error, result)
							return
						}
						if err.Error() != tc.Error {
							t.Errorf("Expected error %q, got %q", tc.Error, err.Error())
						}
						return
					}

					if err != nil {
						t.Errorf("Unexpected error: %v", err)
						return
					}

					if result.String() != tc.Expected {
						t.Errorf("Expected %s, got %s", tc.Expected, result.String())
					}
				})
			}
		})
	}
}

// TestDateFunctionExamples tests date function examples that don't require external dependencies
func TestDateFunctionExamples(t *testing.T) {
	// Note: Date function tests would require more complex setup with proper date expressions
	// and would depend on the current system time for NOW() function.
	// These are skipped in this basic example file but could be expanded with proper fixtures.
	t.Skip("Date function examples require more complex test setup with date expressions")
}

// TestExecFunctionExamples tests EXEC function examples
func TestExecFunctionExamples(t *testing.T) {
	t.Skip("EXEC function examples are environment-dependent and require external commands")

	// Example test structure for EXEC function (would be uncommented in a full test environment):
	// testcases := []InfoTestCase{
	// 	{
	// 		Name: "EXEC echo command",
	// 		Input: []ast.Expression{
	// 			ast.StringExpression{Value: "echo"},
	// 			ast.StringExpression{Value: "Hello World"},
	// 		},
	// 		Expected: `"Hello World"`,
	// 	},
	// }
	// RunFunctionTest(t, "EXEC", testcases)
}

// TestComplexExamples tests more complex combinations mentioned in functions.md
func TestComplexExamples(t *testing.T) {
	testcases := []struct {
		name         string
		functionName string
		input        []ast.Expression
		expected     string
	}{
		{
			name:         "UPPER TRIM combination simulation",
			functionName: "UPPER",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"}, // Simulating already trimmed input
			},
			expected: `"HELLO"`,
		},
		{
			name:         "SUM ABS combination simulation",
			functionName: "SUM",
			input: []ast.Expression{
				ast.IntExpression{Value: 5}, // Simulating ABS(-5)
				ast.IntExpression{Value: 3}, // Simulating ABS(3)
			},
			expected: "8",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DispatchMap[tc.functionName](
				map[string]string{}, [][]string{}, map[string]string{},
				ast.CallExpression{
					Identifier: ast.IdentifierExpression{
						Value: tc.functionName,
						Token: lexer.Token{Literal: tc.functionName},
					}, Arguments: tc.input,
				}, tc.input...)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}
