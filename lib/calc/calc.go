package calc

import (
	"fmt"
	"math"
)

var functions = map[string]interface{}{
	"abs":  math.Abs,
	"ceil": math.Ceil,
	"cos":  math.Cos,
	"tan":  math.Tan,
	"log":  math.Log,
	"max":  math.Max,
	"min":  math.Min,
	"pow":  math.Pow,
	"sqrt": math.Sqrt,
}

func call(funcName string, args []float64) (float64, error) {
	f, ok := functions[funcName]
	if !ok {
		return 0, fmt.Errorf("unknown function %s", funcName)
	}
	switch f := f.(type) {
	case func() float64:
		return f(), nil
	case func(float64) float64:
		return f(args[0]), nil
	case func(float64, float64) float64:
		return f(args[0], args[1]), nil
	case func(float64, float64, float64) float64:
		return f(args[0], args[1], args[2]), nil
	default:
		return 0, fmt.Errorf("invalid function %s", funcName)
	}
}

func calculate(n *Node) (float64, error) {
	switch n.Kind {
	case addNode:
		left, err := calculate(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.Right)
		if err != nil {
			return 0, err
		}
		return left + right, nil
	case subNode:
		left, err := calculate(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.Right)
		if err != nil {
			return 0, err
		}
		return left - right, nil
	case mulNode:
		left, err := calculate(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.Right)
		if err != nil {
			return 0, err
		}
		return left * right, nil
	case divNode:
		left, err := calculate(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.Right)
		if err != nil {
			return 0, err
		}
		return left / right, nil
	case numNode:
		return n.Val, nil
	case funcNode:
		args := []float64{}
		for _, arg := range n.Args {
			val, err := calculate(arg)
			if err != nil {
				return 0, err
			}
			args = append(args, val)
		}
		return call(n.FuncName, args)
	}
	return 0, fmt.Errorf("unknown node type: %s", n.Kind)
}

// Calculate calculates expr
func Calculate(expr string) (float64, error) {
	tokens, err := Tokenize(expr)
	if err != nil {
		return 0, err
	}
	p := newParser(tokens)
	n, err := p.parse()
	if err != nil {
		return 0, err
	}
	return calculate(n)
}
