package strcase

func ToCamel(s string) string {
	return toCamelCaseOrPascalCase(s, false)
}
