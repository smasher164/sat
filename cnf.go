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

type StringDumper interface {
	DumpString() string
}

func (l Literal) DumpString() string {
	s := ""
	if !l.Truth {
		s += "¬"
	}
	s += strconv.Itoa(l.Ident)
	return s
}

func (c Clause) DumpString() string {
	s := "("
	if len(c) > 0 {
		s += c[0].DumpString()
	}
	for i := 1; i < len(c); i++ {
		s += " v " + c[i].DumpString()
	}
	s += ")"
	return s
}

func (c CNF) DumpString() string {
	var s string
	if len(c) > 0 {
		s = c[0].DumpString()
		for i := 1; i < len(c); i++ {
			s += " ∧ " + c[i].DumpString()
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

func sliceLast(trail []Literal) []Literal {
	if len(trail) == 0 {
		return trail
	}
	return trail[:len(trail)-1]
}

func DPLL(formula CNF, trail []Literal) (bool, []Literal) {
	// empty formula
	if len(formula) == 0 {
		return true, trail
	}
	for i := range formula {
		// empty clause
		if len(formula[i]) == 0 {
			return false, trail
		}
		// unit propagation
		if len(formula[i]) == 1 {
			// get literal from unit clause
			l := formula[i][0]
			// ensure only one representation of literal in trail
			trail = sliceLast(trail)
			trail = append(trail, l)
			return DPLL(Simplify(formula, l), trail)
		}
	}
	// choose literal for unit propagation, and greedily store in trail
	l := formula[0][0]
	trail = append(trail, l)
	if truth, trail := DPLL(Simplify(formula, l), trail); truth {
		return true, trail
	} else {
		trail = sliceLast(trail)
		l.Truth = !l.Truth
		trail = append(trail, l)
		return DPLL(Simplify(formula, l), trail)
	}
}

func not(l Literal) Literal {
	return Literal{l.Ident, !l.Truth}
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
