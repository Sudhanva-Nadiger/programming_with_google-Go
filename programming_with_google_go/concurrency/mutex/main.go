package main

import (
	"fmt"
	"sync"
)

// mutex makes sure mutual exclusion

var on sync.Once
var m sync.Mutex

func doOnce() {
	fmt.Println("initializing...")
}

func doSomething() {
	on.Do(doOnce)
	fmt.Println("Doing something...")
}

func main() {
	go doSomething()
	go doSomething()
}
