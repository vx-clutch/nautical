package compiler

import "fmt"

type token struct {
	kind  string
	value string
}

func Compiler(file []byte) (program []byte) {
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
		if current == 6 {
			break
		}
	}
	return tokens
}

func codeGeneration(tokens []token) []byte {
	program := "assembly fr fr"
	fmt.Println(tokens)
	return []byte(program)
}
