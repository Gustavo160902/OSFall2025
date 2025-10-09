//Create by Deryck Tran and Gustavo Lujan on 10/8/25
//This program implements and benchmarks a concurrent linked list with either a basic mutex or hand-over-hand locking

package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrent Linked List: Rewritten
type Node struct {
	key  int
	next *Node
}

type List struct {
	head *Node
	lock sync.Mutex
}

// Concurrent Linked List: Rewritten Insert
func (l *List) Insert(key int) {
	newNode := &Node{key: key}
	l.lock.Lock()
	newNode.next = l.head
	l.head = newNode
	l.lock.Unlock()
}

// Concurrent Linked List: Rewritten Lookup
func (l *List) Lookup(key int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

// Hand-over-hand locking linked list
type HHNode struct {
	key  int
	next *HHNode
	lock sync.Mutex
}

type HHList struct {
	head *HHNode
}

// Hand-over-hand locking Insert
func (l *HHList) Insert(key int) {
	newNode := &HHNode{key: key}
	newNode.lock.Lock()
	newNode.next = l.head
	l.head = newNode
	newNode.lock.Unlock()
}

// Hand-over-hand locking Lookup
func (l *HHList) Lookup(key int) bool {
	curr := l.head
	if curr == nil {
		return false
	}
	curr.lock.Lock()
	for curr != nil {
		next := curr.next
		if next != nil {
			next.lock.Lock()
		}
		// Hand over hand locking
		if curr.key == key {
			curr.lock.Unlock()
			if next != nil {
				next.lock.Unlock()
			}
			return true
		}
		curr.lock.Unlock()
		curr = next
	}
	return false
}

// Benchmarking function
func benchmark(listType string, insertFunc func(int), lookupFunc func(int) bool, goroutines int, insertRatio int) time.Duration {
	// Start the timer
	start := time.Now()
	// WaitGroup waits for all goroutines to finish
	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		// Add a goroutine to the WaitGroup
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Decide to insert or lookup based on the ratio
			if i%100 < insertRatio {
				insertFunc(i)
			} else {
				lookupFunc(i)
			}
		}(i)
	}
	wg.Wait()
	// Stop the timer and calculates the average time per operation
	temp := time.Since(start) / time.Duration(goroutines)
	return temp
}

func main() {
	goroutines := 10000 // Increase number of goroutines
	preinsert := 10000  // Increase the initial number of nodes

	// Concurrent Linked List: Rewritten Initial list
	basic := &List{}
	// Pre-insert nodes
	for i := 0; i < preinsert; i++ {
		basic.Insert(i)
	}
	basicInsert := func(i int) { basic.Insert(i) }
	basicLookup := func(i int) bool { return basic.Lookup(i) }

	// Hand-over-Hand Initial list
	hh := &HHList{}
	// Pre-insert nodes
	for i := 0; i < preinsert; i++ {
		hh.Insert(i)
	}
	hhInsert := func(i int) { hh.Insert(i) }
	hhLookup := func(i int) bool { return hh.Lookup(i) }

	fmt.Println("Benchmarking with 10,000 operations and 10,000 preloaded nodes...")

	// Different workloads

	// Insert-Heavy (90% insert)
	fmt.Println("\nInsert-Heavy (90% insert):")
	fmt.Printf("Basic List: %v\n", benchmark("basic", basicInsert, basicLookup, goroutines, 90))
	fmt.Printf("HH List:    %v\n", benchmark("hh", hhInsert, hhLookup, goroutines, 90))
	// Lookup-Heavy (90% lookup)
	fmt.Println("\nLookup-Heavy (90% lookup):")
	fmt.Printf("Basic List: %v\n", benchmark("basic", basicInsert, basicLookup, goroutines, 0))
	fmt.Printf("HH List:    %v\n", benchmark("hh", hhInsert, hhLookup, goroutines, 0))
	// Balanced (50% insert, 50% lookup)
	fmt.Println("\nBalanced (50/50):")
	fmt.Printf("Basic List: %v\n", benchmark("basic", basicInsert, basicLookup, goroutines, 50))
	fmt.Printf("HH List:    %v\n", benchmark("hh", hhInsert, hhLookup, goroutines, 50))
}
