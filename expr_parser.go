package rpn

import (
	"fmt"
)

type parser struct {
	tokens []string
	rpn    *RPN
	i      int
}

func newParser(tokens []string) *parser {
	return &parser{
		tokens: tokens,
		rpn:    new(RPN),
	}
}

func (c *parser) parse() (rpn *RPN, err error) {

	defer func() {
		switch r := recover().(type) {
		case nil:

		case error:
			err = r
		default:
			err = fmt.Errorf("%v", r)
		}
	}()

	c.parseExpression()
	return c.rpn, nil
}

func (c *parser) parseExpression() {

	c.parseTerm()

	token := c.currentToken()
	for token == "+" || token == "-" {
		c.next()

		c.parseTerm()
		c.rpn.Tokens = append(c.rpn.Tokens, token)
		token = c.currentToken()
	}
}

func (c *parser) parseTerm() {

	c.parseDivTerm()

	token := c.currentToken()
	for token == "*" {
		c.next()

		c.parseDivTerm()
		c.rpn.Tokens = append(c.rpn.Tokens, token)
		token = c.currentToken()
	}
}

func (c *parser) parseDivTerm() {

	c.parseFactor()

	token := c.currentToken()
	for token == "/" {
		c.next()

		c.parseFactor()
		c.rpn.Tokens = append(c.rpn.Tokens, token)
		token = c.currentToken()
	}
}

func (c *parser) parseFactor() {

	token := c.currentToken()
	if token == "(" {
		c.next()
		c.parseExpression()
		c.next()
	} else {
		c.parseNumber()
	}
}

func (c *parser) parseNumber() {

	token := c.currentToken()

	if c.isNumber(token) {

		if c.isSign(token) {
			c.next()
			token += c.currentToken()
		}

		_, err := tokenToDecimal(token)
		if err != nil {
			panic(err)
		}

	} else {
		token = "$" + token
	}

	c.rpn.Tokens = append(c.rpn.Tokens, token)
	c.next()
}

func (c *parser) next() {

	if c.i < len(c.tokens) {
		c.i++
		return
	}

	err := fmt.Errorf("finish")
	panic(err)
}

func (c *parser) currentToken() string {
	if c.i < len(c.tokens) {
		return c.tokens[c.i]
	}
	return ""
}

func (c *parser) isSign(t string) bool {

	if t[0] == '+' || t[0] == '-' {
		return true
	}
	return false
}

func (_ *parser) isNumber(t string) bool {

	c := t[0]
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
