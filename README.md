# sat
Explores various implementations and techniques involving SAT solvers.

Currently, master has a traditional recursive implementation of DPLL, see [here](http://www.cs.miami.edu/home/geoff/Courses/CSC648-12S/Content/DPLL.shtml) for an explanation. It will fail on test data that has a large number of literals/clauses, because of very deep recursion in the DPLL procedure. Possible improvements include:

1. Output a solution to the formula if its satisfiable.
	1. Recording a solution may require a trail-type data structure that's preserved when backtracking.
2. Use an iterative algorithm. This may either require maintaining a stack for backtracking, or using the trail-type data structure to undo a change to the formula.
3. Include more heuristics when choose a literal in unit propagation.
4. Concurrency/streaming
	1. Naively splitting formula when simplifying (order shouldn't matter).
	2. Trying both v and ¬v at the same time.
	3. Avoid storing the entire formula in memory at once. Use channel of literals/clauses instead of a slice.
5. Avoid recopying entire formula when simplifying.
6. etc...