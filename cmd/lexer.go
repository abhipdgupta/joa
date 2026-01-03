package joa

import (
	"fmt"
	"iter"
	"strconv"
)

type Lexer struct {
	input  string
	tokens []Token
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		input:  source,
		tokens: make([]Token, 0),
	}

}

func (l *Lexer) Tokens() iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for _, t := range l.tokens {
			if !yield(t) {
				return
			}
		}
	}
}
func (l *Lexer) Lex() {
	input := l.input

	var i int = 0
	var lineNum int = 1
	for i < len(input) {
		char := input[i]

		switch {
		case char == ' ' || char == '\t':
			i = i + 1
		case char == '\n' || char == '\r':
			if char == '\r' && i+1 < len(input) && input[i+1] == '\n' {
				i = i + 1
			}
			i = i + 1
			lineNum = lineNum + 1
		case char == '{':
			l.tokens = append(l.tokens, Token{
				kind:  TokenOpenBrace,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == '}':
			l.tokens = append(l.tokens, Token{
				kind:  TokenCloseBrace,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == ':':
			l.tokens = append(l.tokens, Token{
				kind:  TokenColon,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == ',':
			l.tokens = append(l.tokens, Token{
				kind:  TokenComma,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == '[':
			l.tokens = append(l.tokens, Token{
				kind:  TokenOpenBracket,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == ']':
			l.tokens = append(l.tokens, Token{
				kind:  TokenCloseBracket,
				start: i,
				end:   i + 1,
			})
			i = i + 1
		case char == '"':
			var requirdIndex int = -1
			var j int = i + 1
			for j < len(input) {
				if input[j] == '\r' && j+1 < len(input) && input[j+1] == '\n' {
					lineNum = lineNum + 1
					j = j + 2
					continue
				}
				if input[j] == '\n' {
					lineNum = lineNum + 1
					j = j + 1
					continue
				}
				if input[j] == '\\' {
					j = j + 2
					continue
				}
				if input[j] == '"' {
					requirdIndex = j - (i + 1)
					break
				}
				j = j + 1
			}
			if requirdIndex == -1 {
				panic("Invalid character at position")
			}
			var str = input[i+1 : i+requirdIndex+1]
			l.tokens = append(l.tokens, Token{
				kind:  TokenString,
				data:  str,
				start: i + 1,
				end:   i + requirdIndex,
			})
			i = i + requirdIndex + 2
		case char == 'n':
			var strToCheck = input[i : i+4]

			if strToCheck != "null" {
				panic("Invalid character expected null")
			}
			l.tokens = append(l.tokens, Token{
				kind:  TokenNull,
				data:  strToCheck,
				start: i,
				end:   i + 4,
			})
			i = i + 4
		case char == 't':
			var strToCheck = input[i : i+4]
			if strToCheck != "true" {
				panic("Invalid character expected true")
			}
			l.tokens = append(l.tokens, Token{
				kind:  TokenTrue,
				data:  strToCheck,
				start: i,
				end:   i + 4,
			})
			i = i + 4
		case char == 'f':
			var strToCheck = input[i : i+5]
			if strToCheck != "false" {
				panic("Invalid character expected false")
			}
			l.tokens = append(l.tokens, Token{
				kind:  TokenFalse,
				data:  strToCheck,
				start: i,
				end:   i + 5,
			})
			i = i + 5
		case char == '-' || l.isDigit(char):
			var isFloat bool = false
			var start = i
			if char == '-' {
				i = i + 1
			}
			for i < len(input) && l.isDigit(input[i]) {
				i = i + 1
			}
			if input[i] == '.' {
				isFloat = true
				i = i + 1
				for i < len(input) && l.isDigit(input[i]) {
					i = i + 1
				}
			}
			if input[i] == 'e' || input[i] == 'E' {
				isFloat = true
				i = i + 1
				if input[i] == '+' || input[i] == '-' {
					i = i + 1
				}
				for i < len(input) && l.isDigit(input[i]) {
					i = i + 1
				}
			}
			str := input[start:i]
			l.tokens = append(l.tokens, Token{
				kind: func() TokenType {
					if isFloat {
						return TokenFloat
					} else {
						return TokenInteger
					}
				}(),
				data: func() any {
					if isFloat {
						f, err := strconv.ParseFloat(str, 64)
						if err != nil {
							panic("Invalid `float`")
						}
						return f
					} else {
						in, err := strconv.ParseInt(str, 10, 64)
						if err != nil {
							panic("Invalid `Integer`")
						}
						return in
					}
				}(),
				start: start,
				end:   i,
			})
		default:
			m := fmt.Sprintf("Invalid character %q at Line: %d, Pos: %d", char, lineNum, i)
			panic(m)
		}
	}
	fmt.Println("Line no", lineNum)

}

func (l *Lexer) isDigit(s byte) bool {
	return s >= '0' && s <= '9'
}
