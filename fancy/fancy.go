//go:generate go run generate/cmd.go

package fancy

import (
	"bytes"
	"text/template"
)

var pageT *template.Template

func init() { pageT = template.Must(template.New("").Parse(pageTemplate)) }

// ErrorPage produces a nice looking HTML error page.
func ErrorPage(title, message, ti, ver string) []byte {
	footer := "<footer>Powered by tls-redirector/" + ver + "</footer>"

	dat := struct {
		Title   string
		Message string
		TI      string
		Footer  string
	}{title, message, ti, footer}

	var buf bytes.Buffer
	err := pageT.Execute(&buf, dat)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
