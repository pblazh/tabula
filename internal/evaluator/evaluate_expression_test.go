package evaluator

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/testutil"
)

func TestInfixExpressionEvaluate(t *testing.T) {
	testcases := []struct {
		name    string
		input   string
		expects string
	}{
		// Values
		{
			name:    "string",
			input:   `"hello"`,
			expects: "\"hello\"",
		},
		{
			name:    "int",
			input:   `9`,
			expects: "9",
		},
		{
			name:    "float",
			input:   `9.01`,
			expects: "9.01",
		},
		{
			name:    "bool",
			input:   `true`,
			expects: "true",
		},
		// Addition
		{
			name:    "int + int",
			input:   `5 + 3`,
			expects: "8",
		},
		{
			name:    "int + float",
			input:   `5 + 3.5`,
			expects: "8.50",
		},
		{
			name:    "float + int",
			input:   `5.5 + 3`,
			expects: "8.50",
		},
		{
			name:    "float + float",
			input:   `5.5 + 3.2`,
			expects: "8.70",
		},
		// Subtraction
		{
			name:    "int - int",
			input:   `10 - 3`,
			expects: "7",
		},
		{
			name:    "int - float",
			input:   `10 - 3.5`,
			expects: "6.50",
		},
		{
			name:    "float - int",
			input:   `10.5 - 3`,
			expects: "7.50",
		},
		{
			name:    "float - float",
			input:   `10.5 - 3.2`,
			expects: "7.30",
		},
		// Multiplication
		{
			name:    "int * int",
			input:   `5 * 3`,
			expects: "15",
		},
		{
			name:    "int * float",
			input:   `5 * 2.5`,
			expects: "12.50",
		},
		{
			name:    "float * int",
			input:   `4.5 * 2`,
			expects: "9.00",
		},
		{
			name:    "float * float",
			input:   `2.5 * 3.0`,
			expects: "7.50",
		},
		// Division
		{
			name:    "int / int",
			input:   `10 / 2`,
			expects: "5",
		},
		{
			name:    "int / float",
			input:   `10 / 2.5`,
			expects: "4.00",
		},
		{
			name:    "float / int",
			input:   `10.5 / 2`,
			expects: "5.25",
		},
		{
			name:    "float / float",
			input:   `10.5 / 2.5`,
			expects: "4.20",
		},
		// Equality
		{
			name:    "int == int (true)",
			input:   `5 == 5`,
			expects: "true",
		},
		{
			name:    "int == int (false)",
			input:   `5 == 3`,
			expects: "false",
		},
		{
			name:    "int == float (true)",
			input:   `5 == 5.0`,
			expects: "true",
		},
		{
			name:    "int == float (false)",
			input:   `5 == 5.01`,
			expects: "false",
		},
		{
			name:    "float == int (true)",
			input:   `5.0 == 5`,
			expects: "true",
		},
		{
			name:    "float == int (false)",
			input:   `5.01 == 5`,
			expects: "false",
		},
		{
			name:    "float == float (true)",
			input:   `3.14 == 3.14`,
			expects: "true",
		},
		{
			name:    "float == float (false)",
			input:   `3.14 == 2.71`,
			expects: "false",
		},
		{
			name:    "bool == bool (true)",
			input:   `true == true`,
			expects: "true",
		},
		{
			name:    "bool == bool (false)",
			input:   `true == false`,
			expects: "false",
		},
		// Inequality
		{
			name:    "int != int (true)",
			input:   `5 != 3`,
			expects: "true",
		},
		{
			name:    "int != int (false)",
			input:   `5 != 5`,
			expects: "false",
		},
		{
			name:    "int != float (true)",
			input:   `5 != 3.5`,
			expects: "true",
		},
		{
			name:    "int != float (false)",
			input:   `5 != 5.0`,
			expects: "false",
		},
		{
			name:    "float != int (true)",
			input:   `3.5 != 5`,
			expects: "true",
		},
		{
			name:    "float != int (false)",
			input:   `5.0 != 5`,
			expects: "false",
		},
		{
			name:    "float != float (true)",
			input:   `3.14 != 2.71`,
			expects: "true",
		},
		{
			name:    "float != float (false)",
			input:   `3.14 != 3.14`,
			expects: "false",
		},
		{
			name:    "bool != bool (true)",
			input:   `true != false`,
			expects: "true",
		},
		{
			name:    "bool != bool (false)",
			input:   `true != true`,
			expects: "false",
		},
		// Less Than Comparison
		{
			name:    "int < int (true)",
			input:   `3 < 5`,
			expects: "true",
		},
		{
			name:    "int < int (false)",
			input:   `5 < 3`,
			expects: "false",
		},
		{
			name:    "int < int (equal)",
			input:   `5 < 5`,
			expects: "false",
		},
		{
			name:    "int < float (true)",
			input:   `3 < 5.5`,
			expects: "true",
		},
		{
			name:    "int < float (false)",
			input:   `5 < 3.5`,
			expects: "false",
		},
		{
			name:    "float < int (true)",
			input:   `3.5 < 5`,
			expects: "true",
		},
		{
			name:    "float < int (false)",
			input:   `5.5 < 3`,
			expects: "false",
		},
		{
			name:    "float < float (true)",
			input:   `3.2 < 5.7`,
			expects: "true",
		},
		{
			name:    "float < float (false)",
			input:   `5.7 < 3.2`,
			expects: "false",
		},
		// Greater Than Comparison
		{
			name:    "int > int (true)",
			input:   `5 > 3`,
			expects: "true",
		},
		{
			name:    "int > int (false)",
			input:   `3 > 5`,
			expects: "false",
		},
		{
			name:    "int > int (equal)",
			input:   `5 > 5`,
			expects: "false",
		},
		{
			name:    "int > float (true)",
			input:   `5 > 3.5`,
			expects: "true",
		},
		{
			name:    "int > float (false)",
			input:   `3 > 5.5`,
			expects: "false",
		},
		{
			name:    "float > int (true)",
			input:   `5.5 > 3`,
			expects: "true",
		},
		{
			name:    "float > int (false)",
			input:   `3.5 > 5`,
			expects: "false",
		},
		{
			name:    "float > float (true)",
			input:   `5.7 > 3.2`,
			expects: "true",
		},
		{
			name:    "float > float (false)",
			input:   `3.2 > 5.7`,
			expects: "false",
		},
		// Function calls
		{
			name:    "SUM with integers",
			input:   `SUM(1, 2, 3)`,
			expects: "6",
		},
		{
			name:    "SUM with floats",
			input:   `SUM(1.5, 2.5)`,
			expects: "4.00",
		},
		{
			name:    "SUM with mixed int and float (int first)",
			input:   `SUM(5, 2.5)`,
			expects: "7.50",
		},
		{
			name:    "SUM with mixed float and int (float first)",
			input:   `SUM(2.5, 5)`,
			expects: "7.50",
		},
		{
			name:    "SUM with single argument",
			input:   `SUM(42)`,
			expects: "42",
		},
		{
			name:    "SUM with no arguments",
			input:   `SUM()`,
			expects: "0",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			var input [][]string
			result, err := EvaluateExpression(expr, make(map[string]string), input, make(map[string]string), "target")
			if err != nil {
				t.Errorf("Unexpects error: %v", err)
				return
			}

			if result.String() != tc.expects {
				t.Errorf("Expected %s, got %s", tc.expects, result.String())
			}
		})
	}
}

