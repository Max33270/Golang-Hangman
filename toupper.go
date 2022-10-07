package piscine

func ToUpper(s string) string {
	tableau := []rune(s)
	for i, v := range tableau {
		if (v >= 97 && v <= 122) || (v > 223 && v < 253) {
			tableau[i] = v - 32
		}
	}
	return (string(tableau))
}
