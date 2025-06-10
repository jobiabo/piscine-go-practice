package piscine

func RetainFirstHalf(s string) string {
	length := len(s)

	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	midpoint := length / 2
	return s[:midpint]
}
