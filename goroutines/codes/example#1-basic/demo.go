package main

import (
	"fmt"
	"time"
)

func processA() {
	time.Sleep(1 * time.Second)
	fmt.Println("process A is already executed")
}

func processB() {
	time.Sleep(1 * time.Second)
	fmt.Println("process B is already executed")
}

func processC() {
	time.Sleep(1 * time.Second)
	fmt.Println("process C is already executed")
}

func processD() {
	time.Sleep(1 * time.Second)
	fmt.Println("process D is already executed")
}

func processE() {
	time.Sleep(1 * time.Second)
	fmt.Println("process E is already executed")
}

func main() {
	startTime := time.Now()
	processA()
	processB()
	processC()
	processD()
	processE()
	timeConsumption := time.Since(startTime)
	fmt.Printf("process ended after %v", timeConsumption)
}
