package parser

import (
	"fmt"
	"strings"
)

func ParseHeader(header string) error {
	headerSlice := strings.Split(header, "\n")
	fmt.Printf("first value: %v\n", headerSlice)

	return nil
}
