package joa

import "fmt"

//go:generate stringer -type=TokenType
type TokenType string

const (
	TokenOpenBrace  TokenType = "TokenOpenBrace"  // {
	TokenCloseBrace TokenType = "TokenCloseBrace" // }

	TokenOpenBracket  TokenType = "TokenOpenBracket"  // [
	TokenCloseBracket TokenType = "TokenCloseBracket" // ]

	TokenColon TokenType = "TokenColon" // :
	TokenComma TokenType = "TokenComma" // ,

	TokenString  TokenType = "TokenString"  // anything between " and "
	TokenInteger TokenType = "TokenInteger" // a integer
	TokenFloat   TokenType = "TokenFloat"   // float

	TokenKeyword TokenType = "TokenKeyword" // eg. null, true , false
)

type Token struct {
	kind TokenType

	data any

	start int

	end int
}

var _ fmt.Stringer = (*Token)(nil)

func (t Token) String() string {
	return fmt.Sprintf("Token(kind=%v, data=%v, start=%v, end=%v)", t.kind, t.data, t.start, t.end)
}
