package rpn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func tokenToDecimal(tok string) (decimal.Decimal, error) {

	if strings.HasPrefix(tok, "0x") {
		v, err := strconv.ParseUint(tok, 0, 64)
		if err != nil {
			return decimal.Zero, err
		}
		return decimal.NewFromString(fmt.Sprint(v))
	}

	return decimal.NewFromString(tok)
}
