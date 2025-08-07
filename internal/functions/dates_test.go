package functions

import (
	"testing"
	"time"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func parseDate(value string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", value)
	return t
}

func TestDatesParsing(t *testing.T) {
	testcases := []struct {
		name     string
		f        string
		input    []ast.Expression
		expected string
		error    string
	}{
		// parsing
		{
			name: "parse valid input",
			f:    "TODATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2025-08-07"},
			},
			expected: "<date 2025-08-07 00:00:00>",
		},
		{
			name:  "parse empty input",
			f:     "TODATE",
			input: []ast.Expression{},
			error: "TODATE(string, string) expected 2 arguments, but got 0 in (TODATE), at <: input:0:0>",
		},
		{
			name:  "parse too few arguments",
			f:     "TODATE",
			input: []ast.Expression{ast.StringExpression{Value: "2006-01-02"}},
			error: "TODATE(string, string) expected 2 arguments, but got 1 in (TODATE <str \"2006-01-02\">), at <: input:0:0>",
		},
		{
			name: "parse too many arguments",
			f:    "TODATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-01"},
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2006-01-03"},
			},
			error: "TODATE(string, string) expected 2 arguments, but got 3 in (TODATE <str \"2006-01-01\"> <str \"2006-01-02\"> <str \"2006-01-03\">), at <: input:0:0>",
		},
		{
			name: "parse invalid layout",
			f:    "TODATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "not a layout"},
				ast.StringExpression{Value: "2025-08-07"},
			},
			error: "failed TODATE(string, string) with <: input:0:0> at parsing time \"2025-08-07\" as \"not a layout\": cannot parse \"2025-08-07\" as \"not a layout\"",
		},
		{
			name: "parse invalid input",
			f:    "TODATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "not a date"},
			},
			error: "failed TODATE(string, string) with <: input:0:0> at parsing time \"not a date\" as \"2006-01-02\": cannot parse \"not a date\" as \"2006\"",
		},

		// formating
		{
			name: "format valid input",
			f:    "FROMDATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006.01.02"},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<str \"2025.08.07\">",
		},
		{
			name:  "format empty input",
			f:     "FROMDATE",
			input: []ast.Expression{},
			error: "FROMDATE(string, date) expected 2 arguments, but got 0 in (FROMDATE), at <: input:0:0>",
		},
		{
			name: "format too few arguments",
			f:    "FROMDATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
			},
			error: "FROMDATE(string, date) expected 2 arguments, but got 1 in (FROMDATE <str \"2006-01-02\">), at <: input:0:0>",
		},
		{
			name: "format too many arguments",
			f:    "FROMDATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-01"},
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2006-01-03"},
			},
			error: "FROMDATE(string, date) expected 2 arguments, but got 3 in (FROMDATE <str \"2006-01-01\"> <str \"2006-01-02\"> <str \"2006-01-03\">), at <: input:0:0>",
		},
		// values
		{
			name: "day valid input",
			f:    "DAY",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 7>",
		},
		{
			name: "day invalid input",
			f:    "DAY",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "DAY(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (DAY <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "hour valid input",
			f:    "HOUR",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 13>",
		},
		{
			name: "hour invalid input",
			f:    "HOUR",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "HOUR(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (HOUR <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "minute valid input",
			f:    "MINUTE",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 41>",
		},
		{
			name: "minute invalid input",
			f:    "MINUTE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "MINUTE(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (MINUTE <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "month valid input",
			f:    "MONTH",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 8>",
		},
		{
			name: "month invalid input",
			f:    "MONTH",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "MONTH(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (MONTH <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "second valid input",
			f:    "SECOND",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 55>",
		},
		{
			name: "second invalid input",
			f:    "SECOND",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "SECOND(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (SECOND <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "year valid input",
			f:    "YEAR",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 2025>",
		},
		{
			name: "year invalid input",
			f:    "YEAR",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "YEAR(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (YEAR <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		{
			name: "weekday valid input",
			f:    "WEEKDAY",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			expected: "<int 4>",
		},
		{
			name: "weekday invalid input",
			f:    "WEEKDAY",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "WEEKDAY(date) got a wrong argument <str \"2025-08-07 13:41:55\"> in (WEEKDAY <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			funcRef, ok := DispatchMap[tc.f]
			if !ok {
				t.Errorf("Unsupported function: %s", tc.f)
				return
			}

			result, err := funcRef(ast.CallExpression{
				Identifier: ast.IdentifierExpression{
					Value: tc.f,
					Token: lexer.Token{Literal: tc.f},
				}, Arguments: tc.input,
			}, tc.input...)

			if tc.error != "" {
				if err == nil {
					t.Errorf("Expected error %q but got result: %v", tc.error, result)
					return
				}
				if err.Error() != tc.error {
					t.Errorf("Expected error %q, got %q", tc.error, err.Error())
				}
				return
			}

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
