package compiler

import "fmt"

type token struct {
	kind  string
	value string
}

func Compile(file []byte) (program []byte) {
	tokens := tokenizer(file)
	program = codeGeneration(tokens)
	return
}

func tokenizer(src []byte) []token {
	source := string(src[:])
	source += "\n"
	current := 0
	tokens := []token{}

	for current < len([]rune(source)) {
		char := string([]rune(source)[current])
		if char == "(" {
			tokens = append(tokens, token{kind: "lparen", value: "("})
			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{kind: "rparen", value: ")"})
			current++
			continue
		}
		if char == "[" {
			tokens = append(tokens, token{kind: "lbracket", value: "["})
			current++
			continue
		}
		if char == "]" {
			tokens = append(tokens, token{kind: "rbracket", value: "]"})
			current++
			continue
		}
		if char == "{" {
			tokens = append(tokens, token{kind: "lbrace", value: "{"})
			current++
			continue
		}
		if char == "}" {
			tokens = append(tokens, token{kind: "rbrace", value: "}"})
			current++
			continue
		}
		if char == ";" {
			tokens = append(tokens, token{kind: "semicolon", value: ";"})
			current++
			continue
		}
		// FIX: don't ask, i do not know
		if current == 6 {
			break
		}
	}
	return tokens
}

// TODO:
func isAlpha()  {}
func isNumber() {}

func codeGeneration(ir []token) []byte {
	base := "global _start\n_start:\n\t"
	program := "assembly fr fr"
	fmt.Println(ir)
	program = base + program
	return []byte(program)
}
