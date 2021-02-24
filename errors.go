package main

import (
	"bytes"
	"text/template"

	_ "embed"
)

var (
	//go:embed templates/acme404Message.html
	acme404Message string

	//go:embed templates/acme404TI.html
	acme404TI string

	//go:embed templates/emptyHostHeader.html
	emptyHostHeader string

	//go:embed templates/genericMessage.html
	genericMessage string

	//go:embed templates/hostHeaderIsIPTechInfo.html
	hostHeaderIsIPTechInfo string

	//go:embed templates/pageTemplate.html
	pageTemplate string
)

var pageT *template.Template

func init() { pageT = template.Must(template.New("").Parse(pageTemplate)) }

// errorPage produces a nice looking HTML error page.
func errorPage(title, message, ti, ver string) []byte {
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
