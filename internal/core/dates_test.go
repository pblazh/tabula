package core

import (
	"testing"
	"time"

	"github.com/pblazh/tabula/internal/ast"
)

func parseDate(value string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", value)
	return t
}

func TestTODATE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2025-08-07"},
			},
			Expected: "<2025-08-07 00:00:00>",
 		},
 		{
 			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "TODATE(layout:string, value:string):date expects 2 arguments, got 0 in TODATE(), at <: input:0:0>",
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
			},
			Error: "TODATE(layout:string, value:string):date expects 2 arguments, got 1 in TODATE(\"2006-01-02\"), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-01"},
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2006-01-03"},
			},
			Error: "TODATE(layout:string, value:string):date expects 2 arguments, got 3 in TODATE(\"2006-01-01\", \"2006-01-02\", \"2006-01-03\"), at <: input:0:0>",
		},
		{
			Name: "invalid layout",
			Input: []ast.Expression{
				ast.StringExpression{Value: "not a layout"},
				ast.StringExpression{Value: "2025-08-07"},
			},
			Error: "failed TODATE(layout:string, value:string):date with <: input:0:0> at parsing time \"2025-08-07\" as \"not a layout\": cannot parse \"2025-08-07\" as \"not a layout\"",
		},
		{
			Name: "invalid input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "not a date"},
			},
			Error: "failed TODATE(layout:string, value:string):date with <: input:0:0> at parsing time \"not a date\" as \"2006-01-02\": cannot parse \"not a date\" as \"2006\"",
		},
		{
			Name: "with time format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02 15:04:05"},
				ast.StringExpression{Value: "2025-08-07 13:41:55"},
			},
			Expected: "<2025-08-07 13:41:55>",
		},
		{
			Name: "different date format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "01/02/2006"},
				ast.StringExpression{Value: "08/07/2025"},
			},
			Expected: "<2025-08-07 00:00:00>",
		},
	}

	RunFunctionTest(t, "TODATE", testcases)
}

func TestFROMDATE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid input",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006.01.02"},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			Expected: "\"2025.08.07\"",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "FROMDATE(layout:string, source:date):string expects 2 arguments, got 0 in FROMDATE(), at <: input:0:0>",
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
			},
			Error: "FROMDATE(layout:string, source:date):string expects 2 arguments, got 1 in FROMDATE(\"2006-01-02\"), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-01"},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
				ast.StringExpression{Value: "extra"},
			},
			Error: "FROMDATE(layout:string, source:date):string expects 2 arguments, got 3 in FROMDATE(\"2006-01-01\", <2025-08-07 13:41:55>, \"extra\"), at <: input:0:0>",
		},
		{
			Name: "with time format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02 15:04:05"},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			Expected: "\"2025-08-07 13:41:55\"",
		},
		{
			Name: "different output format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "01/02/2006"},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			Expected: "\"08/07/2025\"",
		},
		{
			Name: "wrong first argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2006},
				ast.DateExpression{Value: parseDate("2025-08-07 13:41:55")},
			},
			Error: "FROMDATE(layout:string, source:date):string received invalid argument 2006 in FROMDATE(2006, <2025-08-07 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong second argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2006-01-02"},
				ast.StringExpression{Value: "2025-08-07"},
			},
			Error: "FROMDATE(layout:string, source:date):string received invalid argument \"2025-08-07\" in FROMDATE(\"2006-01-02\", \"2025-08-07\"), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "FROMDATE", testcases)
}

func TestDAY(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "17",
		},
		{
			Name: "first day of month",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-01 00:00:00")},
			},
			Expected: "1",
		},
		{
			Name: "last day of month",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-31 23:59:59")},
			},
			Expected: "31",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "DAY(value:date):number expects 1 argument, got 0 in DAY(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-18 13:41:55")},
			},
			Error: "DAY(value:date):number expects 1 argument, got 2 in DAY(<2025-08-17 13:41:55>, <2025-08-18 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17"},
			},
			Error: "DAY(value:date):number received invalid argument \"2025-08-17\" in DAY(\"2025-08-17\"), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "DAY", testcases)
}

func TestMONTH(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "8",
		},
		{
			Name: "january",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-01-15 00:00:00")},
			},
			Expected: "1",
		},
		{
			Name: "december",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-12-25 23:59:59")},
			},
			Expected: "12",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "MONTH(value:date):number expects 1 argument, got 0 in MONTH(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-18 13:41:55")},
			},
			Error: "MONTH(value:date):number expects 1 argument, got 2 in MONTH(<2025-08-17 13:41:55>, <2025-08-18 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 8},
			},
			Error: "MONTH(value:date):number received invalid argument 8 in MONTH(8), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "MONTH", testcases)
}

