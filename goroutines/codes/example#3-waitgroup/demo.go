package main

import (
	"fmt"
	"sync"
	"time"
)

func processA(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("process A is already executed")
}

func processB(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("process B is already executed")
}

func processC(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("process C is already executed")
}

func processD(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("process D is already executed")
}

func processE(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("process E is already executed")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	startTime := time.Now()
	go processA(&wg)
	go processB(&wg)
	go processC(&wg)
	go processD(&wg)
	go processE(&wg)
	wg.Wait()
	timeConsumption := time.Since(startTime)
	fmt.Printf("process ended after %v", timeConsumption)
}
