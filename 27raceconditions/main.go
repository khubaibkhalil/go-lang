package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("One R")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()

	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Two R")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Three R")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
