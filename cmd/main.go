package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smasher164/sat"
)

func main() {
	cnf, err := sat.DIMACS(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sat.DPLL(cnf))
}
