package sat

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func DIMACS(r io.Reader) (CNF, error) {
	scanner := bufio.NewScanner(r)
	var cnf CNF
	var tmpCl Clause
	for scanner.Scan() {
		s := scanner.Text()
		switch {
		case len(s) == 0:
		case s[0] == 'c':
		case s[0] == 'p':
			problem := strings.Fields(s)
			if len(problem) != 4 {
				return cnf, errors.New("invalid problem statement")
			}
			if capVar, err := strconv.Atoi(problem[2]); err != nil {
				return cnf, err
			} else {
				// fmt.Println("Number of variables", capVar)
				tmpCl = make(Clause, 0, capVar)
			}
			if capCl, err := strconv.Atoi(problem[3]); err != nil {
				return cnf, err
			} else {
				// fmt.Println("Number of clauses", capCl)
				cnf = make(CNF, 0, capCl)
			}
		case tmpCl != nil:
			literals := strings.Fields(s)
			for _, l := range literals {
				i, err := strconv.Atoi(l)
				if err != nil {
					return cnf, err
				}
				if i == 0 {
					// append tmpCl slice to cnf
					cnf = append(cnf, tmpCl)
					tmpCl = make([]Literal, 0, cap(tmpCl))
				} else if i > 0 {
					tmpCl = append(tmpCl, Literal{i, true})
				} else {
					tmpCl = append(tmpCl, Literal{-i, false})
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return cnf, err
	}
	return cnf, nil
}
