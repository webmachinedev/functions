package functions

import (
	"encoding/json"
	"net/http"
)

func FunctionIndex() (functions []string, err error) {
	r, err := http.Get("https://raw.githubusercontent.com/webmachinedev/functions/main/index.json")
	if err != nil {
		return nil, err
	}

	var functionids []string
	json.NewDecoder(r.Body).Decode(&functionids)

	return functionids, nil
}
