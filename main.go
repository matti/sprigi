package main

import (
	"bytes"
	"fmt"

	// To generate HTML output, see html/template, which has the same interface as this package but automatically secures HTML output against certain attacks.

	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func main() {
	things := []string{
		`eka`,

		`lolcat {{"toka" | upper | repeat 5 }}`,

		`first
{{ if true -}}
second in if - and actually this is not sprig, but standard go template
{{ else -}}
never here, duh
{{ end -}}
lastly`,

		`{{printf "%q" (print "lol" "cat" "dog" )}}`,

		`{{- env "HOME" }}`,

		`{{- printf "%q" (env "HOME") }}`,

		// "WARNING: Some notable implementations of Sprig (such as Kubernetes Helm) do not provide these functions for security reasons."
		`{{- expandenv "Your USER is set to $USER" }}`,

		`Your USER is set to {{ expandenv "$USER" }}`,

		`{{ betterAdd 1 2 }}`,

		`{{ longcat | upper}}`,

		// when map
		// `{{ .lol | upper }}`,

		// when struct
		// `{{ .Lol | upper }}`,
	}

	data := map[string]string{
		"yes": "box",
		"lol": "cat",
	}

	// type Data struct {
	// 	Yes string
	// 	Lol string
	// }

	// data := Data{
	// 	Yes: "box",
	// 	Lol: "cat",
	// }

	funcMap := sprig.FuncMap()
	funcMap["betterAdd"] = func(a, b int) int {
		return a + b
	}
	funcMap["longcat"] = func() string {
		return "longcat is looooooooooooooooong"
	}

	for _, thing := range things {
		fmt.Println("thing: ", thing)

		var tmpl *template.Template

		if t, err := template.New("jepjep").Funcs(funcMap).Parse(thing); err != nil {
			panic(err)
		} else {
			tmpl = t
		}

		var result string
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			panic(err)
		}
		result = buf.String()
		fmt.Println("sprigified:", result)
		fmt.Println()
	}
}
