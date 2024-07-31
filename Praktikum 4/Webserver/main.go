package main

import (
	"flag"
	"net/http"
	"strconv"
	"strings"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18



// / UPN calculator
// / only support + - * /
// / only supports 2 operands and 1 operator as of now
type Calculation struct {
	rValue   Value
	lValue   Value
	operator rune
}

// 1 1 + 2 1 + +

func calculationParser(s string) Value {
	tokens := strings.Fields(s)
	stack := []Value{}

	for _, token := range tokens {
		if number, err := strconv.Atoi(token); err == nil {
			stack = append(stack, Number{value: number})
		} else {
			if len(stack) < 2 {
				panic("Insufficient values in the expression")
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			calc := Calculation{lValue: left, rValue: right, operator: rune(token[0])}
			stack = append(stack, calc)
		}
	}

	if len(stack) != 1 {
		panic("Error in calculation: stack has multiple elements")
	}
	return stack[0]
}

func (c Calculation) evaluate() int {
	switch c.operator {
	case '+':
		return c.lValue.evaluate() + c.rValue.evaluate()
	case '-':
		return c.lValue.evaluate() - c.rValue.evaluate()
	case '*':
		return c.lValue.evaluate() * c.rValue.evaluate()
	case '/':
		if c.rValue.evaluate() != 0 {
			panic("Division by zero")
		}
		return c.lValue.evaluate() / c.rValue.evaluate()
	}
	return 0
}

type Value interface {
	evaluate() int
}

type Number struct {
	value int
}

func (n Number) evaluate() int {
	return n.value
}

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		calc := r.URL.Query().Get("calc")
		result := calculationParser(calc).evaluate()
		w.Write([]byte(strconv.Itoa(result)))
	})

	http.ListenAndServe(*addr, nil)

}
