package calc

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type TokenKind string

const (
	reservedToken TokenKind = "reserved"
	numberToken   TokenKind = "number"
	identToken    TokenKind = "ident"
	eosToken      TokenKind = "eos"
)

type Token struct {
	Kind TokenKind `json:"kind"`
	Val  float64   `json:"val"`
	Str  string    `json:"str"`
}

type InvalidTokenError struct {
	Input    string `json:"Input"`
	Position int    `json:"Position"`
}

func (e *InvalidTokenError) Error() string {
	curr := ""
	pos := e.Position
	for _, line := range strings.Split(e.Input, "\n") {
		len := len(line)
		curr += line + "\n"
		if pos < len {
			return curr + strings.Repeat(" ", pos) + "^ invalid token"
		}
		pos -= len + 1
	}
	return ""
}

const operators = "+-*/(),"

func isOperator(char rune) bool {
	for _, op := range operators {
		if char == op {
			return true
		}
	}
	return false
}

func numberPrefix(chars []rune, i *int, n int) (float64, error) {
	val := 0.0
	len := 0
	for *i < n {
		curr, err := strconv.ParseFloat(string(chars[*i-len:*i+1]), 64)
		if err != nil {
			break
		}
		val = curr
		len++
		*i++
	}
	if len > 0 {
		return val, nil
	}
	return 0, errors.New("expected a number")
}

func isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isAlNum(char rune) bool {
	return isAlpha(char) || (char >= '0' && char <= '9')
}

func Tokenize(input string) ([]Token, error) {
	chars := []rune(input)
	i := 0
	n := len(chars)
	tokens := []Token{}
	for i < n {
		char := chars[i]
		if unicode.IsSpace(char) {
			i++
			continue
		}

		if isAlpha(char) {
			start := i
			i++
			for i < n && isAlNum(chars[i]) {
				i++
			}
			tokens = append(tokens,
				Token{Kind: identToken, Str: string(chars[start:i])})
			continue
		}

		if isOperator(char) {
			tokens = append(tokens, Token{Kind: reservedToken, Str: string(char)})
			i++
			continue
		}

		if val, err := numberPrefix(chars, &i, n); err == nil {
			tokens = append(tokens, Token{Kind: numberToken, Val: val})
			continue
		}

		return nil, &InvalidTokenError{Input: input, Position: i}
	}
	tokens = append(tokens, Token{Kind: eosToken})
	return tokens, nil
}
