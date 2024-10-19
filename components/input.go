package components

import (
	"fmt"
)

func Input(params ...map[string]interface{}) (interface{}, error) {
	if len(params) > 0 {
		if value, ok := params[0]["label"]; ok {
			fmt.Printf("%s ", value)
		}

		if value, ok := params[0]["type"]; ok {
			if value == "number" {
				var inputAngka int
				_, err := fmt.Scanln(&inputAngka)
				return inputAngka, err
			}
		}
	}

	var inputTeks string
	_, err := fmt.Scanln(&inputTeks)
	return inputTeks, err
}
