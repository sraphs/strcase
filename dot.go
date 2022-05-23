package strcase

const dot = '.'

// converts a string to dot.case
func ToDot(s string) string {
	return ToDelimited(s, dot)
}

func ToDotWithIgnore(s string, ignore string) string {
	return toScreamingDelimited(s, dot, ignore, false)
}
