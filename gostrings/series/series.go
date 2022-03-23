package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	series := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		series = append(series, s[i:i+n])
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	res, ok := First(n, s)
	if ok {
		return res
	}
	panic("UnsafeFirst: no possible series")
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return s, false
	}
	return s[:n], true
}
