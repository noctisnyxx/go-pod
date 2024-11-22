package main

import (
	"fmt"
	"time"
)

func processA() {
	time.Sleep(time.Second)
	fmt.Println("process A finished")
}

func processB() {
	time.Sleep(time.Second)
	fmt.Println("process B finished")
}

func processC() {
	time.Sleep(time.Second)
	fmt.Println("process C finished")
}

func main() {
	startTime := time.Now()
	go processA()
	go processB()
	go processC()
	time.Sleep(2 * time.Second)
	timeConsumption := time.Since(startTime)
	fmt.Printf("process ended after %v", timeConsumption)
}
