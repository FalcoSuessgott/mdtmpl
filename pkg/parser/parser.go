package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/FalcoSuessgott/mkdgo/pkg/template"
)

const commentRegex = `<!---\s*(.*?)\s*--->`

func Parse(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)

	re := regexp.MustCompile(commentRegex)

	var resultFile bytes.Buffer
	ln := 1
	for scanner.Scan() {

		if re.MatchString(scanner.Text()) {
			resultFile.Write(scanner.Bytes())
			resultFile.Write([]byte("\n"))

			b := getcommand(scanner.Text())
			result, err := template.Render([]byte(b))
			if err != nil {
				return nil, fmt.Errorf("cannot render template at line %d: %v", ln, err)
			}

			resultFile.Write(result.Bytes())
			resultFile.Write([]byte("\n"))
		} else {
			resultFile.Write(scanner.Bytes())
			resultFile.Write([]byte("\n"))
		}

		ln++
	}

	return resultFile.Bytes(), nil
}

func getcommand(s string) string {
	return regexp.MustCompile(commentRegex).FindStringSubmatch(s)[1]
}
