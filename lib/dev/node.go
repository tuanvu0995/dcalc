package dev

import (
	"dcalc/lib/calc"
	"fmt"
)

type NodeOperator string

const (
	Add      NodeOperator = "+"
	Subtract NodeOperator = "-"
	Multiply NodeOperator = "*"
	Divide   NodeOperator = "/"
)

type Node struct {
	Tokens   []calc.Token `json:"tokens"`
	Buffer   string       `json:"buffer"`
	Result   float64      `json:"result"`
	Name     string       `json:"name"`
	Operator NodeOperator `json:"operator"`
	Done     bool         `json:"done"`
	Message  string       `json:"message"`
}

func (node *Node) Init(name string) {
	node.Tokens = []calc.Token{}
	node.Buffer = ""
	node.Result = 0.0
	node.Name = name
	node.Operator = Add
	node.Done = false
	node.Message = ""
}

// Checker

func (node *Node) IsEmpty() bool {
	return len(node.Tokens) == 0
}

// Action
func (node *Node) Add(value string) int {
	newBuffer := node.Buffer + value
	tokens, err := calc.Tokenize(newBuffer)
	if err != nil {
		return -1
	}

	node.Buffer = newBuffer
	node.Tokens = tokens
	return len(node.Buffer) - 1
}

func (node *Node) Pop() {
	if len(node.Buffer) > 0 {
		node.Buffer = node.Buffer[:len(node.Buffer)-1]
		tokens, err := calc.Tokenize(node.Buffer)
		if err != nil {
			fmt.Println("Token nizer error: ", err)
			return
		}
		node.Tokens = tokens
	} else {
		node.Buffer = ""
		node.Tokens = []calc.Token{}
	}
}

func (node *Node) Clear() {
	node.Buffer = ""
	node.Tokens = []calc.Token{}
}

func (node *Node) Calculate() (float64, bool) {
	val, err := calc.Calculate(node.Buffer)
	if err != nil {
		node.Message = err.Error()
		return 0.0, false
	}

	node.Result = val

	return val, true
}
