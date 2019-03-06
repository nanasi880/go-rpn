# go-rpn

## Simply
```go
package main

import (
    "fmt"
    
    "github.com/shopspring/decimal"
    "go.nanasi880.dev/rpn"
)

func main() {
	const expr = `5 * (3 + 7) + 50`
	r, err := rpn.Parse(expr)
	if err != nil {
		panic(err)
	}

	v, err := r.Eval(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(v.String())
	// Output:
	// 100
}
```

## Variable value

```go
package main

import (
    "fmt"
    
    "github.com/shopspring/decimal"
    "go.nanasi880.dev/rpn"
)

func main() {
	const expr = `5 * (3 + 7) + 50 + x`
	r, err := rpn.Parse(expr)
	if err != nil {
		panic(err)
	}

	v, err := r.Eval(func(name string) (decimal.Decimal, error) {
		if name == "x" {
			return decimal.New(10, 0), nil
		}
		return decimal.Zero, fmt.Errorf("undefined variable: %s", name)
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(v.String())
	// Output:
	// 110
}
```
