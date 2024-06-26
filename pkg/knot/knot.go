package knot

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println("help")
	os.Exit(0)
}

func build() {
	filename := os.Args[len(os.Args)-1]
	oFlag := flag.String("o", strings.TrimSuffix(filename, ".nm"), fmt.Sprintf("Specify output location\n\tcorrect usage: build -o %v <output>", filename))
	flag.Parse()
	fmt.Println(*oFlag)
	fmt.Println(filename)
}
