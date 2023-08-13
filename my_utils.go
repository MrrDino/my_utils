package myutils

func InSlice(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}
