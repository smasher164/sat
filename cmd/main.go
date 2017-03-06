package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smasher164/sat"
)

func dumpTrail(trail []sat.Literal) {
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
	trail := []sat.Literal{}
	res, trail := sat.DPLL(cnf, trail)
	fmt.Println(res)
	dumpTrail(trail)
}
