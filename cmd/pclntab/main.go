package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vanstee/pclntab"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("exepcted at least one argument")
	}

	symbolTable, err := pclntab.ParsePclntabPath(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse pclntab: %v", err)
	}

	signal.Ignore(syscall.SIGPIPE)

	for _, f := range symbolTable.Funcs {
		fmt.Printf("%016x %s\n", f.Sym.Value, f.Sym.Name)
	}
}
