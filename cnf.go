package sat

// A Literal is represented by an identifier (int) and a truth value (bool).
type Literal struct {
	Ident int
	Truth bool
}

// A Clause is a disjunction of Literals. In conjunctive normal form (CNF),
// the structure implies that the literals are OR'ed together, like (a ∨ b ∨ ¬c).
type Clause []Literal

// A CNF is a conjunction of Clauses in conjunctive normal form,
// structured as clauses AND'ed together, like (a ∨ b) ∧ (b ∨ ¬c).
type CNF []Clause

// Makes a deep copy of the CNF
func (c CNF) clone() CNF {
	to := make(CNF, len(c), cap(c))
	for i := range to {
		to[i] = make(Clause, len(c[i]), cap(c[i]))
		copy(to[i], c[i])
	}
	return to
}

type solver struct {
	formula CNF
	res     map[Literal]bool
	sat     bool
}

func DPLL(formula CNF) (bool, []Literal) {
	s := &solver{
		formula: formula,
		res:     make(map[Literal]bool),
	}
	s.dpll()

	keys := make([]Literal, len(s.res))
	i := 0
	for k := range s.res {
		keys[i] = k
		i++
	}
	return s.sat, keys
}

func (s *solver) dpll() {
	// empty formula
	if len(s.formula) == 0 {
		s.sat = true
		return
	}
	for i := range s.formula {
		// empty clause
		if len(s.formula[i]) == 0 {
			s.sat = false
			return
		}
		// unit propagation
		if len(s.formula[i]) == 1 {
			// get literal from unit clause
			l := s.formula[i][0]
			// update the value of the literal in the result
			delete(s.res, not(l))
			s.res[l] = true
			s.simplify(l)
			s.dpll()
			return
		}
	}
	// choose literal for unit propogation, and greedily store in res
	l := s.formula[0][0]
	s.res[l] = true
	fcopy := s.formula.clone()
	s.simplify(l)
	if s.dpll(); s.sat {
		return
	} else {
		// fmt.Printf("l: %v, last: %v\n", l, s.res[len(s.res)-1])
		delete(s.res, l)
		s.formula = fcopy
		l.Truth = !l.Truth
		s.res[l] = true
		s.simplify(l)
		s.dpll()
		return
	}
}

func not(l Literal) Literal {
	return Literal{l.Ident, !l.Truth}
}

func (s *solver) simplify(p Literal) {
	for i := 0; i < len(s.formula); i++ {
		for j := 0; i >= 0 && j < len(s.formula[i]); j++ {
			if s.formula[i][j].Ident == p.Ident {
				if s.formula[i][j].Truth == p.Truth {
					// delete clause
					i--
					s.formula = append(s.formula[:i+1], s.formula[i+2:]...)
				} else {
					// delete element
					s.formula[i] = append(s.formula[i][:j], s.formula[i][j+1:]...)
				}
			}
		}
	}
}
