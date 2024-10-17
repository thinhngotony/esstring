package main

import (
	"fmt"
	"log"

	"github.com/gookit/goutil"
	"github.com/iancoleman/strcase"
)

// ToSnakeCase converts an interface{} to a snake_case string.
// If the output is empty, it logs a warning and returns the original input.
func ToSnakeCase(input interface{}) (resp string) {
	defer func() {
		if len(resp) == 0 {
			log.Printf("Warning: cannot convert to snake case, returning original input [%v]", input)
			resp = goutil.String(input)
		}
	}()

	return strcase.ToSnake(goutil.String(input))
}

func main() {
	testCases := []interface{}{
		"HelloWorld",
		"someVariableName",
		123,
		456.789,
		true,
		"",
		nil,
		[]string{"some", "list"},
		map[string]string{"key": "value"},
	}

	for _, test := range testCases {
		result := ToSnakeCase(test)
		fmt.Printf("Input: %v, Output: %s\n", test, result)
	}
}
