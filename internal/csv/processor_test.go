package csv

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
			name:   "malformed",
			input:  "one,two\n1,2,3\n4\n",
			output: "",
			error:  "tabula failed to read CSV, record on line 2: wrong number of fields",
		},
		{
			name:   "single value",
			input:  "one\n",
			output: "one\n",
		},
		{
			name:   "no code",
			input:  "one,two,three\n1,2,3\n",
			output: "one,two,three\n1,2,3\n",
		},
		{
			name:   "assign",
			config: Config{Execute: `let A1="ONE"`},
			input:  "one,two,three\n1,2,3\n",
			output: "ONE,two,three\n1,2,3\n",
		},
		{
			name:   "wrong code",
			config: Config{Execute: `let A1=`},
			input:  "one,two,three\n1,2,3\n",
			output: "",
			error:  "failed to parse program, unexpected  at <input>:1:8",
		},
		{
			name:   "comments",
			input:  "#first\none,two,three\n#second\n1,2,3\n#third\n",
			output: "#first\none,two,three\n#second\n1,2,3\n#third\n",
		},
		{
			name:   "code comments",
			input:  "one,two,three\n1,2,3\n#tabula let A1=\"ONE\"\n",
			output: "ONE,two,three\n1,2,3\n#tabula let A1=\"ONE\"\n",
		},
		{
			name:   "open range expression",
			config: Config{Execute: `let C1 = SUM(A1:A);`},
			input:  "1,4\n2,5\n3,6\n",
			output: "1,4,6\n2,5,\n3,6,\n",
		},
		{
			name:   "open row range expression",
			config: Config{Execute: `let D1 = SUM(A1:1);`},
			input:  "1,2,3\n4,5,6",
			output: "1,2,3,6\n4,5,6,\n",
		},
		{
			name:   "open sheet range expression",
			config: Config{Execute: `let D1 = SUM(A1:);`},
			input:  "1,2,3\n4,5,6\n",
			output: "1,2,3,21\n4,5,6,\n",
		},
		{
			name:   "open start range expression",
			config: Config{Execute: `let D1 = SUM(:C1);`},
			input:  "1,2,3\n4,5,6\n",
			output: "1,2,3,21\n4,5,6,\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			var writer strings.Builder

			result, comments, err := Process(&tc.config, reader)

			if tc.error == "" && err != nil {
				t.Errorf("Unexpected error %s", err)
			}
			if tc.error != "" && err == nil {
				t.Errorf("Expected error %s", err)
			}

			err = Write(&writer, result, comments)
			if err != nil {
				t.Errorf("Unexpected error %s", err)
			}

			output := writer.String()

			if tc.output != output {
				t.Errorf("Expected '%v', but got '%v'", tc.output, output)
			}
		})
	}
}
