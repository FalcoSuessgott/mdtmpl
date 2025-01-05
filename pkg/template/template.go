package template

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/FalcoSuessgott/mdtmpl/pkg/commit"
	"github.com/Masterminds/semver/v3"
	"github.com/Masterminds/sprig/v3"
	"github.com/acarl005/stripansi"
)

const (
	gitCommitMsgFile    = ".git/COMMIT_EDITMSG"
	gitLatestTagCommand = "git describe --tags --abbrev=0"
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
		cmd.Env = os.Environ()
		cmd.Dir = os.Getenv("PWD")

		out, err := cmd.Output()
		if err != nil {
			return "", err
		}

		return string(out), nil
	},
	"hook": func(command string) (string, error) {
		cmd := exec.Command("sh", "-c", command)
		cmd.Env = os.Environ()
		cmd.Dir = os.Getenv("PWD")

		_, err := cmd.Output()
		if err != nil {
			return "", err
		}

		return "", nil
	},
	"code": func(language, content string) string {
		return fmt.Sprintf("```%s\n%s\n```", language, content)
	},
	"conventionalCommitBump": func() (string, error) {
		f, err := os.Open(gitCommitMsgFile)
		if err != nil {
			return "", err
		}

		b, err := io.ReadAll(f)
		if err != nil {
			return "", err
		}

		cmd := strings.Split(gitLatestTagCommand, " ")
		//nolint: gosec
		version, err := exec.Command(cmd[0], cmd[1:]...).Output()
		if err != nil {
			return "", fmt.Errorf("failed to get latest tag: %w", err)
		}

		semverF, err := commit.ParseConventionalCommit(bytes.TrimSpace(b))
		if err != nil {
			return "", fmt.Errorf("failed to parse commit as conventional: %w", err)
		}

		sv, err := semver.NewVersion(string(bytes.TrimSpace(version)))
		if err != nil {
			return "", fmt.Errorf("failed to parse version as semantic version: %w", err)
		}

		v := semverF(sv)
		if bytes.HasPrefix(version, []byte("v")) {
			v = "v" + v
		}

		return v, nil
	},
	"truncate":  strings.TrimSpace,
	"stripansi": stripansi.Strip,
}

type RendererOptions func(*Renderer)

type Renderer struct {
	tmplFile string
}

func WithTemplateFile(f string) RendererOptions {
	return func(p *Renderer) {
		p.tmplFile = f
	}
}

// Render renders the given content using the sprig template functions.
// nolint: funlen, cyclop
func Render(content []byte, vars interface{}, opts ...RendererOptions) (bytes.Buffer, error) {
	var r Renderer

	for _, opt := range opts {
		opt(&r)
	}

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

			res, err := Render(b, nil, opts...)
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

			res, err := Render(b, m, opts...)
			if err != nil {
				return "", fmt.Errorf("failed to render template: %w", err)
			}

			return res.String(), nil
		},
		"toc": func() (string, error) {
			// Read the markdown file
			out, err := os.ReadFile(r.tmplFile)
			if err != nil {
				return "", fmt.Errorf("failed to read file %s: %w", r.tmplFile, err)
			}

			// Regular expression to match markdown headings
			re := regexp.MustCompile(`(?m)^(#{1,6})\s+(.*)`)

			// Find all headings
			matches := re.FindAllStringSubmatch(string(out), -1)

			// Generate the table of contents
			var toc strings.Builder

			for _, match := range matches {
				level := len(match[1])
				heading := match[2]
				anchor := strings.ToLower(strings.ReplaceAll(heading, " ", "-"))
				toc.WriteString(fmt.Sprintf("%s- [%s](#%s)\n", strings.Repeat("  ", level-1), heading, anchor))
			}

			return toc.String(), nil
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
