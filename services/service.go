package service

import (
	"errors"
	"fmt"
	"strings"
)

type StringService interface {
	MakeUpperCase(string) (string, error)
	MakeCount(string) int
}

type StringStruct struct {
}

func (StringStruct) MakeUpperCase(s string) (string, error) {
	fmt.Println("request with s:", s)
	if len(s) == 0 {
		return "", errors.New("empty string")
	}
	output := strings.ToUpper(s)
	fmt.Println("output with upper is:", output)
	return output, nil
}

func (StringStruct) MakeCount(s string) int {
	return len(s)
}
