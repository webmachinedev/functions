package functions

import (
	"errors"

	"github.com/webmachinedev/types"
)

func GetFunction(id string) (types.Function, error) {
	functions := map[string]types.Function{}
	if function, ok := functions[id]; ok {
		return function, nil
	} else {
		return types.Function{}, errors.New("not found")
	}
}
