package strcase_test

import (
	"fmt"

	"github.com/sraphs/strcase"
)

func Example() {
	fmt.Println(strcase.ToCamel("AnyKind.of_string"))
	fmt.Println(strcase.ToPascal("AnyKind.of_string"))
	fmt.Println(strcase.ToSnake("AnyKind.of_string"))
	fmt.Println(strcase.ToSnakeWithIgnore("AnyKind.of_string", "."))
	fmt.Println(strcase.ToDot("AnyKind.of_string"))
	fmt.Println(strcase.ToDotWithIgnore("AnyKind.of_string", "_"))
	fmt.Println(strcase.ToKebab("AnyKind.of_string"))
	fmt.Println(strcase.ToKebabWithIgnore("AnyKind.of_string", "."))
	// Output:
	// anyKindOfString
	// AnyKindOfString
	// any_kind_of_string
	// any_kind.of_string
	// any.kind.of.string
	// any.kind.of_string
	// any-kind-of-string
	// any-kind.of-string
}
