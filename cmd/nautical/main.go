package main

import (
	"flag"
	"fmt"
	"nautical/pkg/create"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		help()
	}
	if os.Args[1] == "build" {
		build()
	}
	if os.Args[1] == "new" {
		project()
	}
}

func help() {
	fmt.Println("help")
	os.Exit(0)
}
func build() {
	filename := os.Args[3]
	oFlag := flag.String("o", strings.TrimSuffix(filename, ".nm"), "Specify output location")
	fmt.Println(oFlag)
}
func project() {
	create.Create(os.Args[2])
}
