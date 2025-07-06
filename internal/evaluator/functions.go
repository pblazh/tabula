package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
)

// SumInts accepts a list of Int AST nodes and returns the sum of their values
func SumInts(nodes []ast.IntExpression) int {
	sum := 0
	for _, node := range nodes {
		sum += node.Value
	}
	return sum
}