func TestPrefixExpressionEvaluate(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expects  string
		hasError bool
	}{
		{
			name:     "negation of int",
			input:    `-5`,
			expects:  "-5",
			hasError: false,
		},
		{
			name:     "negation of float",
			input:    `-3.5`,
			expects:  "-3.50",
			hasError: false,
		},
		{
			name:     "negation of negative int",
			input:    `-(-5)`,
			expects:  "5",
			hasError: false,
		},
		{
			name:     "logical not of true",
			input:    `!true`,
			expects:  "false",
			hasError: false,
		},
		{
			name:     "logical not of false",
			input:    `!false`,
			expects:  "true",
			hasError: false,
		},
		{
			name:    "addition and multiplication precedence",
			input:   `2 + 3 * 4`,
			expects: "14",
		},
		{
			name:    "multiplication and addition precedence",
			input:   `3 * 4 + 2`,
			expects: "14",
		},
		{
			name:    "parentheses override precedence",
			input:   `(2 + 3) * 4`,
			expects: "20",
		},
		{
			name:    "complex arithmetic",
			input:   `10 - 2 * 3 + 1`,
			expects: "5",
		},
		{
			name:    "division and multiplication",
			input:   `12 / 3 * 2`,
			expects: "8",
		},
		{
			name:    "comparison with arithmetic",
			input:   `5 + 3 > 7`,
			expects: "true",
		},
		{
			name:    "equality with arithmetic",
			input:   `2 * 3 == 6`,
			expects: "true",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			var input [][]string
			result, err := EvaluateExpression(expr, make(map[string]string), input, make(map[string]string), "target")
			if err != nil {
				t.Errorf("Unexpects error: %v", err)
				return
			}

			if result.String() != tc.expects {
				t.Errorf("Expected %s, got %s", tc.expects, result.String())
			}
		})
	}
}

