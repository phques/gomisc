package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/phques/gomisc/ordone"
)

func main() {
	// ctx.done() will be closed after 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	// start moo and oink generators
	go func() {
		for {
			select {
			case cows <- "moo":
				time.Sleep(time.Millisecond * 200)
			case <-ctx.Done():
				println("moo generator stopped")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case pigs <- "oink":
				time.Sleep(time.Millisecond * 250)
			case <-ctx.Done():
				println("oink generator stopped")
				return
			}
		}
	}()

	var wg = new(sync.WaitGroup)
	wg.Add(2)
	go consumeCows(pigs, ctx.Done(), wg)
	go consumePigs(cows, ctx.Done(), wg)
	wg.Wait()

	time.Sleep(time.Second)
}

func consumePigs(pigs <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ordone.OrDone(done, pigs) {
		fmt.Println(val)
	}

	println("consumePigs stopped")
}

func consumeCows(cows <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ordone.OrDone(done, cows) {
		fmt.Println(val)
	}

	println("consumeCows stopped")
}
