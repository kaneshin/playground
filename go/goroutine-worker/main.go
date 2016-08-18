package main

import (
	"fmt"
	"sync"
	"time"
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

	// DispatcherOptions specifies the optional parameters to the Dispatcher.
	DispatcherOptions struct {
		WorkersCount        int
		QueueCapacity       int
		WorkerQueueCapacity int
		WorkerWaitDuration  time.Duration
		Func                func(interface{})
	}

	// worker represents the worker that executes the job.
	worker struct {
		dispatcher *Dispatcher
		data       chan interface{}
		signal     chan struct{}
		wait       time.Duration
		fn         func(interface{})
		quit       chan struct{}
	}
)

// NewDispatcher returns a pointer of Dispatcher.
func NewDispatcher(opt DispatcherOptions) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, opt.WorkersCount),
		queue: make(chan interface{}, opt.QueueCapacity),
		quit:  make(chan struct{}),
	}
	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		w := worker{
			dispatcher: d,
			wait:       opt.WorkerWaitDuration,
			data:       make(chan interface{}, opt.WorkerQueueCapacity),
			quit:       make(chan struct{}),
		}
		if cap(w.data) > 1 {
			w.signal = make(chan struct{})
		}
		d.workers[i] = &w
	}
	if opt.Func != nil {
		d.SetFunc(opt.Func)
	}
	return d
}

// Add adds a given value to the queue of the dispatcher.
func (d *Dispatcher) Add(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

// SetFunc sets a function to dispatch dequeued values.
func (d *Dispatcher) SetFunc(f func(interface{})) {
	for _, w := range d.workers {
		w.fn = f
	}
}

// Start starts the specified dispatcher but does not wait for it to complete.
func (d *Dispatcher) Start() {
	for _, w := range d.workers {
		w.start()
	}
	go func() {
		for {
			select {
			case v, ok := <-d.queue:
				if !ok {
					return
				}

				worker := <-d.pool
				worker.data <- v
				if cap(worker.data) > 1 {
					worker.signal <- struct{}{}
				}

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

// Destroy stops the dispatcher to execute and closes all channels.
// The dispatcher stops gracefully if the given boolean is false.
func (d *Dispatcher) Destroy(immediately bool) {
	d.Stop(immediately)

	close(d.queue)
	for _, w := range d.workers {
		if w.signal != nil {
			close(w.signal)
		}
		close(w.data)
	}
}

func (w *worker) newTimer() *time.Timer {
	t := time.AfterFunc(w.wait, func() func() {
		var l sync.Mutex
		return func() {
			l.Lock()
			defer l.Unlock()

			count := len(w.data)
			if count == 0 {
				return
			}

			list := make([]interface{}, count)
			for i := 0; i < count; i++ {
				list[i] = <-w.data
			}
			if w.fn != nil {
				w.fn(list)
			}

			w.dispatcher.wg.Add(-count)
		}
	}())

	if !t.Stop() {
		<-t.C
	}

	return t
}

func (w *worker) startSingle() {
	for {
		// register the current worker into the dispatch pool
		w.dispatcher.pool <- w

		select {
		case v, ok := <-w.data:
			if !ok {
				// receive a close signal
				return
			}

			if w.fn != nil {
				w.fn(v)
			}
			w.dispatcher.wg.Done()

		case <-w.quit:
			// receive a signal to stop
			return
		}
	}
}

func (w *worker) startMulti() {
	t := w.newTimer()
	for {
		// register the current worker into the dispatch pool.
		w.dispatcher.pool <- w

		select {
		case _, ok := <-w.signal:
			if !ok {
				// receive a close signal
				return
			}
			if len(w.data) == cap(w.data) {
				t.Reset(0)
			} else {
				t.Reset(w.wait)
			}

		case <-w.quit:
			// receive a signal to stop
			return
		}
	}
}

func (w *worker) start() {
	if cap(w.data) > 1 {
		go w.startMulti()
	} else {
		go w.startSingle()
	}
}

func main() {
	opt := DispatcherOptions{
		WorkersCount:        30,
		QueueCapacity:       100000,
		WorkerQueueCapacity: 10,
		WorkerWaitDuration:  1000 * time.Millisecond,
		Func: func(v interface{}) {
			time.Sleep(300 * time.Millisecond)
			fmt.Println(v)
		},
	}

	d := NewDispatcher(opt)
	d.Start()

	for i := 0; i < 1000; i++ {
		d.Add(i)
	}

	d.Wait()
}
