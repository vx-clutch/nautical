package compiler_errors

import (
	"fmt"
	"os"
)

func errPrint(err string, correct string) {
	msg := fmt.Sprintf("error %v\n\tcorrect usage: %v", err, correct)
	fmt.Println(msg)
}

func Arguments() {
	errPrint("Not enough arguments", "knot <sub-command> <arg0> <arg...>")
	os.Exit(0)
}
