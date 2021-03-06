package tdd_golang

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	s := splitString(input)
	n := 0
	temp := 0

	for _, v := range s {
		temp, _ = strconv.Atoi(v)
		if temp < 0 {
			return 0, errors.New(fmt.Sprintf("negatives not allowed: %v", v))
		}
		if temp > 1000 {
			temp = 0
		}
		n += temp
	}
	return n, nil
}

func splitString(input string) []string {
	var s []string
	if len(input) > 2 && input[0:2] == "//" {
		s = strings.FieldsFunc(input[2:len(input)], Split)
	} else {
		s = strings.FieldsFunc(input, Split)
	}
	return s
}

func Split(r rune) bool {
	return r == ',' || r == '\n' || r == ';'
}
