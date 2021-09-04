package functions

import "github.com/webmachinedev/types"

func AllFunctions() map[string]types.Function {
	pkg := GetPackage("github.com/webmachinedev/functions")
	return pkg.Functions
}
