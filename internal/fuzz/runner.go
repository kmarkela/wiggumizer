package fuzz

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
)

type workUnit struct {
	endpoint, parameter string
	parBody             bool
	hi                  *historyparser.HistoryItem
}

// Worker represents a single worker that executes HTTP requests.
type worker struct {
	workQ    <-chan *workUnit
	res      chan<- *http.Response
	wordlist []string
	c        *http.Client
}

// func newWorker(hi <-chan *historyparser.HistoryItem, res chan<- *http.Response, wordlist []string, tr *http.Transport) *worker {
func newWorker(wq <-chan *workUnit, wordlist []string, tr *http.Transport) *worker {
	return &worker{
		workQ: wq,
		// res:      res,
		wordlist: wordlist,
		c:        &http.Client{Transport: tr},
	}
}

// Start the worker to execute HTTP requests.
func (w *worker) start(rateLimiter <-chan time.Time) {
	for wu := range w.workQ {

		for _, word := range w.wordlist {
			if rateLimiter != nil {
				<-rateLimiter // Wait for rate limit if provided
			}

			if wu.parBody {
				w.fuzzBody(wu, word)
				continue
			}

			w.fuzzGet(wu, word)

			// w.res <- w.doRequest(hi)

		}
	}
}

func (w *worker) fuzzGet(wu *workUnit, word string) {

	// prepare url
	oldParam := fmt.Sprintf("%s=%s", wu.parameter, wu.hi.Req.Parameters.Get[wu.parameter])
	newParam := fmt.Sprintf("%s=%s", wu.parameter, url.QueryEscape(word))
	endpoint := strings.Replace(wu.hi.Path, oldParam, newParam, 1)
	url := fmt.Sprintf("%s%s", wu.hi.Host, endpoint)

	w.doRequest(url, nil, wu.hi)

}

func (w *worker) fuzzBody(wu *workUnit, word string) {

	body := w.encodeBody(wu.hi.Req.ContentType, word, wu.parameter, wu.hi.Req.Parameters.Post)
	w.doRequest(wu.hi.Host+wu.hi.Path, body, wu.hi)
}

func (w *worker) doRequest(url string, body io.Reader, hi *historyparser.HistoryItem) error {

	req, err := http.NewRequest(hi.Method, url, body)
	if err != nil {
		return err
	}

	for _, v := range strings.Split(hi.Req.Headers, "\n") {
		keyVal := strings.Split(v, ":")
		if len(keyVal) != 2 {
			// TODO: log it
			continue
		}
		req.Header.Set(keyVal[0], strings.TrimSpace(keyVal[1]))

	}

	res, err := w.c.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	return nil

}

func (f *Fuzzer) Run(bh *historyparser.BrowseHistory) {

	// Create rate limiter if maxReq > 0
	var rateLimiter <-chan time.Time
	if f.maxReq > 0 {
		rateLimiter = time.Tick(time.Second / time.Duration(f.maxReq))
	}

	wq := make(chan *workUnit, f.workers)
	// results := make(chan string, f.workers)

	// Create and start workers
	var wg sync.WaitGroup
	for i := 0; i < f.workers; i++ {
		worker := newWorker(wq, f.wordlist, f.tr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.start(rateLimiter)
		}()
	}

	// distribute work
	for _, v := range bh.HistoryItems {

		endpoint := v.Host + strings.Split(v.Path, "?")[0]
		f.fuzzHistory.add(endpoint)

		for k := range v.Req.Parameters.Get {

			// skip parameters that were fuzzed alredy
			if f.fuzzHistory.h[endpoint].Contains("get-" + k) {
				continue
			}

			f.fuzzHistory.h[endpoint].Add("get-" + k)

			wq <- &workUnit{
				hi:        &v,
				endpoint:  endpoint,
				parameter: k,
			}

		}

		for k := range v.Req.Parameters.Post {

			// skip parameters that were fuzzed alredy
			if f.fuzzHistory.h[endpoint].Contains("post-" + k) {
				continue
			}

			f.fuzzHistory.h[endpoint].Add("post-" + k)

			wq <- &workUnit{
				hi:        &v,
				endpoint:  endpoint,
				parameter: k,
				parBody:   true,
			}

		}
	}

	close(wq)
	wg.Wait()

	// Close the results channel
	// close(results)

	// // Process results
	// for result := range results {
	// 	fmt.Println(result)
	// }

}
