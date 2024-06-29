package knot

import (
	"flag"
	"fmt"
	"nautical/pkg/compiler"
	compiler_errors "nautical/pkg/errors"
	"os"
	"strings"
)

func help() {
	fmt.Println("help")
	os.Exit(0)
}

func build() {
	if has(3) {
		compiler_errors.Arguments()
		os.Exit(0)
	}
	input := os.Args[3]
	output := flag.String("o", strings.TrimSuffix(os.Args[3], ".nm"), "error: incorrect usage\n\tcorrect usage: knot build -o <output> <filename>.nm")
	flag.Parse()
	if strings.HasSuffix(input, ".nm") {
		fmt.Println()
		os.Exit(0)
	}
	inputContents, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	program := compiler.Compiler(inputContents)
	os.WriteFile(*output, program, 0666)
}

func has(num int) bool {
	if len(os.Args) > num {
		return false
	}
	return true
}

func is(num int) bool {
	if len(os.Args) == num {
		return true
	}
	return false
}

func Parse(pOption *string) {
	opt := *pOption
	switch opt {
	case "build":
		build()
	case "help":
		help()
	default:
		help()
	}
}
