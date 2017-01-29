package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	args []string
)

func init() {
	flag.Parse()

	args = flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Missing filename: specify one argument")
		os.Exit(1)
	}
}

func tail(c context.Context, name string) error {
	fi, err := os.Open(name)
	if err != nil {
		return err
	}
	defer fi.Close()

	const size = 10
	ch := make(chan string, size)

	// prepare
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		if len(ch) >= size {
			<-ch
		}
		ch <- scanner.Text()
	}

	// output
	go func() {
		for {
			t, ok := <-ch
			if !ok {
				return
			}
			fmt.Fprintln(os.Stdout, t)
		}
	}()

	var mu sync.Mutex
	go func() {
		for {
			if ch == nil {
				return
			}
			scanner := bufio.NewScanner(fi)
			for scanner.Scan() {
				mu.Lock()
				ch <- scanner.Text()
				mu.Unlock()
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	<-c.Done()
	mu.Lock()
	close(ch)
	ch = nil
	mu.Unlock()
	return nil
}

func main() {
	var wg sync.WaitGroup

	c, cancel := context.WithCancel(context.Background())
	for _, n := range args {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tail(c, n)
		}()
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		for {
			s := <-sig
			switch s {
			case syscall.SIGINT:
				cancel()
			}
		}
	}()

	wg.Wait()
	time.Sleep(time.Millisecond * 300)
}
