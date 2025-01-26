package main

import (
	"flag"
	"log"
)

type operatorType int

const(
	openBracket operatorType =iota
	closedBracket 
	otherOperator
)

var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

type stack struct{
	elems []rune
}

func (s *stack) push(e rune){
	s.elems= append(s.elems,e)
}

func (s *stack) pop() *rune{
	if len(s.elems) == 0{
		return nil
	}
	e:= s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return &e
}

func getOperatorType(op rune) operatorType{
	for o,c := range bracketPairs{
		switch op {
		case o:
			return openBracket
		case c:
			return closedBracket
		}
	}
	return otherOperator
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	var s stack
	for _,e := range expr{
		switch getOperatorType(e) {
		case openBracket:
			s.push(e)
		case closedBracket:
			j:= s.pop()
			if j==nil || e!=bracketPairs[*j]{
				return false
			}
		}
	}
	return len(s.elems) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool){ 
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
