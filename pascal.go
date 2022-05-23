package strcase

func ToPascal(s string) string {
	return toCamelCaseOrPascalCase(s, true)
}
