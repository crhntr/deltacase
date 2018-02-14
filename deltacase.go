package deltacase

import "strings"

func Camel(str string) string {
	parts, out := slice(str), ""
	if len(parts) > 0 {
		out = parts[0]
		if len(parts) > 1 {
			for _, part := range parts[1:] {
				out += strings.Title(part)
			}
		}
	}
	return out
}

func Pascal(str string) string {
	out := ""
	for _, part := range slice(str) {
		out += strings.Title(part)
	}
	return out
}

func Kebab(str string) string { return strings.Join(slice(str), "-") }

func Snake(str string) string { return strings.Join(slice(str), "_") }

func Spaced(str string) string { return strings.Join(slice(str), " ") }

func UpperSnake(str string) string { return strings.ToUpper(Snake(str)) }

func slice(str string) []string {
	str = strings.TrimSpace(str)
	parts := []string{}
	if strings.Contains(str, " ") {
		parts = strings.Split(str, " ")
	} else if strings.Contains(str, "-") {
		parts = strings.Split(str, "-")
	} else if strings.Contains(str, "_") {
		parts = strings.Split(str, "_")
	} else {
		if allUpper(str) {
			parts = []string{str}
		} else {
			previous, i, char := 0, 0, rune(0)
			for i, char = range str {
				if char >= 'A' && char <= 'Z' && i > 0 {
					parts = append(parts, str[previous:i])
					previous = i
				}
			}
			parts = append(parts, str[previous:len(str)])
		}
	}
	for i := range parts {
		parts[i] = strings.ToLower(strings.TrimSpace(parts[i]))
	}
	return parts
}

func allUpper(str string) bool {
	for _, char := range str {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}
