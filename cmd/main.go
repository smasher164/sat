package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smasher164/sat"
)

func dumpRes(trail []sat.Literal) {
	for i := range trail {
		fmt.Printf("%s ", trail[i].DumpString())
	}
	fmt.Println()
}

func main() {
	cnf, err := sat.DIMACS(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	sol, res := sat.DPLL(cnf)
	fmt.Println(sol)
	dumpRes(res)
}