func TestYEAR(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "2025",
		},
		{
			Name: "different year",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("1999-12-31 23:59:59")},
			},
			Expected: "1999",
		},
		{
			Name: "leap year",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2024-02-29 12:00:00")},
			},
			Expected: "2024",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "YEAR(value:date):number expects 1 argument, got 0 in YEAR(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2024-08-17 13:41:55")},
			},
			Error: "YEAR(value:date):number expects 1 argument, got 2 in YEAR(<2025-08-17 13:41:55>, <2024-08-17 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025"},
			},
			Error: "YEAR(value:date):number received invalid argument \"2025\" in YEAR(\"2025\"), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "YEAR", testcases)
}

func TestHOUR(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date with hour",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "13",
		},
		{
			Name: "midnight",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 00:00:00")},
			},
			Expected: "0",
		},
		{
			Name: "late evening",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 23:59:59")},
			},
			Expected: "23",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "HOUR(value:date):number expects 1 argument, got 0 in HOUR(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 14:41:55")},
			},
			Error: "HOUR(value:date):number expects 1 argument, got 2 in HOUR(<2025-08-17 13:41:55>, <2025-08-17 14:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 13},
			},
			Error: "HOUR(value:date):number received invalid argument 13 in HOUR(13), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "HOUR", testcases)
}

func TestMINUTE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date with minute",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "41",
		},
		{
			Name: "zero minutes",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:00:55")},
			},
			Expected: "0",
		},
		{
			Name: "max minutes",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:59:55")},
			},
			Expected: "59",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "MINUTE(value:date):number expects 1 argument, got 0 in MINUTE(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 13:42:55")},
			},
			Error: "MINUTE(value:date):number expects 1 argument, got 2 in MINUTE(<2025-08-17 13:41:55>, <2025-08-17 13:42:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 41},
			},
			Error: "MINUTE(value:date):number received invalid argument 41 in MINUTE(41), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "MINUTE", testcases)
}

func TestSECOND(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date with second",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "55",
		},
		{
			Name: "zero seconds",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:00")},
			},
			Expected: "0",
		},
		{
			Name: "max seconds",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:59")},
			},
			Expected: "59",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "SECOND(value:date):number expects 1 argument, got 0 in SECOND(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:56")},
			},
			Error: "SECOND(value:date):number expects 1 argument, got 2 in SECOND(<2025-08-17 13:41:55>, <2025-08-17 13:41:56>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 55},
			},
			Error: "SECOND(value:date):number received invalid argument 55 in SECOND(55), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "SECOND", testcases)
}

func TestWEEKDAY(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "sunday",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")}, // Sunday
			},
			Expected: "0",
		},
		{
			Name: "monday",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-18 13:41:55")}, // Monday
			},
			Expected: "1",
		},
		{
			Name: "saturday",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-16 13:41:55")}, // Saturday
			},
			Expected: "6",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "WEEKDAY(value:date):number expects 1 argument, got 0 in WEEKDAY(), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-18 13:41:55")},
			},
			Error: "WEEKDAY(value:date):number expects 1 argument, got 2 in WEEKDAY(<2025-08-17 13:41:55>, <2025-08-18 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			Error: "WEEKDAY(value:date):number received invalid argument 1 in WEEKDAY(1), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "WEEKDAY", testcases)
}

func TestNOW(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name:     "no arguments",
			Input:    []ast.Expression{},
			Expected: "",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "extra"},
			},
			Error: "NOW():date expects 0 arguments, got 1 in NOW(\"extra\"), at <: input:0:0>",
		},
	}

	// Special handling for NOW since we can't predict the exact time
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Error != "" {
				RunFunctionTest(t, "NOW", []InfoTestCase{tc})
				return
			}

			// For successful cases, we need to verify it returns a date type
			call := ast.CallExpression{
				Identifier: ast.IdentifierExpression{Value: "NOW"},
			}
			result, err := DispatchMap["NOW"](call, tc.Input...)
			if err != nil {
				t.Errorf("Unexpects error: %v", err)
				return
			}

			if _, ok := result.(ast.DateExpression); !ok {
				t.Errorf("Expected DateExpression, got %T", result)
			}
		})
	}
}

