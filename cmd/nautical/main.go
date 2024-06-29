package main

import (
	"nautical/pkg/knot"
	"os"
)

func main() {
	var arg0 string
	if len(os.Args) > 1 {
		arg0 = os.Args[1]
	}
	knot.Parse(&arg0)
}
