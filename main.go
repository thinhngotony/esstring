package main

import (
	"fmt"
	"log"
	"main/esstring"
)

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
		result := esstring.ToSnakeCase(test)
		fmt.Printf("Input: %v, Output: %s\n", test, result)
	}

	log.Println("======================================")

	snake := "example_snake_case_string"
	pascal := esstring.ToPascalCase(snake)
	log.Println(pascal) // Output: ExampleSnakeCaseString
}
