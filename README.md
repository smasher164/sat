# sat
Explores various implementations and techniques involving SAT solvers.

Currently, master has a traditional recursive implementation of DPLL, see [__here__](http://www.cs.miami.edu/home/geoff/Courses/CSC648-12S/Content/DPLL.shtml) for an explanation. It will fail on test data that has a large number of literals/clauses, because of very deep recursion in the DPLL procedure. Possible improvements include:

- [x] Output a solution to the formula if its satisfiable.
	[__Trail structure preserved through recursion__](https://concurrency.cs.uni-kl.de/documents/albert_schimpf_bachelors_thesis.pdf)
	- [x] Recording a solution may require a trail-type data structure that's preserved when backtracking.
- [ ] Use an iterative algorithm. This may either require maintaining a stack for backtracking, or using the trail-type data structure to undo a change to the formula.
- [ ] Include more heuristics when choose a literal in unit propagation.
- [ ] Concurrency/streaming
	- [ ] Naively splitting formula when simplifying (order shouldn't matter).
	- [ ] Trying both v and ¬v at the same time.
	- [ ] Avoid storing the entire formula in memory at once. Use channel of literals/clauses instead of a slice.
- [ ] Avoid recopying entire formula when simplifying.
	etc...