package main

import (
	"fmt"

	"github.com/webmachinedev/functions"
)

func main() {
	// f, _ := functions.GetFunction("getfunction")

	// switch d := f.Decl.(type) {
	// case *ast.FuncDecl:
	// 	fmt.Println(d.Type.Params)
	// default:
	// 	panic("not func decl")
	// }
	// for _, param := range f.Decl.(*ast.FuncDecl).Type.Params.List {
	// 	fmt.Println(param.Names[0].Name)
	// 	fmt.Println(param.Type)
	// }


	files, err := functions.GenerateVercelAPI()
	fmt.Println(err)
	for name, contents := range files {
		fmt.Println(name)
		fmt.Println(contents)
	}
}