func TestOperationErrors(t *testing.T) {
	testcases := []struct {
		name         string
		input        string
		expectsError string
	}{
		// Division by zero errors
		{
			name:         "int division by zero",
			input:        `10 / 0`,
			expectsError: `division by zero at <DIV:/ test:1:4>`,
		},
		{
			name:         "float division by zero",
			input:        `10.5 / 0.0`,
			expectsError: `division by zero at <DIV:/ test:1:6>`,
		},
		{
			name:         "int division by float zero",
			input:        `10 / 0.0`,
			expectsError: `division by zero at <DIV:/ test:1:4>`,
		},
		{
			name:         "float division by int zero",
			input:        `10.5 / 0`,
			expectsError: `division by zero at <DIV:/ test:1:6>`,
		},
		// Type mismatch errors for arithmetic operations
		{
			name:         "bool + int",
			input:        `true + 5`,
			expectsError: `operator <PLUS:+ test:1:6> is not supported for type: boolean and integer`,
		},
		{
			name:         "int + bool",
			input:        `5 + true`,
			expectsError: `operator <PLUS:+ test:1:3> is not supported for type: integer and boolean`,
		},
		{
			name:         "bool - int",
			input:        `true - 5`,
			expectsError: `operator <MINUS:- test:1:6> is not supported for type: boolean and integer`,
		},
		{
			name:         "bool * int",
			input:        `true * 5`,
			expectsError: `operator <MULT:* test:1:6> is not supported for type: boolean and integer`,
		},
		{
			name:         "bool / int",
			input:        `true / 5`,
			expectsError: `operator <DIV:/ test:1:6> is not supported for type: boolean and integer`,
		},
		// Type mismatch errors for comparison operations
		{
			name:         "int < bool",
			input:        `5 < true`,
			expectsError: `operator <LESS:< test:1:3> is not supported for type: integer and boolean`,
		},
		{
			name:         "bool > int",
			input:        `true > 5`,
			expectsError: `operator <GREATER:> test:1:6> is not supported for type: boolean and integer`,
		},
		// Prefix operation errors
		{
			name:         "negation of bool",
			input:        `-true`,
			expectsError: `<MINUS:- test:1:1> is not supported for type: boolean`,
		},
		{
			name:         "logical not of int",
			input:        `!5`,
			expectsError: `<NOT:! test:1:1> is not supported for type: integer`,
		},
		{
			name:         "logical not of float",
			input:        `!3.14`,
			expectsError: `<NOT:! test:1:1> is not supported for type: float`,
		},
		// Function call errors
		{
			name:         "SUM with unsupported boolean first argument",
			input:        `SUM(true)`,
			expectsError: `SUM(values:number...):number received invalid argument true in SUM(true), at <IDENT:SUM test:1:1>`,
		},
		{
			name:         "SUM with mixed incompatible types in integer sum",
			input:        `SUM(5, "hello")`,
			expectsError: `SUM(values:number...):number received invalid argument "hello" in SUM(5, "hello"), at <IDENT:SUM test:1:1>`,
		},
		{
			name:         "SUM with mixed incompatible types in float sum",
			input:        `SUM(5.5, true)`,
			expectsError: `SUM(values:number...):number received invalid argument true in SUM(5.50, true), at <IDENT:SUM test:1:1>`,
		},
		{
			name:         "SUM with mixed incompatible types in string sum",
			input:        `SUM("hello", 42)`,
			expectsError: `SUM(values:number...):number received invalid argument "hello" in SUM("hello", 42), at <IDENT:SUM test:1:1>`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)
			var input [][]string
			formats := make(map[string]string)
			result, err := EvaluateExpression(expr, make(map[string]string), input, formats, "target")
			if err == nil {
				t.Errorf("Expected error, got result: %s", result.String())
				return
			}

			if err.Error() != tc.expectsError {
				t.Errorf("Expected error %q, got %q", tc.expectsError, err.Error())
			}
		})
	}
}

func TestRangeExpressionTokenPreservation(t *testing.T) {
	testcases := []struct {
		name       string
		input      string
		expectsPos string
	}{
		{
			name:       "simple horizontal range A1:C1",
			input:      `A1:C1`,
			expectsPos: "test:1:3",
		},
		{
			name:       "simple vertical range A1:A3",
			input:      `A1:A3`,
			expectsPos: "test:1:3",
		},
		{
			name:       "larger range A1:C3",
			input:      `A1:C3`,
			expectsPos: "test:1:3",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			// Verify it's a RangeExpression
			rangeExpr, ok := expr.(ast.RangeExpression)
			if !ok {
				t.Fatalf("Expected RangeExpression, got %T", expr)
			}

			// Check the Token position of the range expression itself
			if rangeExpr.Token.Position.String() != tc.expectsPos {
				t.Errorf("Expected range token position %s, got %s", tc.expectsPos, rangeExpr.Token.Position.String())
			}

			// The key test: verify that when we expand the range expression,
			// the generated IdentifierExpression objects have the same Token as the original range
			// This simulates what EvaluateRangeExpression does internally
			cells := make([]ast.IdentifierExpression, len(rangeExpr.Value))
			for i, cell := range rangeExpr.Value {
				cells[i] = ast.IdentifierExpression{Token: rangeExpr.Token, Value: cell}
			}

			// Verify each generated IdentifierExpression has the correct Token
			for i, cell := range cells {
				if cell.Token.Position.String() != tc.expectsPos {
					t.Errorf("Generated identifier %d (%s) has wrong token position expects %s, got %s",
						i, cell.Value, tc.expectsPos, cell.Token.Position.String())
				}

				if cell.Token.Type != rangeExpr.Token.Type {
					t.Errorf("Generated identifier %d (%s) has wrong token type expects %v, got %v",
						i, cell.Value, rangeExpr.Token.Type, cell.Token.Type)
				}
			}
		})
	}
}
