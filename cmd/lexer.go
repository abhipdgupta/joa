package joa

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
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
		case char == '\n':
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
		case char == ' ':
			i = i + 1
		case char == '"':
			var requirdIndex = strings.Index(input[i+1:], "\"")
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
				kind:  TokenKeyword,
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
				kind:  TokenKeyword,
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
				kind:  TokenKeyword,
				data:  strToCheck,
				start: i,
				end:   i + 5,
			})
			i = i + 5
		case l.isDigit(char):
			var isFloat bool = false
			var start = i
			for l.isDigit(input[i]) && i < len(input) {
				i = i + 1
			}
			if input[i] == '.' {
				isFloat = true
				i = i + 1
				for l.isDigit(input[i]) && i < len(input) {
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
							panic("Invalid float")
						}
						return f
					} else {
						in, err := strconv.ParseFloat(str, 64)
						if err != nil {
							panic("Invalid float")
						}
						return in
					}
				}(),
				start: start,
				end:   i,
			})
		default:
			panic("Invalid character")
		}
	}
	fmt.Println("Line no", lineNum)

}

func (l *Lexer) isDigit(s byte) bool {
	return s >= '0' && s <= '9'
}
