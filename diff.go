package main

// diff finds strings in b but not in a
func diff(a, b []string) []string {
	res := make([]string, 0, len(b))
	for _, p := range b {
		var found bool
		for _, q := range a {
			if p == q {
				found = true
			}
		}
		if !found {
			res = append(res, p)
		}
	}
	return res
}
