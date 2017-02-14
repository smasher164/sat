package sat

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	testDir = "testdata/Formulas"
)

type fCNF struct {
	cnf      CNF
	filename string
}

func fpaths() []string {
	f, err := os.Open(testDir)
	if err != nil {
		log.Fatalln(err)
	}
	filenames, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fpaths := []string{}
	for i := range filenames {
		fpath := filepath.Join(testDir, filenames[i])
		if filepath.Ext(fpath) == ".cnf" {
			fpaths = append(fpaths, fpath)
		}
	}
	return fpaths
}

func tofCNF(fpaths []string) []fCNF {
	testfCNFs := make([]fCNF, 0, len(fpaths))
	for _, fp := range fpaths {
		fr, err := os.Open(fp)
		if err != nil {
			log.Fatalln(err)
		}
		cnf, err := DIMACS(fr)
		if err != nil {
			log.Fatalln(err)
		}
		testfCNFs = append(testfCNFs, fCNF{cnf: cnf, filename: fr.Name()})
		fr.Close()
	}
	return testfCNFs
}

func cdpll(formula CNF, ch chan<- bool, stop <-chan struct{}) {
	cdpll_1 := func() bool {
		trail := make(map[Literal]bool)
		res, _ := DPLL(formula, trail)
		return res
	}
	select {
	case ch <- cdpll_1():
	case <-stop:
	}
}

func TestDPLL(t *testing.T) {
	testfCNFs := tofCNF(fpaths())

	t.Parallel()
	for _, fcnf := range testfCNFs {
		fcnf := fcnf
		t.Run(fcnf.filename, func(t *testing.T) {
			// defer func(filename string) {
			// 	if r := recover(); r != nil {
			// 		t.Errorf("panic: calling DPLL() for %v\n", filename)
			// 	}
			// }(fcnf.filename)

			res, stop := make(chan bool), make(chan struct{}, 1)
			go cdpll(fcnf.cnf, res, stop)
			var r bool
			select {
			case r = <-res:
			case <-time.After(3 * time.Second):
				stop <- struct{}{}
				t.Fatalf("test for %v timed out", fcnf.filename)
			}

			if want := fcnf.filename[len(fcnf.filename)-5]; want == 's' {
				if !r {
					t.Fatalf("Want: %v, Got: %v", want, "u")
				}
			} else if want == 'u' {
				if r {
					t.Fatalf("Want: %v, Got: %v", want, "s")
				}
			} else {
				log.Fatalf("Error parsing %v for satisfiability condition (s|u).\n", fcnf.filename)
			}
		})
	}
}
