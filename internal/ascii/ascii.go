package ascii

import (
	"fmt"
	"os"
)

func Create(content []string) (map[rune][]string, error) {
	if len(content) != 855 {
		return nil, fmt.Errorf("non-valid number of lines")
	}
	maps := make(map[rune][]string)
	element := rune(32)
	previous, current := 1, 9
	for current <= len(content) {
		maps[element] = content[previous:current]
		element++
		previous = current + 1
		current += 9
	}

	return maps, nil
}

func Print(maps map[rune][]string, words []string) error {
	if len(os.Args[1:]) != 1 {
		return fmt.Errorf("not enough arguments")
	}
	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, key := range word {
				fmt.Print(maps[key][i])
			}
			if i != 8 {
				fmt.Println()
			}
		}
	}
	return nil
}
