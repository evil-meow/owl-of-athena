package validation

import "unicode"

func IsRFC1123(s string) bool {
	if len(s) > 63 || !isLower(s) {
		return false
	}

	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
