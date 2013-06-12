package batch

import (
	"net/http"
	"io/ioutil"
)

const concurrentDefault = 10

type entry struct {
	url string
	callback Callback
}

type batch struct {
	maxConcurrent int
	pool []*entry
}

type Callback func(url string, body string, err error)

func (b *batch) SetMaxConcurrent(maxConcurrent int) (previous int) {
	previous = b.maxConcurrent
	b.maxConcurrent = maxConcurrent
	return
}

func (b *batch) MaxConcurrent() (maxConcurrent int) {
	maxConcurrent = b.maxConcurrent
	return
}

func (b *batch) AddEntry(url string, callback Callback) {
	b.pool = append(b.pool, &entry{url, callback})
	return
}

func New() (*batch) {
	return &batch{concurrentDefault, []*entry{}}
}

func (b *batch) Run() {
	// create and fill our working queue
	queue := make(chan *entry, len(b.pool))
	for _, entry := range b.pool {
		queue <- entry
	}
	close(queue)
	var total_threads int
	if b.maxConcurrent <= len(b.pool) {
		total_threads = b.maxConcurrent
	} else {
		total_threads = len(b.pool)
	}
	waiters := make(chan bool, total_threads)
	var threads int
	for threads = 0; threads < total_threads; threads++ {
		go process(queue, waiters, threads)
	}
	for ; threads > 0; threads-- {
		<-waiters
	}
}

func process(queue chan *entry, waiters chan bool, thread_num int) {
	for entry := range queue {
		response, err := http.Get(entry.url)
		if err != nil {
			entry.callback(entry.url, "", err)
			continue
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			entry.callback(entry.url, "", err)
			continue
		}
		entry.callback(entry.url, string(contents), nil)
	}
	waiters <- true
}
