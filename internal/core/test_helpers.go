package core

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

// InfoTestCase defines the structure for info function test cases
type InfoTestCase struct {
	Name     string
	Input    []ast.Expression
	Expected string
	Error    string
}

// RunFunctionTest executes test cases for info functions (ISNUMBER, ISTEXT, ISLOGICAL, ISBLANK)
func RunFunctionTest(t *testing.T, functionName string, testcases []InfoTestCase) {
	t.Helper()
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := DispatchMap[functionName](ast.CallExpression{
				Identifier: ast.IdentifierExpression{
					Value: functionName,
					Token: lexer.Token{Literal: functionName},
				}, Arguments: tc.Input,
			}, tc.Input...)

			if tc.Error != "" {
				if err == nil {
					t.Errorf("Expected error %q but got result: %v", tc.Error, result)
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
}

