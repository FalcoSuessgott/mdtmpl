package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/FalcoSuessgott/mdtmpl/pkg/template"
)

const commentRegex = `<!---\s*(.*?)\s*--->`

func Parse(r io.Reader) ([]byte, error) {
	var resultFile bytes.Buffer

	re := regexp.MustCompile(commentRegex)

	ln := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			resultFile.Write(scanner.Bytes())
			resultFile.WriteString("\n")

			b := parseInstruction(scanner.Text())

			result, err := template.Render([]byte(b), nil)
			if err != nil {
				return nil, fmt.Errorf("cannot render template at line %d: %w", ln, err)
			}

			resultFile.Write(result.Bytes())
			resultFile.WriteString("\n")
		} else {
			resultFile.Write(scanner.Bytes())
			resultFile.WriteString("\n")
		}

		ln++
	}

	return resultFile.Bytes(), nil
}

func parseInstruction(s string) string {
	return regexp.MustCompile(commentRegex).FindStringSubmatch(s)[1]
}
