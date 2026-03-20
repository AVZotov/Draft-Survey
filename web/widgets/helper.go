package widgets

import "strconv"

func formatFloat(f float64) string {
	if f == 0 {
		return ""
	}
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func formatInt(i int) string {
	if i == 0 {
		return "-"
	}
	return strconv.Itoa(i)
}
