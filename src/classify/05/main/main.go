package main

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] == ' ' {
			s = s[:i] + "%20" + s[i+1:]
			i += 3
			j += 2
		} else {
			i++
		}
	}
	return s
}
