// Created by Deryck Tran and Gustavo Lujan - 10/2/25
// This program implements TicketLocks and SpinLocks to benchmark their times
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// TicketLock
type TicketLock struct {
	ticket uint32
	turn   uint32
}

// Lock the ticket lock using FAA substitute
func (l *TicketLock) Lock() {
	ticket := atomic.AddUint32(&l.ticket, 1) - 1
	for atomic.LoadUint32(&l.turn) != ticket {
		runtime.Gosched()
	}
}

// Unlock the Ticket Lock
func (l *TicketLock) Unlock() {
	atomic.AddUint32(&l.turn, 1)
}

// SpinLock
type SpinLock struct {
	flag uint32
}

// Lock the Spinlock using compare and swap
func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32(&l.flag, 0, 1) {
		runtime.Gosched()
	}
}

// Unlock the Spinlock
func (l *SpinLock) Unlock() {
	atomic.StoreUint32(&l.flag, 0)
}

// Benchmark the different locks with different goroutines
func benchmarkLock(lock interface {
	Lock()
	Unlock()
}, goroutines int, iterations int) {
	//WaitGroup waits for all goroutines to finish
	var wg sync.WaitGroup
	waitTimes := make([]time.Duration, goroutines)

	//Start the goroutines
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			totalWait := time.Duration(0)
			for j := 0; j < iterations; j++ {
				start := time.Now()
				lock.Lock()
				totalWait += time.Since(start)
				// Critical section
				lock.Unlock()
			}
			waitTimes[idx] = totalWait
		}(i)
	}
	wg.Wait()
	// Print average wait time
	sum := time.Duration(0)
	for _, wt := range waitTimes {
		sum += wt
	}
	fmt.Printf("Average wait time: %v\n", sum/time.Duration(goroutines*iterations))
}

func main() {
	//Different goroutines or threads to test
	goroutines := []int{2, 4, 8, 16, 32, 64, 128}
	//Amount of iterations each goroutine will do
	iterations := 1000

	//Benchmark the TicketLock
	fmt.Println("Benchmarking TicketLock:")
	for _, g := range goroutines {
		fmt.Printf("Goroutines: %d\n", g)
		benchmarkLock(&TicketLock{}, g, iterations)
	}

	//Benchmark the SpinLock
	fmt.Println("\nBenchmarking SpinLock:")
	for _, g := range goroutines {
		fmt.Printf("Goroutines: %d\n", g)
		benchmarkLock(&SpinLock{}, g, iterations)
	}
}
