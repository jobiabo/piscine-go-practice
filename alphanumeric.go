package piscine

func AlphaNumeric(s string) bool {

	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true

}
