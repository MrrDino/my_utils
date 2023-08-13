package myutils

func Contains(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}
	}
	return false
}
