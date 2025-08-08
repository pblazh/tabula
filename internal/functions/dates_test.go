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
		{
			name: "now invalid input",
			f:    "NOW",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			error: "NOW() expected 0 arguments, but got 1 in (NOW <str \"2025-08-07 13:41:55\">), at <: input:0:0>",
		},
		// DATE function tests
		{
			name: "date valid input",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 7},
			},
			expected: "<date 2025-08-07 00:00:00>",
		},
		{
			name:  "date empty input",
			f:     "DATE",
			input: []ast.Expression{},
			error: "DATE(year, month, day) expected 3 arguments, but got 0 in (DATE), at <: input:0:0>",
		},
		{
			name: "date too few arguments",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
			},
			error: "DATE(year, month, day) expected 3 arguments, but got 2 in (DATE <int 2025> <int 8>), at <: input:0:0>",
		},
		{
			name: "date too many arguments",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 7},
				ast.IntExpression{Value: 12},
			},
			error: "DATE(year, month, day) expected 3 arguments, but got 4 in (DATE <int 2025> <int 8> <int 7> <int 12>), at <: input:0:0>",
		},
		{
			name: "date invalid year type",
			f:    "DATE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025"},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 7},
			},
			error: "DATE(year, month, day) got a wrong argument <str \"2025\"> in (DATE <str \"2025\"> <int 8> <int 7>), at <: input:0:0>",
		},
		{
			name: "date invalid month type",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.StringExpression{Value: "8"},
				ast.IntExpression{Value: 7},
			},
			error: "DATE(year, month, day) got a wrong argument <str \"8\"> in (DATE <int 2025> <str \"8\"> <int 7>), at <: input:0:0>",
		},
		{
			name: "date invalid day type",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
				ast.StringExpression{Value: "7"},
			},
			error: "DATE(year, month, day) got a wrong argument <str \"7\"> in (DATE <int 2025> <int 8> <str \"7\">), at <: input:0:0>",
		},
		{
			name: "date with edge case values",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 2000},
				ast.IntExpression{Value: 12},
				ast.IntExpression{Value: 31},
			},
			expected: "<date 2000-12-31 00:00:00>",
		},
		{
			name: "date with minimum values",
			f:    "DATE",
			input: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 1},
			},
			expected: "<date 0001-01-01 00:00:00>",
		},
		// DATEDIF function tests
		{
			name: "datedif valid years",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2023-01-01 00:00:00")},
				ast.StringExpression{Value: "Y"},
			},
			expected: "<int 3>",
		},
		{
			name: "datedif valid months",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-06-01 00:00:00")},
				ast.StringExpression{Value: "M"},
			},
			expected: "<int 5>",
		},
		{
			name: "datedif valid days",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-10 00:00:00")},
				ast.StringExpression{Value: "D"},
			},
			expected: "<int 9>",
		},
		{
			name: "datedif valid days ignoring months and years (MD)",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-15 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-02-20 00:00:00")},
				ast.StringExpression{Value: "MD"},
			},
			expected: "<int 5>",
		},
		{
			name: "datedif valid months ignoring years (YM)",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-03-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-06-01 00:00:00")},
				ast.StringExpression{Value: "YM"},
			},
			expected: "<int 3>",
		},
		{
			name: "datedif valid days ignoring years (YD)",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-02-15 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-03-20 00:00:00")},
				ast.StringExpression{Value: "YD"},
			},
			expected: "<int 5>",
		},
		{
			name:  "datedif empty input",
			f:     "DATEDIF",
			input: []ast.Expression{},
			error: "DATEDIF(from, to, unit) expected 3 arguments, but got 0 in (DATEDIF), at <: input:0:0>",
		},
		{
			name: "datedif too few arguments",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
			},
			error: "DATEDIF(from, to, unit) expected 3 arguments, but got 2 in (DATEDIF <date 2020-01-01 00:00:00> <date 2021-01-01 00:00:00>), at <: input:0:0>",
		},
		{
			name: "datedif too many arguments",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
				ast.StringExpression{Value: "Y"},
				ast.StringExpression{Value: "extra"},
			},
			error: "DATEDIF(from, to, unit) expected 3 arguments, but got 4 in (DATEDIF <date 2020-01-01 00:00:00> <date 2021-01-01 00:00:00> <str \"Y\"> <str \"extra\">), at <: input:0:0>",
		},
		{
			name: "datedif invalid from type",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.StringExpression{Value: "2020-01-01"},
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
				ast.StringExpression{Value: "Y"},
			},
			error: "DATEDIF(from, to, unit) got a wrong argument <str \"2020-01-01\"> in (DATEDIF <str \"2020-01-01\"> <date 2021-01-01 00:00:00> <str \"Y\">), at <: input:0:0>",
		},
		{
			name: "datedif invalid to type",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.StringExpression{Value: "2021-01-01"},
				ast.StringExpression{Value: "Y"},
			},
			error: "DATEDIF(from, to, unit) got a wrong argument <str \"2021-01-01\"> in (DATEDIF <date 2020-01-01 00:00:00> <str \"2021-01-01\"> <str \"Y\">), at <: input:0:0>",
		},
		{
			name: "datedif invalid unit type",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
				ast.IntExpression{Value: 1},
			},
			error: "DATEDIF(from, to, unit) got a wrong argument <int 1> in (DATEDIF <date 2020-01-01 00:00:00> <date 2021-01-01 00:00:00> <int 1>), at <: input:0:0>",
		},
		{
			name: "datedif invalid unit value",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
				ast.StringExpression{Value: "X"},
			},
			error: "DATEDIF(from, to, unit) got a wrong argument <str \"X\"> in (DATEDIF <date 2020-01-01 00:00:00> <date 2021-01-01 00:00:00> <str \"X\">), at <: input:0:0>",
		},
		{
			name: "days valid days",
			f:    "DAYS",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-10 00:00:00")},
			},
			expected: "<int 9>",
		},
		// DATEDIF negative distance test cases (start > end)
		{
			name: "datedif negative years",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2023-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.StringExpression{Value: "Y"},
			},
			expected: "<int -3>",
		},
		{
			name: "datedif negative months",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-06-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.StringExpression{Value: "M"},
			},
			expected: "<int -5>",
		},
		{
			name: "datedif negative days",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-10 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
				ast.StringExpression{Value: "D"},
			},
			expected: "<int -9>",
		},
		{
			name: "datedif negative days MD",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2021-02-20 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-15 00:00:00")},
				ast.StringExpression{Value: "MD"},
			},
			expected: "<int -5>",
		},
		{
			name: "datedif negative months YM",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2021-06-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-03-01 00:00:00")},
				ast.StringExpression{Value: "YM"},
			},
			expected: "<int -3>",
		},
		{
			name: "datedif negative days YD",
			f:    "DATEDIF",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2021-03-20 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-02-15 00:00:00")},
				ast.StringExpression{Value: "YD"},
			},
			expected: "<int -5>",
		},
		// DAYS negative distance test cases
		{
			name: "days negative days",
			f:    "DAYS",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2020-01-10 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
			},
			expected: "<int -9>",
		},
		{
			name: "days negative across years",
			f:    "DAYS",
			input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2021-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2020-01-01 00:00:00")},
			},
			expected: "<int -366>",
		},
		// DATEVALUE function test cases
		{
			name: "datevalue iso date format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07"},
			},
			expected: "<date 2025-08-07 00:00:00>",
		},
		{
			name: "datevalue iso datetime format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 15:30"},
			},
			expected: "<date 2025-08-07 15:30:00>",
		},
		{
			name: "datevalue european date format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "07.08.2025"},
			},
			expected: "<date 2025-08-07 00:00:00>",
		},
		{
			name: "datevalue european datetime format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "07.08.2025 15:30"},
			},
			expected: "<date 2025-08-07 15:30:00>",
		},
		{
			name: "datevalue european datetime with seconds",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "07.08.2025 15:30:45"},
			},
			expected: "<date 2025-08-07 15:30:45>",
		},
		{
			name: "datevalue us date format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "08/07/2025"},
			},
			expected: "<date 2025-08-07 00:00:00>",
		},
		{
			name: "datevalue us datetime format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "08/07/2025 15:30"},
			},
			expected: "<date 2025-08-07 15:30:00>",
		},
		{
			name: "datevalue us datetime with seconds",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "08/07/2025 15:30:45"},
			},
			expected: "<date 2025-08-07 15:30:45>",
		},
		{
			name: "datevalue datetime format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07 15:30:45"},
			},
			expected: "<date 2025-08-07 15:30:45>",
		},
		{
			name: "datevalue time only format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "15:04:05"},
			},
			expected: "<date 0000-01-01 15:04:05>",
		},
		{
			name: "datevalue kitchen time format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "3:04PM"},
			},
			expected: "<date 0000-01-01 15:04:00>",
		},
		{
			name:  "datevalue empty input",
			f:     "DATEVALUE",
			input: []ast.Expression{},
			error: "DATEVALUE(string) expected 1 argument, but got 0 in (DATEVALUE), at <: input:0:0>",
		},
		{
			name: "datevalue too many arguments",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-07"},
				ast.StringExpression{Value: "extra"},
			},
			error: "DATEVALUE(string) expected 1 argument, but got 2 in (DATEVALUE <str \"2025-08-07\"> <str \"extra\">), at <: input:0:0>",
		},
		{
			name: "datevalue invalid argument type",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.IntExpression{Value: 20250807},
			},
			error: "DATEVALUE(string) got a wrong argument <int 20250807> in (DATEVALUE <int 20250807>), at <: input:0:0>",
		},
		{
			name: "datevalue invalid date string",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "not a date"},
			},
			error: "failed DATEVALUE(string) with <: input:0:0> at <nil>",
		},
		{
			name: "datevalue empty string",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			error: "failed DATEVALUE(string) with <: input:0:0> at <nil>",
		},
		{
			name: "datevalue invalid date format",
			f:    "DATEVALUE",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025/13/45"},
			},
			error: "failed DATEVALUE(string) with <: input:0:0> at <nil>",
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
