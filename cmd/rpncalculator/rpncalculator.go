package main

import (
	"strings"
	"fmt"
	"errors"
	"strconv"
	"bufio"
	"os"
)

func main() {
	calc := New()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		calc.evaluate(scanner.Text())
		if top, ok := calc.top(); ok {
			fmt.Println(top)
		}
	}

	if scanner.Err() != nil {
		fmt.Errorf("%v", scanner.Err())
	}
}

type Stack []float64

type operator func(calculator *RpnCalculator) error

type RpnCalculator struct {
	stack Stack
	operations map[string]operator
}

func (calculator *RpnCalculator) evaluate(input string) error {
	tokens := strings.Fields(input)
	for _, token := range tokens {
		if operator, ok := calculator.operations[token]; ok {
			err := operator(calculator)
			if err != nil {
				fmt.Printf("Error: %v", err)
				return err
			}
		} else {
			f, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return err
			} else {
				calculator.push(f)
			}
		}
	}
	return nil
}

func (calculator *RpnCalculator) top() (float64, bool) {
	if len(calculator.stack) > 0 {
		return calculator.stack[len(calculator.stack) - 1], true
	} else {
		return 0, false
	}
}

func (calculator *RpnCalculator) pop() (float64, bool) {
	s := calculator.stack
	if len(s) > 0 {
		v := s[len(s) - 1]
		calculator.stack = s[:len(s) - 1]
		return v, true
	} else {
		return 0, false
	}
}

func (calculator *RpnCalculator) push(value float64) {
	calculator.stack = append(calculator.stack, value)
}

func binop(op func(float64, float64) float64) operator {
	return func(calc *RpnCalculator) error {
		y, ok2 := calc.pop()
		x, ok1 := calc.pop()
		if ok1 && ok2 {
			res := op(x, y)
			calc.push(res)
			return nil
		} else {
			return errors.New("couldn't pop two operands")
		}
	}
}

func New() RpnCalculator {
	calc := RpnCalculator{operations: make(map[string]operator)}

	calc.operations["+"] = binop(func(x float64, y float64) float64 {
		return x + y
	})
	calc.operations["-"] = binop(func(x float64, y float64) float64 {
		return x - y
	})
	calc.operations["*"] = binop(func(x float64, y float64) float64 {
		return x * y
	})
	calc.operations["/"] = binop(func(x float64, y float64) float64 {
		return x / y
	})

	return calc
}