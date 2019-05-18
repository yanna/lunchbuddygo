package core

func isIntInSlice(val int, list []int) bool {
	for _, currVal := range list {
		if currVal == val {
			return true
		}
	}
	return false
}
