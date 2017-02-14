package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smasher164/sat"
)

func dumpTrail(trail map[sat.Literal]bool) {
	for l, _ := range trail {
		fmt.Printf("%s ", l)
	}
	fmt.Println()
}

func main() {
	cnf, err := sat.DIMACS(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	trail := make(map[sat.Literal]bool)
	res, trail := sat.DPLL(cnf, trail)
	fmt.Println(res)
	dumpTrail(trail)
}
