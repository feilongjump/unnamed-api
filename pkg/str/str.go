package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 转为复数
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular 转为单数
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake 转为 snake_case
func Snake(str string) string {
	return strcase.ToSnake(str)
}

// Camel 转为 CamelCase
func Camel(str string) string {
	return strcase.ToCamel(str)
}

// LowerCamel 转为 lowerCamelCase
func LowerCamel(str string) string {
	return strcase.ToLowerCamel(str)
}
