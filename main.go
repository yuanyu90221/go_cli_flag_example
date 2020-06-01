package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// global var
var (
	operator    string
	showProcess bool
)

const (
	plus      string = "+"
	minus     string = "-"
	multplied string = "*"
	dividedBy string = "/"
)

// command set up
func init() {
	// set -o *
	flag.StringVar(&operator, "o", "+", "the math `operator`")
	// bool use -p=true
	flag.BoolVar(&showProcess, "p", false, "`show` process")
	flag.Usage = usage
}

// -h
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: math [options] [root]\n")
	fmt.Fprintf(os.Stderr, " Currently, there are four URI routes could be used:\n")
	flag.PrintDefaults()
}
func main() {
	// parse command
	flag.Parse()
	// read arg
	number1, _ := strconv.Atoi(flag.Arg(0))
	number2, _ := strconv.Atoi(flag.Arg(1))

	if showProcess {
		fmt.Printf("%d %s %d = ", number1, operator, number2)
	}
	// do operator
	switch operator {
	case plus:
		fmt.Printf("%d \n", number1+number2)
	case minus:
		fmt.Printf("%d \n", number1-number2)
	case multplied:
		fmt.Printf("%d \n", number1*number2)
	case dividedBy:
		fmt.Printf("%0.2f \n", float64(number1)/float64(number2))
	}
}
