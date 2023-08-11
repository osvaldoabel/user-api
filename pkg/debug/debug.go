package debug

import (
	"encoding/json"
	"fmt"
	"os"
)

func DUMP(variable interface{}) {
	res, _ := json.MarshalIndent(variable, "", "  ")

	fmt.Println("= = = = = = = = = = = =")
	fmt.Println(string(res))
	fmt.Println("= = = = = = = = = = = =")
}

func DD(variable interface{}, die bool) {
	DUMP(variable)
	if die {
		os.Exit(1)
	}
}
