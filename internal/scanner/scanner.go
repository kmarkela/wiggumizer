package scanner

import (
	"sync"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
	"github.com/kmarkela/Wiggumizeng/internal/scanner/splugin"
)

type cInOut struct {
	jobs    chan int
	results chan map[string]splugin.Finding
}

type Scanner struct {
	checkers map[string]splugin.Checker
	verbose  bool
}

func NewScanner(v bool) *Scanner {

	var s = &Scanner{verbose: v}
	s.checkers = make(map[string]splugin.Checker)
	s.registerCheckers()

	// it will be posible to enable verbose only for specific checks in futher version
	for _, v := range s.checkers {
		v.SetVerbose(s.verbose)
	}
	return s

}

func (s *Scanner) Scan(bh *historyparser.BrowseHistory, fname string, numJobs int) {
	var wgW, wgR sync.WaitGroup
	var results = make(map[string][]splugin.Finding)

	var ch = cInOut{
		jobs:    make(chan int, numJobs),
		results: make(chan map[string]splugin.Finding, numJobs),
	}

	// listen for results
	wgR.Add(1)
	go func() {
		for i := range ch.results {
			for k, v := range i {
				results[k] = append(results[k], v)
			}
		}
		wgR.Done()
	}()

	// start workers
	for w := 1; w <= numJobs; w++ {
		go s.worker(bh, &ch, &wgW)
	}

	// distribute work
	wgW.Add(numJobs)
	for i := 0; i <= len(bh.HistoryItems)-1; i++ {
		ch.jobs <- i
	}
	close(ch.jobs)
	wgW.Wait()

	close(ch.results)
	wgR.Wait()

	// save MD report
	s.saveReport(fname, results, bh.ListOfHosts.Keys())
}

func (s *Scanner) worker(bh *historyparser.BrowseHistory, c *cInOut, wg *sync.WaitGroup) {
	for i := range c.jobs {
		c.results <- s.runAllcheckes(&bh.HistoryItems[i])
	}
	wg.Done()
}

func (s *Scanner) runAllcheckes(bi *historyparser.HistoryItem) map[string]splugin.Finding {
	var results = make(map[string]splugin.Finding)

	for _, v := range s.checkers {
		finding, found := v.Check(bi)
		if !found {
			continue
		}
		results[v.GetName()] = finding
	}

	return results
}

func (s *Scanner) RunAcheck(name string) {}
