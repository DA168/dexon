package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dexon-foundation/dexon/core/vm/sqlvm/ast"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/parser"
)

func main() {
	var detail bool
	flag.BoolVar(&detail, "detail", false, "print struct detail")

	flag.Parse()

	n, err := parser.Parse([]byte(flag.Arg(0)))
	fmt.Printf("detail: %t\n", detail)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:\n%+v\n", err)
		os.Exit(1)
	}
	ast.PrintAST(os.Stdout, n, "  ", detail)
}
