package functions

import "github.com/webmachinedev/types"

func AllPackages() map[string]*types.Package {
	pkgids := []string{
		"github.com/webmachinedev/functions",
		"github.com/webmachinedev/types",
	}

	for _, pkgid := range pkgids {
		pkg := GetPackage(pkgid)
		
	}
}
