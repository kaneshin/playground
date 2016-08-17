package main

import (
	"log"
	"sync"
	"time"

	"github.com/kaneshin/playground/go/concurrency-pattern/payload"
)

type (
	Dispatcher struct {
		pool    chan *worker
		queue   chan interface{}
		wg      sync.WaitGroup
		workers []*worker
		quit    chan struct{}
	}
	worker struct {
		dispatcher *Dispatcher
		wait       time.Duration
		signal     chan struct{}
		data       chan interface{}
		iterator   func(interface{})
		quit       chan struct{}
	}
	DispatcherOptions struct {
		NumberOfWorkers     int
		QueueCapacity       int
		WorkerQueueCapacity int
		WorkerWaitDuration  time.Duration
		Iterator            func(interface{})
	}
)

func (d *Dispatcher) Add(v interface{}) {
	d.wg.Add(1)
	d.queue <- v
}

func (d *Dispatcher) SetIterator(f func(interface{})) {
	for _, w := range d.workers {
		w.iterator = f
	}
}

func (d *Dispatcher) Start() {
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

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

func (d *Dispatcher) Stop(immediately bool) {
	if !immediately {
		d.Wait()
	}

	d.quit <- struct{}{}
	for _, w := range d.workers {
		w.quit <- struct{}{}
	}
}

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

func (w *worker) start() {

	if cap(w.data) > 1 {
		go func() {
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
					if w.iterator != nil {
						w.iterator(list)
					}

					w.dispatcher.wg.Add(-count)
				}
			}())

			if !t.Stop() {
				<-t.C
			}
			for {
				w.dispatcher.pool <- w

				select {
				case _, ok := <-w.signal:
					if !ok {
						return
					}
					if len(w.data) == cap(w.data) {
						t.Reset(0)
					} else {
						t.Reset(w.wait)
					}

				case <-w.quit:
					return

				}
			}
		}()
	} else {
		go func() {
			for {
				w.dispatcher.pool <- w

				select {
				case v, ok := <-w.data:
					if !ok {
						return
					}

					if w.iterator != nil {
						w.iterator(v)
					}
					w.dispatcher.wg.Done()

				case <-w.quit:
					return

				}
			}
		}()
	}
}

func NewDispatcher(opt DispatcherOptions) *Dispatcher {
	d := &Dispatcher{
		pool:  make(chan *worker, opt.NumberOfWorkers),
		queue: make(chan interface{}, opt.QueueCapacity),
		quit:  make(chan struct{}),
	}
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
		d.workers = append(d.workers, &w)
		w.start()
	}
	if opt.Iterator != nil {
		d.SetIterator(opt.Iterator)
	}
	return d
}

var dispatcher *Dispatcher

func init() {
	dispatcher = NewDispatcher(DispatcherOptions{
		NumberOfWorkers:    numMaxWorkers(),
		QueueCapacity:      numMaxQueues(),
		WorkerWaitDuration: 1000 * time.Millisecond,
		Iterator: func(v interface{}) {
			instance, ok := v.(payload.Instance)
			if !ok {
				return
			}
			body, err := instance.Get()
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Get %s  (%d bytes)\n", instance.URL, len(body))
			}
		},
	})

	dispatcher.Start()
}

func do(instance payload.Instance) {
	dispatcher.Add(instance)
}
