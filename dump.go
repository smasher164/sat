package sat

import "strconv"

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
