package functions

import (
	"io"
	"text/template"

	"github.com/webmachinedev/types"
)

func WriteFunction(w io.Writer, f *types.Function) error {
	filetmpl, err := template.New("gofunction").Parse(gofunctiontempl)
	if err != nil {
		panic(err)
	}
	
	err = filetmpl.Execute(w, f)
	return err
}

var gofunctiontempl string = `func {{ .Name }}({{ range .Inputs }}{{ if $index }},{{ end }}{{ .Name }}{{ .Type }}{{ end }}) ({{ range .Outputs }}{{ if $index }},{{ end }}{{ .Name }}{{ .Type }}{{ end }}) {
	{{ .Body }}
}
`