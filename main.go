package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/pkg/file"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	// fmt.Println([]byte(input))
	input = strings.ReplaceAll(input, "\n", "\\n")
	for _, w := range input {
		if w < ' ' || w > '~' {
			log.Fatalf("ERROR: Non-ASCII character found: %q", w)
		}
	}

	lines, err := file.ReadLine("standard.txt")
	if err != nil {
		fmt.Printf("Error: \"%s\"\n", err.Error())
		return
	}

	maps, err := ascii.Create(lines)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	// fmt.Println(len(maps['a']), maps['a'][8] == "")
	// fmt.Println(string(maps['a'][7][len(maps['a'][7])-1]) == " ")
	// _, ok := maps['~']
	// fmt.Println(ok)
	words := strings.Split(input, "\\n")
	// for i := ' '; i <= '~'; i++ {
	// 	_, ok := maps[i]
	// 	fmt.Println(string(i), ok)
	// }
	err = ascii.Print(maps, words)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
