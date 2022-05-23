# strcase

[![CI](https://github.com/sraphs/strcase/actions/workflows/ci.yml/badge.svg)](https://github.com/sraphs/strcase/actions/workflows/ci.yml)

strcase is a go package for converting string case to various cases (e.g. [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase)) to see the full conversion table below.

## Example

```go
s := "AnyKind.of_string"
```

| Function                    | Result               |
| --------------------------- | -------------------- |
| `ToCamel(s)`                | `anyKindOfString`    |
| `ToPascal(s)`               | `AnyKindOfString`    |
| `ToSnake(s)`                | `any_kind_of_string` |
| `ToSnakeWithIgnore(s, '.')` | `any_kind.of_string` |
| `ToDot(s)`                  | `any.kind.of.string` |
| `ToDotWithIgnore(s, '_')`   | `any.kind.of_string` |
| `ToKebab(s)`                | `any-kind-of-string` |
| `ToKebabWithIgnore(s, '.')` | `any-kind.of-string` |

## Usage

```go
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

```

## Custom Acronyms for ToCamel && ToLowerCamel

Often times text can contain specific acronyms which you need to be handled a certain way.
Out of the box `strcase` treats the string "ID" as "Id" or "id" but there is no way to cater
for every case in the wild.

To configure your custom acronym globally you can use the following before running any conversion

```go
import (
    "github.com/sraphs/go/x/strcase"
)

func init() {
    // results in "Api" using ToCamel("API")
    // results in "api" using ToLowerCamel("API")
    strcase.ConfigureAcronym("API", "api")
    
    // results in "PostgreSQL" using ToCamel("PostgreSQL")
    // results in "postgreSQL" using ToLowerCamel("PostgreSQL")
    strcase.ConfigureAcronym("PostgreSQL", "PostgreSQL")

}

```
