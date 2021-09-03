package functions

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"strings"
)

func GetFunction(id string) (function struct{
	Name string
	GoPackageName string
	Inputs []struct{
		Name string
		Type string
	}
	Outputs []struct{
		Name string
		Type string
	}
}, err error) {
	filepath := id+".go"
	fileurl := "https://raw.githubusercontent.com/webmachinedev/functions/main/"+filepath
	r, err := http.Get(fileurl)
	if err != nil {
		return function, err
	}

	filebytes, err := io.ReadAll(r.Body)
	if err != nil {
		return function, err
	}

	file := string(filebytes)

	fset := token.NewFileSet()

	gofile, err := parser.ParseFile(fset, "", file, 0)
	if err != nil {
		return function, err
	}

	for _, obj := range gofile.Scope.Objects {
		if strings.ToLower(obj.Name) == id {
			function.Name = obj.Name
			function.GoPackageName = strings.ToLower(obj.Name)
	
			for _, input := range obj.Decl.(*ast.FuncDecl).Type.Params.List {
				t := file[input.Type.Pos()-1:input.Type.End()-1]
				var name string
				if len(input.Names) == 0 {
					name = GetNameFromType(t)
				} else {
					name = input.Names[0].Name
				}
				function.Inputs = append(function.Inputs, struct{
					Name string
					Type string
				}{
					Name: name,
					Type: t,
				})
			}

			for _, output := range obj.Decl.(*ast.FuncDecl).Type.Results.List {
				t := file[output.Type.Pos()-1:output.Type.End()-1]
				var name string
				if len(output.Names) == 0 {
					name = GetNameFromType(t)
				} else {
					name = output.Names[0].Name
				}
				function.Outputs = append(function.Outputs, struct{
					Name string
					Type string
				}{
					Name: name,
					Type: t,
				})
			}

			return function, nil
		}
	}

	return function, errors.New("function not found in "+ filepath)
}
