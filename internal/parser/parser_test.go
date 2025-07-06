package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "simple assign",
			input:  "let A1 = -10;",
			output: "let A1 = (- <int 10>);",
		},
		{
			name:  "identifier",
			input: "something;",
		},
		{
			name:   "not identifier",
			input:  "!something;",
			output: "(! something);",
		},
		{
			name:   "number",
			input:  "10;",
			output: "<int 10>;",
		},
		{
			name:   "number",
			input:  "-10;",
			output: "(- <int 10>);",
		},
		{
			name:   "true",
			input:  "true;",
			output: "<bool true>;",
		},
		{
			name:   "false",
			input:  "false;",
			output: "<bool false>;",
		},
		{
			name:   "compare with boolen",
			input:  "3 > 5 == false;",
			output: "(== (> <int 3> <int 5>) <bool false>);",
		},
		{
			name:   "compare boolen",
			input:  "true == 3 < 5;",
			output: "(== <bool true> (< <int 3> <int 5>));",
		},
		{
			name:   "infix",
			input:  "5 + 6 - 2;",
			output: "(- (+ <int 5> <int 6>) <int 2>);",
		},
		{
			name:   "infix reverse precedence",
			input:  "5 + 6 * 2;",
			output: "(+ <int 5> (* <int 6> <int 2>));",
		},
		{
			name:   "infix precedence",
			input:  "5 / 6 + 2;",
			output: "(+ (/ <int 5> <int 6>) <int 2>);",
		},
		{
			name:   "multiple statements",
			input:  "let A1 = 5.6;\nlet A2 = x;\n",
			output: "let A1 = <float 5.60>;let A2 = x;",
		},
		{
			name:   "parenteces precedence",
			input:  "(5 + 6) * 2;",
			output: "(* (+ <int 5> <int 6>) <int 2>);",
		},
		{
			name:   "call expression with no arguments",
			input:  "SUM();",
			output: "(SUM);",
		},
		{
			name:   "call expression with one argument",
			input:  "SUM(5);",
			output: "(SUM <int 5>);",
		},
		{
			name:   "call expression with multiple arguments",
			input:  "SUM(5, 6);",
			output: "(SUM <int 5> <int 6>);",
		},
		{
			name:   "cell expression",
			input:  "a1;",
			output: "A1;",
		},
		{
			name:   "range expression",
			input:  "a1:B1;",
			output: "(: A1 B1);",
		},
		{
			name:   "string expression",
			input:  "let a = \"hello\";",
			output: "let a = <str \"hello\">;",
		},
		{
			name:   "multiple statements",
			input:  "let A1 = 5.6;\nlet A2 = x;\nlet A3 = sum(A1:A2);",
			output: "let A1 = <float 5.60>;let A2 = x;let A3 = (sum (: A1 A2));",
		},
		{
			name:   "one statements",
			input:  "A1",
			output: "A1;",
		},
		{
			name:   "no statements",
			input:  "",
			output: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			program, _, err := parser.Parse()
			if err != nil {
				t.Errorf("Unexpected error '%v'", err)
			}

			literal := ""
			for _, statement := range program {
				literal += statement.String()
			}

			output := tc.output
			if output == "" {
				output = tc.input
			}

			if literal != output {
				t.Errorf("Expected '%s' to equal '%s'", literal, output)
			}
		})
	}
}
