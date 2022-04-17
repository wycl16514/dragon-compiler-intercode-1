package main

import (
	"fmt"
	//	"io"
	//	"lexer"
	//	"simple_parser"
	"inter"
	"lexer"
)

func main() {
	/*
		source := "{int x; char y; {bool y; x; y;} x; y;}"
		my_lexer := lexer.NewLexer(source)
		parser := simple_parser.NewSimpleParser(my_lexer)
		err := parser.Parse()
		if err == io.EOF || err == nil {
			fmt.Println("\nparsing success")
		} else {
			fmt.Println("source is ilegal : ", err)
		}*/

	expr_type := inter.NewType("int", lexer.BASIC, 4)
	id_a := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "a"), expr_type)
	id_b := inter.NewID(1, lexer.NewTokenWithString(lexer.ID, "b"), expr_type)
	arith, _ := inter.NewArith(1, lexer.NewTokenWithString(lexer.PLUS, "+"), id_a, id_b)

	//expr := arith.Gen()
	op := arith.Reduce()
	r := op.ToString()

	fmt.Println("\nOp node return temperay register: ", r)
}
