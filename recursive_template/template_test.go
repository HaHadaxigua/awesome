package recursive_template

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func TestRecursiveRender(t *testing.T) {
	funcMap := map[string]any{
		"z": func() any {
			return "{{ x }}"
		},
		"x": func() any {
			return "1"
		},
		"y": func() any {
			return `{{ if myeq "1" z  }} yyy {{ else }} ggg {{ end }}`
		},
		"myeq": func(a, b string) bool {
			return a == b
		},
	}

	content := "{{ y }}"
	cnt := 1000
	for IsTemplate(content) && cnt > 0 {
		tmpl := template.New("sss").Funcs(funcMap)
		tmpl, err := tmpl.Parse(content)
		if err != nil {
			t.Fatal(err)
		}

		var writer bytes.Buffer
		if err = tmpl.Execute(&writer, nil); err != nil {
			t.Fatal(err)
		}
		content = writer.String()
		cnt--
	}
	fmt.Println(content)
}
