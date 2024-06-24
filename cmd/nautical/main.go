package main

import (
	"fmt"
	"nautical/pkg/create"
	"os"
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
	fmt.Println("build")
}
func project() {
	create.Create(os.Args[2])
}
