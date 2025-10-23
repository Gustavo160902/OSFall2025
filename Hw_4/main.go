// Implementation of Two-Lock Queue and Lock-Free Queue HW#4 Problem 2 in Go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type TLNode struct {
	data int
	next *TLNode
}

// two-lock queue with separate locks for head and tail
type TwoLockQueue struct {
	head     *TLNode
	tail     *TLNode
	lockHead sync.Mutex
	lockTail sync.Mutex
}

func NewTwoLockQueue() *TwoLockQueue {
	dummy := &TLNode{}
	return &TwoLockQueue{head: dummy, tail: dummy}
}

func (q *TwoLockQueue) Enqueue(val int) {
	node := &TLNode{data: val}
	q.lockTail.Lock()
	q.tail.next = node
	q.tail = node
	q.lockTail.Unlock()
}

func (q *TwoLockQueue) Dequeue() (int, bool) {
	q.lockHead.Lock()
	defer q.lockHead.Unlock()
	if q.head.next == nil {
		return 0, false
	}
	value := q.head.next.data
	q.head = q.head.next
	return value, true
}

// node for the lock-free queue
type LFNode struct {
	data int
	next atomic.Pointer[LFNode]
}

type LockFreeQueue struct {
	head atomic.Pointer[LFNode]
	tail atomic.Pointer[LFNode]
}

func NewLockFreeQueue() *LockFreeQueue {
	dummy := &LFNode{}
	q := &LockFreeQueue{}
	q.head.Store(dummy)
	q.tail.Store(dummy)
	return q
}

func (q *LockFreeQueue) Enqueue(val int) {
	node := &LFNode{data: val}
	for {
		tail := q.tail.Load()
		next := tail.next.Load()
		if next == nil {
			if tail.next.CompareAndSwap(nil, node) {
				q.tail.CompareAndSwap(tail, node)
				return
			}
		} else {
			q.tail.CompareAndSwap(tail, next)
		}
	}
}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := q.head.Load()
		next := head.next.Load()
		if next == nil {
			return 0, false
		}
		value := next.data
		if q.head.CompareAndSwap(head, next) {
			return value, true
		}
	}
}

// runs a test for each queue type
func runTest(name string, qType string, total int) {
	start := time.Now()

	if qType == "two-lock" {
		q := NewTwoLockQueue()
		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < total; j++ {
					q.Enqueue(j)
					q.Dequeue()
				}
			}()
		}
		wg.Wait()
	} else {
		q := NewLockFreeQueue()
		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < total; j++ {
					q.Enqueue(j)
					q.Dequeue()
				}
			}()
		}
		wg.Wait()
	}

	fmt.Printf("%s finished in %v\n", name, time.Since(start))
}

func main() {
	runTest("Two-Lock Queue", "two-lock", 100000)
	runTest("Lock-Free Queue", "lock-free", 100000)
}
