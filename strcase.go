// package strcase is a go package for converting string case to various cases (e.g. snake case or camel case).
package strcase

import (
	"strings"
)

// ToDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
func ToDelimited(s string, delimiter rune) string {
	return toScreamingDelimited(s, delimiter, "", false)
}

// toScreamingDelimited converts a string to SCREAMING.DELIMITED.SNAKE.CASE
// (in this case `delimiter = '.'; screaming = true`)
// or delimited.snake.case
// (in this case `delimiter = '.'; screaming = false`)
func toScreamingDelimited(s string, delimiter rune, ignore string, screaming bool) string {
	s = strings.TrimSpace(s)
	var b strings.Builder
	b.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters

	for i, v := range s {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(s) {
			next := s[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			// add underscore if next letter case type is changed
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
							b.WriteRune(delimiter)
						}
					}
					b.WriteRune(v)
					if vIsLow || vIsNum || nextIsNum {
						b.WriteRune(delimiter)
					}
					continue
				}
			}
		}

		if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
			// replace space/underscore/hyphen/dot with delimiter
			b.WriteRune(delimiter)
		} else {
			b.WriteRune(v)
		}
	}

	return b.String()
}

// Converts a string to CamelCase or PascalCase
func toCamelCaseOrPascalCase(s string, pascalCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	var b strings.Builder

	capNext := pascalCase
	for i, v := range s {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			b.WriteRune(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			b.WriteRune(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return b.String()
}
