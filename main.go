package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func main() {
	lines := []string{
		`eka`,
		`lolcat {{"toka" | upper | repeat 5 }}`,
	}

	for _, line := range lines {
		var tmpl *template.Template

		if t, err := template.New("jepjep").Funcs(sprig.FuncMap()).Parse(line); err != nil {
			panic(err)
		} else {
			tmpl = t
		}

		var result string
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, nil); err != nil {
			panic(err)
		}
		result = buf.String()
		fmt.Println(result)
	}
}
