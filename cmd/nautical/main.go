package main

import (
	"fmt"
	"nautical/pkg/compiler"
	"os"
	"strings"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("error: not enogh arguments\n\tcorrect usage: knot <filename>.nm")
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[1], ".hm") {
		fmt.Println("error: invalid file extention\n\tcorrect usage: <filename>.nm")
		os.Exit(1)
	}
}

func main() {
	filename := os.Args[1]
	path := strings.TrimSuffix(filename, ".nm")
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 2 {
		if os.Args[2] == "-o" {
			path = os.Args[3]
		}
	} else {
		fmt.Println("error: no argument after output flag\n\tcorrect usage: -o <output>")
		os.Exit(1)
	}
	program := compiler.Compiler(content)
	err = os.WriteFile(path, program, 0666)
	if err != nil {
		panic(err)
	}
}
