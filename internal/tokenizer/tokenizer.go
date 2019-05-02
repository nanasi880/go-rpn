package tokenizer

import "fmt"

type tokenType int

const (
	tokenTypeNone tokenType = iota
	tokenTypeNumber
	tokenTypeSymbol
	tokenTypeVar
)

// Parse is parse expr in tokens
func Parse(expr string) ([]string, error) {

	var (
		token    []rune
		tokType  tokenType
		tokens   []string
		prevRune rune
		hexToken bool
	)
	flushToken := func() {
		if len(token) > 0 {
			tokens = append(tokens, string(token))
			token = token[:0]
			tokType = tokenTypeNone
			hexToken = false
		}
	}
	for i, c := range expr {

		if c == ' ' {
			flushToken()
			goto NEXT
		}

		switch tokType {

		case tokenTypeNone:
			tokType = getTokenType(c)
			token = append(token, c)

		case tokenTypeNumber:

			if isNumber(c, hexToken) {
				token = append(token, c)
				goto NEXT
			}

			if c == 'x' {
				if prevRune == '0' && !containsRune(token, 'x') {
					hexToken = true
					token = append(token, c)
					goto NEXT
				} else {
					return nil, fmt.Errorf("invalid character `x`: %d", i)
				}
			}

			if isSymbol(c) {
				flushToken()
				tokType = tokenTypeSymbol
				token = append(token, c)
				goto NEXT
			}

			return nil, fmt.Errorf("invalid character `%c`: %d", c, i)

		case tokenTypeSymbol:
			// シンボルは複数ワード連結されないので必ずフラッシュする
			flushToken()

			tokType = getTokenType(c)
			token = append(token, c)
			goto NEXT

		case tokenTypeVar:

			if isSymbol(c) {
				flushToken()
				tokType = tokenTypeSymbol
				token = append(token, c)
				goto NEXT
			}

			token = append(token, c)
			goto NEXT
		}

	NEXT:
		prevRune = c
	}

	flushToken()
	return tokens, nil
}

func containsRune(runes []rune, c rune) bool {
	for _, v := range runes {
		if v == c {
			return true
		}
	}
	return false
}

func isNumber(c rune, hex bool) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if !hex {
		return false
	}

	if c >= 'a' && c <= 'f' {
		return true
	}
	if c >= 'A' && c <= 'F' {
		return true
	}

	return false
}

func isSymbol(c rune) bool {
	switch c {
	case '+', '-', '*', '/', '(', ')':
		return true
	}
	return false
}

func getTokenType(c rune) tokenType {
	if isNumber(c, false) {
		return tokenTypeNumber
	}
	if isSymbol(c) {
		return tokenTypeSymbol
	}
	return tokenTypeVar
}
