package piscine

func IsAlpha(s string) bool {
	for _, ch := range s {
		if ch < 'A' || ch > 'Z' {
			return false
		}
	}
	return true
}
