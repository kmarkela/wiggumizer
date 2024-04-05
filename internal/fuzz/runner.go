package fuzz

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/kmarkela/Wiggumizeng/internal/historyparser"
)

// Worker represents a single worker that executes HTTP requests.
type worker struct {
	hi       <-chan *historyparser.HistoryItem
	res      chan<- *http.Response
	wordlist []string
	tr       *http.Transport
}

func newWorker(hi <-chan *historyparser.HistoryItem, res chan<- *http.Response, wordlist []string, tr *http.Transport) *worker {
	return &worker{
		hi:       hi,
		res:      res,
		wordlist: wordlist,
		tr:       tr,
	}
}

// Start the worker to execute HTTP requests.
func (w *worker) Start(rateLimiter <-chan time.Time) {
	for hi := range w.hi {
		for _, word := range w.wordlist {
			if rateLimiter != nil {
				<-rateLimiter // Wait for rate limit if provided
			}
			w.res <- w.doRequest(hi)
		}
	}
}

func (w *worker) doRequest(hi *historyparser.HistoryItem) string {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error fetching %s: %s", url, err)
	}
	defer response.Body.Close()

	// Here you can process the response as needed
	// For simplicity, let's just return the response status code
	return fmt.Sprintf("Response from %s: %s", url, response.Status)
}

func (f *Fuzzer) Run(bh *historyparser.BrowseHistory) {

	// Create rate limiter if maxReq > 0
	var rateLimiter <-chan time.Time
	if f.maxReq > 0 {
		rateLimiter = time.Tick(time.Second / time.Duration(f.maxReq))
	}

	hi := make(chan *historyparser.HistoryItem, f.workers)

	// Create and start workers
	var wg sync.WaitGroup
	for i := 0; i < f.workers; i++ {
		worker := newWorker(hi)
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.Start(rateLimiter)
		}()
	}

	// distribute work
	for _, v := range bh.HistoryItems {

		endpoint := v.Host + strings.Split(v.Path, "?")[0]
		f.fuzzHistory.add(endpoint)

		for k, p := range v.Req.Parameters.Get {

			// skip parameters that were fuzzed alredy
			if f.fuzzHistory.h[endpoint].Contains(k) {
				continue
			}

			f.fuzzHistory.h[endpoint].Add(k)

			hi <- &v

		}

		for k, p := range v.Req.Parameters.Post {

			// TODO: make history Method aware
			// skip parameters that were fuzzed alredy
			if f.fuzzHistory.h[endpoint].Contains(k) {
				continue
			}

			f.fuzzHistory.h[endpoint].Add(k)

			hi <- &v

		}
	}

	close(hi)
	wg.Wait()

}
