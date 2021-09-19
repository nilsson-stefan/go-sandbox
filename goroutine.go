package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func sqrtWorker(chIn chan int, chOut chan int) {
	fmt.Printf("sqrtWorker started\n")
	for i := range chIn {
		sqrt := i * i
		time.Sleep(time.Duration(sqrt) * time.Second)
		chOut <- sqrt
	}
	fmt.Printf("sqrtWorker finished\n")
	waitGroup.Done()
}

func testConcurrentChannels() {
	fmt.Printf("\n\ntestConcurrentChannels\n")

	chIn := make(chan int)
	chOut := make(chan int)

	fmt.Printf("Channels created\n")

	for i := 0; i < 2; i++ {
		waitGroup.Add(1)
		go sqrtWorker(chIn, chOut)
	}

	fmt.Printf("Two workers created\n")

	go func() {
		chIn <- 1
		chIn <- 2
		chIn <- 3
		close(chIn)
	}()

	fmt.Printf("The numbers sent to in channel\n")

	go func() {
		waitGroup.Wait()
		close(chOut)
	}()
	fmt.Printf("Waiting for channel out\n")

	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}
