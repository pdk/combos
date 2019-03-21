package comp

// Combos computes all the combinations of n items from the list of things.
func Combos(n int, things []string) [][]string {

	var x [][]string

	if n <= 0 {
		return x
	}

	if n == 1 {
		for _, item := range things {
			x = append(x, []string{item})
		}
		return x
	}

	for i, item := range things {

		var sub []string
		sub = append(sub, things[:i]...)
		sub = append(sub, things[i+1:]...)

		for _, c := range Combos(n-1, sub) {

			var next []string
			next = append(next, item)
			next = append(next, c...)

			x = append(x, next)
		}
	}

	return x
}
