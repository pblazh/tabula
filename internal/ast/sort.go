package ast

import (
	"fmt"
	"maps"
	"sort"
)

type DependencyInfo struct {
	Name         string
	Statement    Statement
	Dependencies []string
}

type DependencyGraph struct {
	Nodes    map[string]*DependencyInfo
	AdjList  map[string][]string
	InDegree map[string]int
}

func ExtractDependencies(expr Expression) []string {
	var deps []string
	switch e := expr.(type) {
	case IdentifierExpression:
		deps = append(deps, e.Token.Literal)
	case PrefixExpression:
		deps = append(deps, ExtractDependencies(e.Right)...)
	case InfixExpression:
		deps = append(deps, ExtractDependencies(e.Left)...)
		deps = append(deps, ExtractDependencies(e.Right)...)
	case CallExpression:
		deps = append(deps, ExtractDependencies(e.Identifier)...)
		for _, arg := range e.Arguments {
			deps = append(deps, ExtractDependencies(arg)...)
		}
	default:
		// No dependencies for literal expressions
	}
	return removeDependenciesDuplicates(deps)
}

// GetStatementName returns the name/identifier of a statement if it defines one
func GetStatementName(stmt Statement) string {
	switch s := stmt.(type) {
	case LetStatement:
		return s.Identifier.Token.Literal
	default:
		return ""
	}
}

// GetStatementDependencies returns the dependencies of a statement
func GetStatementDependencies(stmt Statement) []string {
	switch s := stmt.(type) {
	case LetStatement:
		return ExtractDependencies(s.Value)
	case ExpressionStatement:
		return ExtractDependencies(s.Value)
	default:
		return []string{}
	}
}

// removeDependenciesDuplicates removes duplicate strings from a slice in-place
func removeDependenciesDuplicates(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	sort.Strings(slice)

	writeIndex := 1
	for readIndex := 1; readIndex < len(slice); readIndex++ {
		if slice[readIndex] != slice[readIndex-1] {
			slice[writeIndex] = slice[readIndex]
			writeIndex++
		}
	}

	// Return sub-slice containing only unique elements
	return slice[:writeIndex]
}

// NewDependencyGraph creates a new dependency graph from a program
func NewDependencyGraph(program Program) *DependencyGraph {
	graph := &DependencyGraph{
		Nodes:    make(map[string]*DependencyInfo),
		AdjList:  make(map[string][]string),
		InDegree: make(map[string]int),
	}

	// First pass: collect all named statements
	for _, stmt := range program {
		name := GetStatementName(stmt)
		if name != "" {
			deps := GetStatementDependencies(stmt)
			graph.Nodes[name] = &DependencyInfo{
				Name:         name,
				Statement:    stmt,
				Dependencies: deps,
			}
			graph.InDegree[name] = 0 // Initialize in-degree
		}
	}

	// Second pass: build adjacency list and calculate in-degrees
	for name, info := range graph.Nodes {
		for _, dep := range info.Dependencies {
			// Only add edges if the dependency is also a defined variable
			if _, exists := graph.Nodes[dep]; exists {
				graph.AdjList[dep] = append(graph.AdjList[dep], name)
				graph.InDegree[name]++
			}
		}
	}

	return graph
}

// Sort performs Kahn's algorithm for topological sorting
func (g *DependencyGraph) Sort() ([]Statement, error) {
	// Copy in-degrees to avoid modifying original
	inDegree := make(map[string]int)
	maps.Copy(inDegree, g.InDegree)

	// Find all nodes with in-degree 0
	queue := []string{}
	for name, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, name)
		}
	}
	// Sort queue for deterministic ordering
	sort.Strings(queue)

	var result []Statement
	processed := 0

	for len(queue) > 0 {
		// Remove a node with in-degree 0
		current := queue[0]
		queue = queue[1:]
		processed++

		// Add to result
		result = append(result, g.Nodes[current].Statement)

		// For each neighbor of current node
		neighbors := g.AdjList[current]
		sort.Strings(neighbors) // Deterministic processing order
		for _, neighbor := range neighbors {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				// Insert in sorted position to maintain deterministic order
				insertPos := 0
				for insertPos < len(queue) && queue[insertPos] < neighbor {
					insertPos++
				}
				queue = append(queue[:insertPos], append([]string{neighbor}, queue[insertPos:]...)...)
			}
		}
	}

	// Check for cycles
	if processed != len(g.Nodes) {
		return nil, fmt.Errorf("circular dependency detected")
	}

	return result, nil
}

// SortProgram sorts a program's statements in topological order
func SortProgram(program Program) (Program, error) {
	graph := NewDependencyGraph(program)
	sortedStatements, err := graph.Sort()
	if err != nil {
		return nil, err
	}

	// Start with sorted let statements
	var result Program
	result = append(result, sortedStatements...)

	// Add expression statements that don't define variables
	for _, stmt := range program {
		if GetStatementName(stmt) == "" {
			result = append(result, stmt)
		}
	}

	return result, nil
}
