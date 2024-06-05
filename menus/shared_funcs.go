package menus

import (
	"errors"
	"strings"
)

func spaceCalculator(budget int, previosWrods ...string) (string, error) {
	s := ""

	for _, word := range previosWrods {
		if len(word) > budget {
			return "", errors.New("previosWrods longer than budget")
		}

		budget -= len(word)
	}

	s += strings.Repeat(" ", budget)
	return s, nil

}
