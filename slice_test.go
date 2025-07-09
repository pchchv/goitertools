package goitertools

func stdRetain(s []int) []int {
	var j int
	for _, v := range s {
		if v == 1 {
			s[j] = v
			j++
		}
	}

	return s[:j]
}

func stdRetainFn(s []int, fn func(v int) bool) []int {
	var j int
	for _, v := range s {
		if fn(v) {
			s[j] = v
			j++
		}
	}

	return s[:j]
}
