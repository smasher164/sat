package sat

import "strconv"

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

func (l Literal) String() string {
	s := ""
	if !l.Truth {
		s += "¬"
	}
	s += strconv.Itoa(l.Ident)
	return s
}

func (c Clause) String() string {
	s := "("
	if len(c) > 0 {
		s += c[0].String()
	}
	for i := 1; i < len(c); i++ {
		s += " v " + c[i].String()
	}
	s += ")"
	return s
}

func (c CNF) String() string {
	var s string
	if len(c) > 0 {
		s = c[0].String()
		for i := 1; i < len(c); i++ {
			s += " ∧ " + c[i].String()
		}
	}
	return s
}

// Makes a deep copy of the CNF
func (c CNF) clone() CNF {
	to := make(CNF, len(c), cap(c))
	for i := range to {
		to[i] = make(Clause, len(c[i]), cap(c[i]))
		copy(to[i], c[i])
	}
	return to
}

func DPLL(formula CNF) bool {
	// empty formula
	if len(formula) == 0 {
		return true
	}
	for i := range formula {
		// empty clause
		if len(formula[i]) == 0 {
			return false
		}
		// unit propagation
		if len(formula[i]) == 1 {
			return DPLL(Simplify(formula, formula[i][0]))
		}
	}
	if v := formula[0][0]; DPLL(Simplify(formula, v)) {
		return true
	} else {
		v.Truth = !v.Truth
		return DPLL(Simplify(formula, v))
	}
}

func Simplify(formula CNF, p Literal) CNF {
	f := formula.clone()
	for i := 0; i < len(f); i++ {
		for j := 0; i >= 0 && j < len(f[i]); j++ {
			if f[i][j].Ident == p.Ident {
				if f[i][j].Truth == p.Truth {
					// delete clause
					i--
					f = append(f[:i+1], f[i+2:]...)
				} else {
					// delete element
					f[i] = append(f[i][:j], f[i][j+1:]...)
				}
			}
		}
	}
	return f
}
