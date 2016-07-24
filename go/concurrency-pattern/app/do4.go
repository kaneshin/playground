// +build do4

package main

import (
	"log"

	"github.com/kaneshin/playground/go/concurrency-pattern/payload"
)

type (
	worker struct {
		pool chan chan payload.Instance
		ch   chan payload.Instance
		quit chan bool
	}
	dispatcher struct {
		maxWorkers int
		pool       chan chan payload.Instance
	}
)

var (
	queue chan payload.Instance
)

func newWorker(pool chan chan payload.Instance) worker {
	return worker{
		pool: pool,
		ch:   make(chan payload.Instance),
		quit: make(chan bool),
	}
}

// start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w worker) start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.pool <- w.ch

			select {
			case instance := <-w.ch:
				body, err := instance.Get()
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("Get %s  (%d bytes)\n", instance.URL, len(body))
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// stop signals the worker to stop listening for work requests.
func (w worker) stop() {
	go func() {
		w.quit <- true
	}()
}

func newDispatcher(maxWorkers int) *dispatcher {
	return &dispatcher{
		maxWorkers: maxWorkers,
		pool:       make(chan chan payload.Instance, maxWorkers),
	}
}

func (d *dispatcher) run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := newWorker(d.pool)
		worker.start()
	}
	go d.dispatch()
}

func (d *dispatcher) dispatch() {
	for {
		select {
		case instance := <-queue:
			// a job request has been received
			go func(instance payload.Instance) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				ch := <-d.pool

				// dispatch the job to the worker job channel
				ch <- instance
			}(instance)
		}
	}
}

func init() {
	queue = make(chan payload.Instance, numMaxQueues())
	dispatcher := newDispatcher(numMaxWorkers())
	dispatcher.run()
}

func do(instance payload.Instance) {
	queue <- instance
}
