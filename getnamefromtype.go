package functions

func GetNameFromType(t string) (name string) {
	switch t {
	case "error":
		return "err"
	case "string":
		return "str"
	default:
		panic("no name for type " + t)
	}
}
