package main

import (
	"flag"
	"fmt"
	"nautical/pkg/compiler"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[1], ".hm") {
		fmt.Println("error: invalid file extention\n\tcorrent usage: <filename>.nm")
		os.Exit(1)
	}
	outputFilePath := flag.String("o", "", "Specify the output file path")
	flag.Parse()
	if *outputFilePath != "" {
		path := outputFilePath
	} else {
		path := outputFilePath
	}
	filename := os.Args[1]
	defaultPath := strings.TrimSuffix(filename, ".nm")
	fmt.Println(defaultPath)
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	program := compiler.Compiler(content)
	err = os.WriteFile(path, program, 0666)
	if err != nil {
		panic(err)
	}
}
