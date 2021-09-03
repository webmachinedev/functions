package functions

import (
	"bytes"

	"github.com/webmachinedev/types"
)

func UpdateFunction(id string, f *types.Function, githubkey string) error {
	bytes := bytes.NewBuffer(nil)
	err := WriteFunction(bytes, f)
	if err != nil {
		return err
	}
	changes := map[string]string{
		id+".go": bytes.String(),
	}
	return WriteToGithub("webmachinedev", "functions", "main", changes, "Update "+id, githubkey)
}
