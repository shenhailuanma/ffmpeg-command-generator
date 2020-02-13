package templates

import (
	"bytes"
	"strings"
	"text/template"
)

func GenerateCommand(name string, templateText string, request interface{}) (string, error) {

	tmpl := template.New(name).Funcs(template.FuncMap{})

	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, request)
	if err != nil {
		return "", err
	}

	var cmdString = buf.String()

	// delete '\n'
	cmdString = strings.Replace(cmdString, "\n", " ", -1)
	cmdString = strings.Replace(cmdString, "\t", " ", -1)
	cmdString = strings.TrimSpace(cmdString)

	return cmdString, nil
}
