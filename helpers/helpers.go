package helpers

import (
	"encoding/json"
	"fmt"
)

func PrintJson(s any) {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal:", err)
	} else {
		fmt.Println(string(pretty))
	}
}
