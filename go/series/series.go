package slice

func All(n int, s string) (r []string) {
	for i := 0; i <= len(s)-n; i++ {
		r = append(r, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
