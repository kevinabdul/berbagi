package stringarr

func Contains(list []string, target string) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}

	return false
}