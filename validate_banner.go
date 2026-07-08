package main

import (
	"errors"
	"fmt"
)

func validate(str map[rune][]string) error {

	if str == nil {
		return errors.New("invalid")
	}
	if len(str) != 95 {
		return errors.New("invalid")
	}

	for i, j := range str {
		if len(j) < 32 || len(j) > 126 {
			return fmt.Errorf("character %c has %d", i, j)
		}
		if rune(i) != 8 {
			return fmt.Errorf("invalid %c", j)
		}
	}
	return nil
}
