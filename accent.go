package piscine

func Accent(a []rune) []rune{
	w := []rune{}
	for v := range a {
		if a[v] > 231 && a[v] < 236 {
			w = append(w, 'e')
		} else if a[v] > 223 && a[v] < 231 {
			w = append(w, 'a')
		} else if a[v] > 235 && a[v] < 240 {
			w = append(w, 'i')
		} else if a[v] == 231 {
			w = append(w, 'c')
		} else if a[v] > 241 && a[v] < 247 {
			w = append(w, 'o')
		} else if a[v] > 248 && a[v] < 253 {
			w = append(w, 'u')
		} else {
			w = append(w, a[v])
		}
	}
	return w
}