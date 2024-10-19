package esstring

import (
	"log"
	"strings"

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

func ToPascalCase(input interface{}) (resp string) {
	defer func() {
		if len(resp) == 0 {
			log.Printf("Warning: cannot convert to Pascal case, returning original input [%v]", input)
			resp = goutil.String(input)
		}
	}()

	str := goutil.String(input)
	// Split the string by underscores
	parts := strings.Split(str, "_")

	for i, part := range parts {
		if len(part) > 0 {
			// Capitalize the first letter of each part
			parts[i] = strings.Title(part)
		}
	}

	// Join the parts back together
	resp = strings.Join(parts, "")
	return resp
}
