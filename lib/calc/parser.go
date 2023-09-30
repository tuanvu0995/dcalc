package calc

import (
	"fmt"
	"math"
	"strings"
)

type NodeKind string

const (
	addNode  NodeKind = "+"
	subNode  NodeKind = "-"
	mulNode  NodeKind = "*"
	divNode  NodeKind = "/"
	funcNode NodeKind = "func"
	numNode  NodeKind = "num"
)

type Node struct {
	Kind  NodeKind
	Left  *Node
	Right *Node

	FuncName string
	Args     []*Node

	Val float64
}

type Parser struct {
	Tokens []Token
	Index  int
}

func newParser(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, Index: 0}
}

func (p *Parser) numberNode() (*Node, error) {
	t := p.Tokens[p.Index]
	if t.Kind != numberToken {
		return nil, fmt.Errorf("expected a number: %s", t.Str)
	}
	p.Index++
	return &Node{Kind: numNode, Val: t.Val}, nil
}

func (p *Parser) constantNode(str string) (*Node, error) {
	constants := map[string]float64{
		"e":   math.E,
		"pi":  math.Pi,
		"phi": math.Phi,

		"sqrt2":   math.Sqrt2,
		"sqrte":   math.SqrtE,
		"sqrtpi":  math.SqrtPi,
		"sqrtphi": math.SqrtPhi,

		"ln2":    math.Ln2,
		"log2e":  math.Log2E,
		"ln10":   math.Ln10,
		"log10e": math.Log10E,
	}
	val, ok := constants[strings.ToLower(str)]
	if !ok {
		return nil, fmt.Errorf("unknown constant: %s", str)
	}
	p.Index++
	return &Node{Kind: numNode, Val: val}, nil
}

func argumentNumber(funcName string) (int, error) {
	f, ok := functions[funcName]
	if !ok {
		return 0, fmt.Errorf("unknown function: %s", funcName)
	}

	switch f.(type) {
	case func() float64:
		return 0, nil
	case func(float64) float64:
		return 1, nil
	case func(float64, float64) float64:
		return 2, nil
	case func(float64, float64, float64) float64:
		return 3, nil
	default:
		return 0, fmt.Errorf("invalid function: %s", funcName)
	}
}

func (p *Parser) functionNode(str string) (*Node, error) {
	funcName := strings.ToLower(str)
	num, err := argumentNumber(funcName)
	if err != nil {
		return nil, err
	}

	if p.consume(")") {
		if num != 0 {
			return nil, fmt.Errorf("%s should have argument(s)", funcName)
		}
		return &Node{Kind: funcNode, FuncName: funcName}, nil
	}

	args := []*Node{}

	n, err := p.add()
	if err != nil {
		return nil, err
	}
	args = append(args, n)

	for p.consume(",") {
		n, err := p.add()
		if err != nil {
			return nil, err
		}
		args = append(args, n)
	}
	if len(args) != num {
		return nil, fmt.Errorf("%s should have %d argument(s) but has %d arguments(s)",
			funcName, num, len(args))
	}
	p.consume(")")
	return &Node{Kind: funcNode, FuncName: funcName, Args: args}, nil
}

func (p *Parser) consume(s string) bool {
	t := p.Tokens[p.Index]
	if t.Kind != reservedToken || t.Str != s {
		return false
	}
	p.Index++
	return true
}

func (p *Parser) parse() (*Node, error) {
	return p.add()

}

func (p *Parser) insert(n *Node, f func() (*Node, error), kind NodeKind) (*Node, error) {
	left := n
	right, err := f()
	if err != nil {
		return n, err
	}
	return &Node{Kind: kind, Left: left, Right: right}, err
}

func (p *Parser) add() (*Node, error) {
	n, err := p.mul()
	if err != nil {
		return nil, err
	}

	for {
		if p.consume("+") {
			n, err = p.insert(n, p.mul, addNode)
			if err != nil {
				return nil, err
			}
		} else if p.consume("-") {
			n, err = p.insert(n, p.mul, subNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
}

func (p *Parser) mul() (*Node, error) {
	n, err := p.unary()
	if err != nil {
		return nil, err
	}

	for {
		if p.consume("*") {
			n, err = p.insert(n, p.unary, mulNode)
			if err != nil {
				return nil, err
			}
		} else if p.consume("/") {
			n, err = p.insert(n, p.unary, divNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
}

func (p *Parser) unary() (*Node, error) {
	if p.consume("+") {
		return p.primary()
	} else if p.consume("-") {
		return p.insert(&Node{Kind: numNode, Val: 0.0}, p.primary, subNode)
	}
	return p.primary()
}

func (p *Parser) primary() (*Node, error) {
	if p.consume("(") {
		n, err := p.add()
		if err != nil {
			return nil, err
		}
		p.consume(")")
		return n, nil
	}

	if p.Tokens[p.Index].Kind == identToken {
		str := p.Tokens[p.Index].Str
		p.Index++
		if p.Index < len(p.Tokens) && p.consume("(") {
			return p.functionNode(str)
		}
		p.Index--
		return p.constantNode(str)
	}
	return p.numberNode()
}
