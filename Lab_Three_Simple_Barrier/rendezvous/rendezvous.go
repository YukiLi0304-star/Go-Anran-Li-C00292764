package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, threadCount int, mu *sync.Mutex, arrived *int, barrier chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	mu.Lock()
	*arrived++
	CurrentArrived := *arrived
	mu.Unlock()

	if CurrentArrived == threadCount {
		for i := 0; i < threadCount-1; i++ {
			barrier <- true
		}
	} else {
		<-barrier
	}

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	arrived := 0
	barrier := make(chan bool)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, threadCount, &mu, &arrived, barrier)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
