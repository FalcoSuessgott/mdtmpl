package template

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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
	"truncate": strings.TrimSpace,
}

// Render renders the given content using the sprig template functions.
// nolint: funlen, cyclop
func Render(content []byte, vars interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer

	tpl, err := template.New("template").
		Option("missingkey=error").
		Funcs(sprig.FuncMap()).Funcs(funcMap).Funcs(template.FuncMap{
		// we define tmpl here so we dont have a cyclic dependency
		"tmpl": func(file string) (string, error) {
			f, err := os.Open(file)
			if err != nil {
				return "", err
			}

			b, err := io.ReadAll(f)
			if err != nil {
				return "", err
			}

			res, err := Render(b, nil)
			if err != nil {
				return "", fmt.Errorf("failed to render template: %w", err)
			}

			return res.String(), nil
		},
		"tmplWithVars": func(file string, vars ...string) (string, error) {
			f, err := os.Open(file)
			if err != nil {
				return "", err
			}

			b, err := io.ReadAll(f)
			if err != nil {
				return "", err
			}

			m := map[string]interface{}{}

			for _, s := range vars {
				parts := strings.Split(s, "=")
				//nolint: mnd
				if len(parts) != 2 {
					return "", fmt.Errorf("invalid variable format: %s", s)
				}

				m[parts[0]] = parts[1]
			}

			res, err := Render(b, m)
			if err != nil {
				return "", fmt.Errorf("failed to render template: %w", err)
			}

			return res.String(), nil
		},
	}).
		Parse(string(content))
	if err != nil {
		return buf, err
	}

	if err := tpl.Execute(&buf, vars); err != nil {
		return buf, err
	}

	return buf, nil
}
