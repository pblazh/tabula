package markdown

import (
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		name   string
		config Config
		input  string
		output string
		error  string
	}{
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "CSV no code",
			input:  "```csv\none,two,three\n1,2,3\n```\n",
			output: "```csv\none,two,three\n1,2,3\n```\n",
		},
		{
			name:   "CSV aligned",
			config: Config{Align: true},
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"```csv",
				"one , two , three",
				"1   , 2   , 3",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "CSV with code",
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"```tabula",
				"let A2 = 99;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"```csv",
				"one,two,three",
				"99,2,3",
				"```",
				"```tabula",
				"let A2 = 99;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "CSV with failing code",
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"<!-- tabula error parsing script, failed to parse program , unexpected ; at <input>:1:10 -->",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table no code",
			input: `
| one | two | three |
| - | - | - |
| 1 | 2 | 3 |
`,
			output: `
| one | two | three |
| --- | --- | ----- |
| 1   | 2   | 3     |
`,
		},
		{
			name: "table with code",
			input: strings.Join([]string{
				"",
				"| one | two | three |",
				"| - | - | - |",
				"| 1 | 2 | 3 |",
				"",
				"```tabula",
				"let A2 = 99;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"",
				"| one | two | three |",
				"| --- | --- | ----- |",
				"| 99  | 2   | 3     |",
				"",
				"```tabula",
				"let A2 = 99;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table with failing code",
			input: strings.Join([]string{
				"",
				"| one | two | three |",
				"| - | - | - |",
				"| 1 | 2 | 3 |",
				"",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"",
				"| one | two | three |",
				"| - | - | - |",
				"| 1 | 2 | 3 |",
				"<!-- tabula error parsing script, failed to parse program , unexpected ; at <input>:1:10 -->",
				"",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var writer strings.Builder

			err := Process(&tc.config, reader, &writer)
			output := writer.String()

			if tc.error == "" && err != nil {
				t.Errorf("Unexpected error %s", err)
			}
			if tc.error != "" && err == nil {
				t.Errorf("Expected error %s", err)
			}
			if tc.output != output {
				t.Errorf("Expected '%v', but got '%v'", tc.output, output)
			}
		})
	}
}
