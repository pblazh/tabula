package markdown

import (
	"strings"
	"sync"
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
			name:  "CSV malformed",
			input: "```csv\none,two\n1,2,3\n#tabula let x = 0;\n```\n",
			output: "```csv\none,two\n1,2,3\n#tabula let x = 0;\n```\n" +
				"<!-- Tabula: cannot read CSV: record on line 2: wrong number of fields -->\n",
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
				"#tabula let x = 0;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"```csv",
				"one , two , three",
				"1   , 2   , 3",
				"#tabula let x = 0;",
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
			name: "CSV with fixed code",
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"<!-- Tabula: message to be removed -->",
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
				"<!-- Tabula: can not parse: unexpected ; at input:1:10 -->",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "CSV with failing again code",
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"<!-- Tabula: previous error -->",
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
				"<!-- Tabula: can not parse: unexpected ; at input:1:10 -->",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table one cell",
			input: strings.Join([]string{
				"",
				"| one |",
				"| - |",
				"| 1 |",
				"",
				"```tabula",
				"let x = 0;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"",
				"| one |",
				"| --- |",
				"| 1   |",
				"",
				"```tabula",
				"let x = 0;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table malformed header",
			input: strings.Join([]string{
				"",
				"| one |",
				"| - | - |",
				"| 1 |",
				"",
				"```tabula",
				"let x = 0;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"",
				"| one |",
				"| - | - |",
				"| 1 |",
				"<!-- Tabula: malformed table header -->",
				"",
				"```tabula",
				"let x = 0;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table malformed body",
			input: strings.Join([]string{
				"",
				"| one |",
				"| - |",
				"| 1 | 2 |",
				"",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"",
				"| one |",
				"| - |",
				"| 1 | 2 |",
				"<!-- Tabula: malformed table at line 1 -->",
				"",
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
| - | - | - |
| 1 | 2 | 3 |
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
			name: "table with previously failed code",
			input: strings.Join([]string{
				"",
				"| one | two | three |",
				"| - | - | - |",
				"| 1 | 2 | 3 |",
				"<!-- Tabula:  previouse error -->",
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
				"<!-- Tabula: can not parse: unexpected ; at input:1:10 -->",
				"",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "table with failing again code",
			input: strings.Join([]string{
				"",
				"| one | two | three |",
				"| - | - | - |",
				"| 1 | 2 | 3 |",
				"<!-- Tabula:  previouse error -->",
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
				"<!-- Tabula: can not parse: unexpected ; at input:1:10 -->",
				"```tabula",
				"let A2 = ;",
				"```",
				"",
			}, "\n"),
		},
		{
			name: "multiple CSV with code concurrent",
			input: strings.Join([]string{
				"```csv",
				"one,two,three",
				"1,2,3",
				"```",
				"```tabula",
				"let A2 = 10;",
				"```",
				"",
				"```csv",
				"one,two,three",
				"4,5,6",
				"```",
				"```tabula",
				"let A2 = 20;",
				"```",
				"",
			}, "\n"),
			output: strings.Join([]string{
				"```csv",
				"one,two,three",
				"10,2,3",
				"```",
				"```tabula",
				"let A2 = 10;",
				"```",
				"",
				"```csv",
				"one,two,three",
				"20,5,6",
				"```",
				"```tabula",
				"let A2 = 20;",
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

// TestProcessConcurrentChunks verifies that multiple processable chunks in a
// single document are each executed with their own correct script and data.
// This guards against closure variable capture bugs where goroutines share
// loop variables and all end up processing the last chunk's data.
func TestProcessConcurrentChunks(t *testing.T) {
	input := strings.Join([]string{
		"```csv",
		"one,two,three",
		"1,2,3",
		"```",
		"```tabula",
		"let A2 = 10;",
		"```",
		"",
		"```csv",
		"one,two,three",
		"4,5,6",
		"```",
		"```tabula",
		"let A2 = 20;",
		"```",
		"",
	}, "\n")

	want := strings.Join([]string{
		"```csv",
		"one,two,three",
		"10,2,3",
		"```",
		"```tabula",
		"let A2 = 10;",
		"```",
		"",
		"```csv",
		"one,two,three",
		"20,5,6",
		"```",
		"```tabula",
		"let A2 = 20;",
		"```",
		"",
	}, "\n")

	const iterations = 100
	var wg sync.WaitGroup
	wg.Add(iterations)
	for range iterations {
		go func() {
			defer wg.Done()
			var writer strings.Builder
			if err := Process(&Config{}, strings.NewReader(input), &writer); err != nil {
				t.Errorf("unexpected error: %s", err)
				return
			}
			if got := writer.String(); got != want {
				t.Errorf("concurrent output mismatch:\nwant: %q\n got: %q", want, got)
			}
		}()
	}
	wg.Wait()
}
