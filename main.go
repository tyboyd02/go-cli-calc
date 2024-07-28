package main

import (
	"fmt"
	"os"

	"github.com/tyboyd02/go-cli-calc/calculator"
	"github.com/tyboyd02/go-cli-calc/parser"
	"github.com/tyboyd02/go-cli-calc/utils"
)

const (
	colorRed   = "\033[0;31m"
	colorReset = "\033[0m"
)

const (
	usage string = `This program is a calculator that follows PEMDAS
	
Supported Symbols: (, ), ^, *, /, +, -

Usage examples: 
	go-cli-calc 2+2*2
		OUTPUT: 6
		
	go-cli-calc "4.2 + 18 / (9 - 3)"
		OUTPUT: 7.2

	go-cli-calc "(1-.5)+2*9+2"
		OUTPUT: 20.5
`
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Fprint(os.Stdout, usage)
		os.Exit(0)
	case 1:
		helpFlags := map[string]bool{
			"-h":     true,
			"--help": true,
			"-help":  true,
			"help":   true,
		}

		if helpFlags[args[0]] {
			fmt.Fprint(os.Stdout, usage)
			os.Exit(0)
		}
	default:
		fmt.Fprintln(os.Stderr, "To many arguments")
		os.Exit(1)
	}

	equation := args[0]
	tokenizedEquation, err := utils.TokenizeEquation(equation)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError:%s Invalid equation\n", colorRed, colorReset)
		fmt.Fprintf(os.Stderr, "%sError:%s %v\n", colorRed, colorReset, err)
		return
	}
	postfixE := parser.ConvertToReversePolishNotation(tokenizedEquation)
	sum, err := calculator.SolvePostfix(postfixE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError:%s Invalid equation\n", colorRed, colorReset)
		fmt.Fprintf(os.Stderr, "%sError:%s %v\n", colorRed, colorReset, err)
		return
	}

	fmt.Println(sum)
}
