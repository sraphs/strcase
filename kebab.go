package strcase

const kebab = '-'

// ToKebab converts a string to kebab-case
func ToKebab(s string) string {
	return ToDelimited(s, kebab)
}

func ToKebabWithIgnore(s string, ignore string) string {
	return toScreamingDelimited(s, kebab, ignore, false)
}
