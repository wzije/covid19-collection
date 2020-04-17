package utils

import (
	"fmt"
	"os"
	"strconv"
)

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("parsed : %x", result)
	return result
}
