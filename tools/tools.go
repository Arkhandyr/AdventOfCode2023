package tools

func contains(slice []int, val int) bool {
	for _, e := range slice {
		if e == val {
			return true
		}
	}
	return false
}
