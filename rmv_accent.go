package piscine

func Rmv_accent(s string) string {
	e := []rune(s)
	if rune(e[0]) > 231 && rune(e[0]) < 236 && len(s) == 2 {
		s = "e"
	} else if rune(e[0]) > 223 && rune(e[0]) < 231 && len(s) == 2 {
		s = "a"
	} else if rune(e[0]) > 235 && rune(e[0]) < 240 && len(s) == 2 {
		s = "i"
	} else if rune(e[0]) == 231 && len(s) == 2 {
		s = "c"
	} else if rune(e[0]) > 241 && rune(e[0]) < 247 && len(s) == 2 {
		s = "o"
	} else if rune(e[0]) > 248 && rune(e[0]) < 253 && len(s) == 2 {
		s = "u"
	}
	return s
}
