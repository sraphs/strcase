package strcase

const snake = '_'

// ToSnake converts a string to snake_case
func ToSnake(s string) string {
	return ToDelimited(s, snake)
}

func ToSnakeWithIgnore(s string, ignore string) string {
	return toScreamingDelimited(s, snake, ignore, false)
}
