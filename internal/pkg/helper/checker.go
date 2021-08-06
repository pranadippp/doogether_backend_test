package helper

import "strings"

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsAlphabet(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') &&
			(charVariable < 'A' || charVariable > 'Z') &&
			(charVariable > ' ') {
			return false
		}
	}
	return true
}
