package main

import (
	"fmt"
	"sync"
)

/*
A race condition occurs when two or more goroutines access a shared variable
concurrently, and at least one of them modifies it, leading to unpredictable results.

in this example both go routine call x++
If both goroutines perform these steps at the same time, they can overwrite each other's updates,
*/

func goroutines() {
	x := 1

	channels := make(chan bool, 2)

	go func() {
		x++
		fmt.Printf("from first routine %v\n", x)
		x++
		channels <- true
	}()

	go func() {
		fmt.Printf("1. from second routine : %v\n", x)
		x++
		fmt.Printf("2. from second routine : %v\n", x)
		channels <- true
	}()

	<-channels
	<-channels
}

func waitGroups() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("from first routine")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("from second routine")
	}()

	wg.Wait()
}

func goRputinesCommunication() {
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()
	ch := make(chan int)

	go func() {
		defer wg.Done()
		ch <- 42
	}()

	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
}

func mul(a, b int, c chan int) {
	c <- a * b
}

func goRoutinesCommunication1() {
	c := make(chan int)
	go mul(2, 3, c)
	go mul(4, 5, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}

func main() {
	goRoutinesCommunication1()
}
