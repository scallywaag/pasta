package helpers

import (
	"encoding/json"
	"fmt"
)

func Prettify(s any) string {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal:", err)
	}

	return string(pretty)
}
