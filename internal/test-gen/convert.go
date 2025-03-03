package test_gen

import (
	"strconv"
)

func strToInt(s string) (i int64) {
	if val, exists := vars[s]; exists {
		i, _ = strconv.ParseInt(val, 10, 64)
	} else {
		i, _ = strconv.ParseInt(s, 10, 64)
	}
	return
}

func strToFLoat(s string) (f float64) {
	if val, exists := vars[s]; exists {
		f, _ = strconv.ParseFloat(val, 64)
	} else {
		f, _ = strconv.ParseFloat(s, 64)
	}
	return
}

func intToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func end(s string) string {
	quoted := `"` + s + `"`
	s, _ = strconv.Unquote(quoted)
	return s
}
