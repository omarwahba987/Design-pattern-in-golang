package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

// Task that will be done over time
func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1.")
	wg.Done()
}
func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2.")
	wg.Done()
}
func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3.")
	wg.Done()
}

// Task done in parallel
func listenForever() {
	for {
		printTime("Listening...")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	go listenForever()

	// Let's start writing the mails
	go writeMail1(&waitGroup)
	go writeMail2(&waitGroup)
	go writeMail3(&waitGroup)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	waitGroup.Wait()
}

//https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca
//https://hub.packtpub.com/concurrency-and-parallelism-in-golang-tutorial/
