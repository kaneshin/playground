package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

type (
	// Dispatcher represents a management workers.
	Dispatcher struct {
		pool    chan *worker
		queue   chan interface{}
		workers []*worker
		wg      sync.WaitGroup
		quit    chan struct{}
	}

	// worker represents the worker that executes the job.
	worker struct {
		dispatcher *Dispatcher
		data       chan interface{}
		quit       chan struct{}
	}
)

const (
	maxWorkers = 3
	maxQueues  = 10000
)

// NewDispatcher returns a pointer of Dispatcher.
func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, maxWorkers),
		queue: make(chan interface{}, maxQueues),
		quit:  make(chan struct{}),
	}
	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := worker{
			dispatcher: d,
			data:       make(chan interface{}),
			quit:       make(chan struct{}),
		}
		d.workers[i] = &w
		w.start()
	}
	return d
}

// Add adds a given value to the queue of the dispatcher.
func (d *Dispatcher) Add(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

// Start starts the specified dispatcher but does not wait for it to complete.
func (d *Dispatcher) Start() {
	go func() {
		for {
			select {
			case v := <-d.queue:
				(<-d.pool).data <- v

			case <-d.quit:
				return
			}
		}
	}()
}

// Wait waits for the dispatcher to exit. It must have been started by Start.
func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

// Stop stops the dispatcher to execute. The dispatcher stops gracefully
// if the given boolean is false.
func (d *Dispatcher) Stop(immediately bool) {
	if !immediately {
		d.Wait()
	}

	d.quit <- struct{}{}
	for _, w := range d.workers {
		w.quit <- struct{}{}
	}
}

func (w *worker) start() {
	go func() {
		for {
			// register the current worker into the dispatch pool
			w.dispatcher.pool <- w

			select {
			case v := <-w.data:
				if str, ok := v.(string); ok {
					get(str)
				}

				w.dispatcher.wg.Done()

			case <-w.quit:
				return
			}
		}
	}()
}

func get(url string) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Goroutine:%d, URL:%s (%d bytes)", runtime.NumGoroutine(), url, len(body))
}

func main() {
	d := NewDispatcher()

	d.Start()
	for i := 0; i < 100; i++ {
		url := fmt.Sprintf("http://placehold.it/%dx%d", i, i)
		d.Add(url)
	}
	d.Wait()
}
