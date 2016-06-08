package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for t := range ticker.C {
			fmt.Printf("tick %s\n", t)
		}
		wg.Done()
	}()
	wg.Wait()
}
