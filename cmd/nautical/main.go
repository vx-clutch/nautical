package main

import (
	"flag"
	"fmt"
	"nautical/pkg/compiler"
	"os"
	"strings"
)

func init() {

}

func main() {
	commandChecks()
	filename := os.Args[1]
	defaultPath := strings.TrimSuffix(filename, ".nm")
	fmt.Println(defaultPath)
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	program := compiler.Compiler(content)
	err = os.WriteFile(defaultPath, program, 0666)
	if err != nil {
		panic(err)
	}
}

func commandChecks() {
	if len(os.Args) < 2 {
		os.Exit(1)
	} // Checks for arguments
	if strings.HasSuffix(os.Args[1], ".hm") {
		fmt.Println("error: invalid file extention\n\tcorrent usage: <filename>.nm")
		os.Exit(1)
	} // Checks for proper file extentions
}
