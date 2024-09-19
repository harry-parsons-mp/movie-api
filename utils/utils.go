package utils

import (
	"fmt"
	"strconv"
)

func Str_to_uint(input string) uint {
	converted, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Printf("problem converting %s into uint", input)
	}
	return uint(converted)
}