func TestDATE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "valid date",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 17},
			},
			Expected: "<2025-08-17 00:00:00>",
		},
		{
			Name: "leap year february",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2024},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 29},
			},
			Expected: "<2024-02-29 00:00:00>",
		},
		{
			Name: "first day of year",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 1},
			},
			Expected: "<2025-01-01 00:00:00>",
		},
		{
			Name: "last day of year",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 12},
				ast.IntExpression{Value: 31},
			},
			Expected: "<2025-12-31 00:00:00>",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "DATE(year:number, month:number, day:number):date expects 3 arguments, got 0 in DATE(), at <: input:0:0>",
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
			},
			Error: "DATE(year:number, month:number, day:number):date expects 3 arguments, got 2 in DATE(2025, 8), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 17},
				ast.IntExpression{Value: 12},
			},
			Error: "DATE(year:number, month:number, day:number):date expects 3 arguments, got 4 in DATE(2025, 8, 17, 12), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025"},
				ast.IntExpression{Value: 8},
				ast.IntExpression{Value: 17},
			},
			Error: "DATE(year:number, month:number, day:number):date received invalid argument \"2025\" in DATE(\"2025\", 8, 17), at <: input:0:0>",
		},
		{
			Name: "invalid date",
			Input: []ast.Expression{
				ast.IntExpression{Value: 2025},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 30}, // February 30th doesn't exist
			},
			Expected: "<2025-03-02 00:00:00>", // Go time.Date normalizes invalid dates
		},
	}

	RunFunctionTest(t, "DATE", testcases)
}

func TestDATEDIF(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "days difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
				ast.StringExpression{Value: "D"},
			},
			Expected: "3",
		},
		{
			Name: "months difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-11-17 13:41:55")},
				ast.StringExpression{Value: "M"},
			},
			Expected: "3",
		},
		{
			Name: "years difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2027-08-17 13:41:55")},
				ast.StringExpression{Value: "Y"},
			},
			Expected: "2",
		},
		{
			Name: "negative difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.StringExpression{Value: "D"},
			},
			Expected: "-3",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "DATEDIF(start:date, end:date, unit:string):number expects 3 arguments, got 0 in DATEDIF(), at <: input:0:0>",
		},
		{
			Name: "invalid unit",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
				ast.StringExpression{Value: "X"},
            },
            Error: "DATEDIF(start:date, end:date, unit:string):number received invalid argument \"X\" in DATEDIF(<2025-08-17 13:41:55>, <2025-08-20 13:41:55>, \"X\"), at <: input:0:0>",
        },
        {
            Name: "wrong argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17"},
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
				ast.StringExpression{Value: "D"},
			},
			Error: "DATEDIF(start:date, end:date, unit:string):number received invalid argument \"2025-08-17\" in DATEDIF(\"2025-08-17\", <2025-08-20 13:41:55>, \"D\"), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "DATEDIF", testcases)
}

func TestDAYS(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "positive difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "-3",
		},
		{
			Name: "negative difference",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
			},
			Expected: "3",
		},
		{
			Name: "same date",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Expected: "0",
		},
		{
			Name: "cross year boundary",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2026-01-01 00:00:00")},
				ast.DateExpression{Value: parseDate("2025-12-31 23:59:59")},
			},
			Expected: "0",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "DAYS(start:date, end:date):number expects 2 arguments, got 0 in DAYS(), at <: input:0:0>",
		},
		{
			Name: "too few arguments",
			Input: []ast.Expression{
				ast.DateExpression{Value: parseDate("2025-08-17 13:41:55")},
			},
			Error: "DAYS(start:date, end:date):number expects 2 arguments, got 1 in DAYS(<2025-08-17 13:41:55>), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17"},
				ast.DateExpression{Value: parseDate("2025-08-20 13:41:55")},
			},
			Error: "DAYS(start:date, end:date):number received invalid argument \"2025-08-17\" in DAYS(\"2025-08-17\", <2025-08-20 13:41:55>), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "DAYS", testcases)
}

func TestDATEVALUE(t *testing.T) {
	testcases := []InfoTestCase{
		{
			Name: "ISO date format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17"},
			},
			Expected: "<2025-08-17 00:00:00>",
		},
		{
			Name: "US date format",
			Input: []ast.Expression{
				ast.StringExpression{Value: "08/17/2025"},
			},
			Expected: "<2025-08-17 00:00:00>",
		},
		{
			Name: "with time",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 13:41:55"},
			},
			Expected: "<2025-08-17 13:41:55>",
		},
		{
			Name:  "empty input",
			Input: []ast.Expression{},
			Error: "DATEVALUE(value:string):date expects 1 argument, got 0 in DATEVALUE(), at <: input:0:0>",
		},
		{
			Name: "wrong argument type",
			Input: []ast.Expression{
				ast.IntExpression{Value: 20250817},
			},
			Error: "DATEVALUE(value:string):date received invalid argument 20250817 in DATEVALUE(20250817), at <: input:0:0>",
		},
		{
			Name: "too many arguments",
			Input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17"},
				ast.StringExpression{Value: "extra"},
			},
			Error: "DATEVALUE(value:string):date expects 1 argument, got 2 in DATEVALUE(\"2025-08-17\", \"extra\"), at <: input:0:0>",
		},
	}

	RunFunctionTest(t, "DATEVALUE", testcases)
}
