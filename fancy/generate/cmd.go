package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("html")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	buf.WriteString("package fancy\n\n")
	buf.WriteString("// Autogenerated -- DO NOT EDIT BY HAND.\n\n")
	buf.WriteString("const (\n")

	for _, file := range files {
		fileName := file.Name()

		if strings.HasPrefix(fileName, ".") {
			continue
		}

		fileName = strings.Split(fileName, ".")[0]

		src, err := ioutil.ReadFile("html" + string(os.PathSeparator) + fileName + ".html")
		if err != nil {
			panic(err)
		}

		buf.WriteString("\t// " + fileName + " is a copy of html/" + fileName + ".html\n")
		buf.WriteString("\t" + fileName + " = `")
		buf.Write(src)
		buf.WriteString("`\n")
	}

	buf.WriteString(")\n")

	ioutil.WriteFile("autogenerated.go", buf.Bytes(), 0o755)
}
