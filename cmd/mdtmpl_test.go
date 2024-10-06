package cmd

import (
	"strings"
	"testing"

	"github.com/FalcoSuessgott/mdtmpl/pkg/parser"
	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestParseConfig(t *testing.T) {
	testCases := []struct {
		name string
		tmpl string
		exp  string
		err  bool
	}{
		{
			name: "simple",
			tmpl: `<!--- {{ "hello!" | upper | repeat 5 }} --->`,
			exp: `<!--- {{ "hello!" | upper | repeat 5 }} --->
HELLO!HELLO!HELLO!HELLO!HELLO!
`,
		},
		{
			name: "exec",
			tmpl: `<!--- {{ exec "echo hallo" | repeat 3 }} --->`,
			exp: `<!--- {{ exec "echo hallo" | repeat 3 }} --->
hallo
hallo
hallo

`,
		},
		{
			name: "fle",
			tmpl: `<!--- {{ file "testdata/cfg.yml" | code "yml" }} --->`,
			exp: `<!--- {{ file "testdata/cfg.yml" | code "yml" }} --->` + "\n```yml" + `
settings:
  cfg: true

` + "```\n",
		},
		{
			name: "tmpl",
			tmpl: `<!--- {{ tmpl "testdata/tmpl.tmpl" }} --->`,
			exp: `<!--- {{ tmpl "testdata/tmpl.tmpl" }} --->
This is a test template

`,
		},
		{
			name: "tmplWithVars",
			tmpl: `<!--- {{ tmplWithVars "testdata/template.tmpl" "version=v1.0.0" "name=kuberbernetes" }} --->`,
			exp: `<!--- {{ tmplWithVars "testdata/template.tmpl" "version=v1.0.0" "name=kuberbernetes" }} --->
This is another template kuberbernetes-v1.0.0

`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := strings.NewReader(tc.tmpl)
			res, err := parser.Parse(s)

			if tc.err {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.exp, string(res))
		})
	}
}
