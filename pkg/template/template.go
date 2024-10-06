package template

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var funcMap template.FuncMap = map[string]any{
	"file": func(file string) (string, error) {
		f, err := os.Open(file)
		if err != nil {
			return "", err
		}

		b, err := io.ReadAll(f)
		if err != nil {
			return "", err
		}

		return string(b), err
	},
	"exec": func(command string) (string, error) {
		cmd := exec.Command("sh", "-c", command)

		out, err := cmd.Output()
		if err != nil {
			return "", err
		}

		return string(out), nil
	},
	"code": func(language, content string) string {
		return fmt.Sprintf("```%s\n%s\n```", language, content)
	},
	"inlineCode": func(content string) string {
		return fmt.Sprintf("`%s`", content)
	},
}

// Render renders the given content using the sprig template functions.
func Render(content []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer

	tpl, err := template.New("template").
		Option("missingkey=error").
		Funcs(sprig.FuncMap()).Funcs(funcMap).
		Parse(string(content))
	if err != nil {
		return buf, err
	}

	if err := tpl.Execute(&buf, nil); err != nil {
		return buf, err
	}

	return buf, nil
}
