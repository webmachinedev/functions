package functions

import (
	"bytes"
	"text/template"
)

var vercelfunctiontmpl string = `package {{ .GoPackageName }}

import (
	"encoding/json"
	"net/http"

	"github.com/webmachinedev/functions"
	"github.com/webmachinedev/types"
)

type Request struct {
{{ range .Inputs }}
	{{ .Name }} {{ .Type }}
{{ end }}}

type Response struct {
{{ range .Outputs }}
	{{ .Name }} {{ .Type }}
{{ end }}}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var res Response

	{{ range $index, $output := .Outputs }}{{ if $index }},{{ end }}{{ $output.Name }}{{ end }} := functions.{{ .Name }}({{ range $index, $input := .Inputs }}{{ if $index }},{{ end }}req.{{ $input.Name }}{{ end }})

    json.NewEncoder(w).Encode(res)
}
`

func GenerateVercelAPI() (mutation map[string]string, err error) {
	filetmpl, err := template.New("vercelfunction").Parse(vercelfunctiontmpl)
	if err != nil {
		panic(err)
	}

	functionids, err := FunctionIndex()
	if err != nil {
		return nil, err
	}

	files := map[string]string{}

	for _, id := range functionids {
		function, err := GetFunction(id)
		if err != nil {
			return nil, err
		}
		
		contents := bytes.NewBuffer(nil)
		err = filetmpl.Execute(contents, function)
		if err != nil {
			return nil, err
		}
		
		path := "api/"+function.GoPackageName+"/index.go"
		files[path] = contents.String()
	}

	return files, nil
}
