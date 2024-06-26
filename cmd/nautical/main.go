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
