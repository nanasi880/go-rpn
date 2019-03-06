package rpn

import "go.nanasi880.dev/rpn/internal/tokenizer"

// Parse is parse the expression
func Parse(expr string) (*RPN, error) {

	tokens, err := tokenizer.Parse(expr)
	if err != nil {
		return nil, err
	}

	return newParser(tokens).parse()
}
