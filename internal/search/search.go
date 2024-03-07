package search

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/kmarkela/Wiggumizeng/cmd"
	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
)

type worker struct {
	bh        *historyparser.BrowseHistory
	jobs      <-chan int
	results   chan<- int
	wg        *sync.WaitGroup
	reg       sParam
	caseInSen bool
}

type Searcher struct {
	sParam sParam
}

type sParam struct {
	req    sReg
	res    sReg
	method []sMatch
	conf   sConf
}

type sOutputType int

const (
	brief sOutputType = iota
	headers
	reqOnly
	full
)

type sConf struct {
	caseInSen bool
	output    sOutputType
}

type sReg struct {
	header      []sMatch
	contentType []sMatch
	body        []sMatch
}

type sMatch struct {
	value    string
	negative bool
}

func (s *Searcher) Search(bh *historyparser.BrowseHistory, numJobs int) {

	fmt.Print("Regexp Search. Type \"help\" for help or \"exit\" to exit \n")

	input := cmd.GetString("Type search query: ")

	switch input {
	case "help", "Help", "HELP":
		help()
	case "exit", "Exit", "EXIT":
		return
	case "":
		// just skip empty str
	default:
		s.sParam = parseInput(input)
		s.output(bh, s.doSearch(bh, numJobs))
	}

	// run self again
	s.Search(bh, numJobs)
}

func (s *Searcher) doSearch(bh *historyparser.BrowseHistory, numJobs int) []int {

	var wgW, wgR sync.WaitGroup
	var results = []int{}

	var jCh = make(chan int, numJobs)
	var rCh = make(chan int, numJobs)

	// listen for results
	wgR.Add(1)
	go func() {
		for i := range rCh {
			results = append(results, i)
		}
		wgR.Done()
	}()

	// start workers
	wi := worker{
		bh:      bh,
		jobs:    jCh,
		results: rCh,
		wg:      &wgW,
		reg:     s.sParam,
	}

	for w := 1; w <= numJobs; w++ {
		go wi.start()
	}

	// distribute work
	wgW.Add(numJobs)
	for i := 0; i <= len(bh.HistoryItems)-1; i++ {
		jCh <- i
	}
	close(jCh)

	wgW.Wait()
	close(rCh)
	wgR.Wait()

	return results
}

func (w *worker) start() {
	for i := range w.jobs {
		if checkMatch(&w.bh.HistoryItems[i], w.reg, w.caseInSen) {
			w.results <- i
		}
	}
	w.wg.Done()
}

func checkMatch(bi *historyparser.HistoryItem, reg sParam, ci bool) bool {
	// check Method
	if !regexMatch(bi.Method, reg.method, reg.conf.caseInSen) {
		return false
	}

	// check req
	if !regexMatch(bi.Req.Body, reg.req.body, reg.conf.caseInSen) || !regexMatch(bi.Req.Headers, reg.req.header, reg.conf.caseInSen) || !regexMatch(bi.Req.ContentType, reg.req.contentType, reg.conf.caseInSen) {
		return false
	}

	// check res
	if !regexMatch(bi.Res.Body, reg.res.body, reg.conf.caseInSen) || !regexMatch(bi.Res.Headers, reg.res.header, reg.conf.caseInSen) || !regexMatch(bi.Res.ContentType, reg.res.contentType, reg.conf.caseInSen) {
		return false
	}

	return true
}

func regexMatch(st string, m []sMatch, ci bool) bool {

	for _, v := range m {
		match, _ := regexp.MatchString(v.value, st)
		if ci {
			match, _ = regexp.MatchString("(?i)"+v.value, st)
		}

		// Found Negative or doesn't find positive
		if (match && v.negative) || (!match && !v.negative) {
			return false
		}
	}

	return true
}
