package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func main() {
	websites := []string{"https://www.facebook.com",
		"https://www.linkedin.com",
		"https://www.instagram.com",
		"https://www.hahahafhalfhaskf.com"}

	var start = time.Now()

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	fmt.Println()

	wg.Wait()
	fmt.Println(signals)
	fmt.Println(time.Since(start))
}

func getStatusCode(endpoint string) {

	defer wg.Done()
	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf(" %v", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)

		mut.Unlock()
		fmt.Printf(" %v %v\n", resp.Status, endpoint)

	}

}
