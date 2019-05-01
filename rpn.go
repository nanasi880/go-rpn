package rpn // import "go.nanasi880.dev/rpn"

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"go.nanasi880.dev/rpn/internal"
)

// RPN is Reverse Polish Notation object
type RPN struct {
	Tokens []string // Parsed expression token list
}

// Resolver is variable value resolver by name
type Resolver func(name string) (decimal.Decimal, error)

// Eval is evaluate expression
func (r *RPN) Eval(resolver Resolver) (decimal.Decimal, error) {

	var (
		zero  decimal.Decimal
		stack internal.Stack
	)
	for _, tok := range r.Tokens {

		if !r.isOperator(tok) {

			var (
				v   decimal.Decimal
				err error
			)
			if r.isVar(tok) {
				if resolver == nil {
					return zero, fmt.Errorf("nil resolver")
				}
				v, err = resolver(tok[1:])
			} else {
				v, err = tokenToDecimal(tok)
			}
			if err != nil {
				return zero, err
			}

			stack.Push(v)
			continue
		}

		a, b, ok := stack.Pop2()
		if !ok {
			return zero, fmt.Errorf("stack error")
		}

		v := r.eval(a, b, tok)
		stack.Push(v)
	}

	v, ok := stack.Pop()
	if !ok {
		return zero, fmt.Errorf("internal")
	}

	return v, nil
}

func (_ *RPN) isOperator(tok string) bool {

	switch tok {

	case "+", "-", "*", "/":
		return true
	}

	return false
}

func (_ *RPN) isVar(tok string) bool {
	return strings.HasPrefix(tok, "$")
}

func (_ *RPN) eval(a decimal.Decimal, b decimal.Decimal, op string) decimal.Decimal {

	switch op {

	case "+":
		return a.Add(b)

	case "-":
		return a.Sub(b)

	case "*":
		return a.Mul(b)

	case "/":
		return a.Div(b)

	}

	panic("internal")
}
