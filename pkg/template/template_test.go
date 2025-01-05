package template

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// nolint: funlen
func TestTemplateFuncMap(t *testing.T) {
	testcases := []struct {
		name string
		opts []RendererOptions
		vars interface{}
		tmpl string
		exp  string
		err  bool
	}{
		{
			name: "simple render",
			tmpl: `{{ .Key }}`,
			vars: map[string]interface{}{"Key": "Value"},
			exp:  "Value",
		},
		{
			name: "missing key",
			tmpl: `{{ .Key }}`,
			err:  true,
		},
		{
			name: "truncate",
			tmpl: `{{ "this is a line\n" | truncate }}`,
			exp:  "this is a line",
		},
		{
			name: "truncate multiple lines",
			tmpl: `{{ "this is a line\n\n\n" | truncate }}`,
			exp:  "this is a line",
		},
		{
			name: "truncate multiple lines",
			tmpl: `{{ "this is a line\n\n\n" | truncate }}`,
			exp:  "this is a line",
		},
		{
			name: "stripansi",
			tmpl: `{{ "\x1b[38;5;140mfoo\x1b[0m bar" | stripansi }}`,
			exp:  "foo bar",
		},
		{
			name: "exec & code",
			tmpl: `{{ exec "echo hallo" | truncate | code "bash" }}`,
			exp:  "```bash\n" + "hallo\n" + "```",
		},
		{
			name: "hook",
			tmpl: `{{ hook "echo hallo" }}`,
			exp:  "",
		},
		{
			name: "toc",
			opts: []RendererOptions{WithTemplateFile("testdata/toc.md")},
			tmpl: `# ToC
{{ toc }}
# 1. Heading
## 2. Heading
### 3. Heading
## 5. Heading`,
			exp: `# ToC
- [ToC](#toc)
- [1. Heading](#1.-heading)
  - [2. Heading](#2.-heading)
    - [3. Heading](#3.-heading)
  - [5. Heading](#5.-heading)

# 1. Heading
## 2. Heading
### 3. Heading
## 5. Heading`,
		},
		{
			name: "file",
			tmpl: `This is text
{{ file "testdata/include.md" }}`,
			exp: `This is text
include this text
`,
		},
		{
			name: "tmpl",
			tmpl: `This is text
{{ tmpl "testdata/tmpl.tmpl" }}`,
			exp: `This is text
HELLO!HELLO!HELLO!HELLO!HELLO!
`,
		},
		{
			name: "tmpl with vars",
			tmpl: `This is text
{{ tmplWithVars "./testdata/tmpl-vars.tmpl" "Var=Kubernetes" }}`,
			exp: `This is text
KUBERNETESKUBERNETESKUBERNETESKUBERNETESKUBERNETES
`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := Render([]byte(tc.tmpl), tc.vars, tc.opts...)

			if tc.err {
				require.Error(t, err, "expected an error but did not get one")

				return
			}

			require.NoError(t, err, "expected no error but got one")
			require.Equal(t, tc.exp, out.String())
		})
	}
}
