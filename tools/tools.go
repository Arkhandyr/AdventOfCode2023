package tools

func contains(slice []any, val any) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}
