## Problem 1 – ABA Problem

The ABA problem happens when a thread is working on a node A but gets paused during the operation.  
While it is paused, another thread changes the node to node B and then changes it back to node A.  
When the first thread continues, it believes that nothing has changed, but in reality, the data was modified while it was paused.

In the enqueue example discussed in class, this issue can happen when one thread is trying to enqueue a node but is paused before completing the Compare-And-Swap step.  
During that time, another thread enqueues a new node B and dequeues the original node A.  
Since node A was removed, its memory can be reused for another node, possibly even one from another queue.  
When the first thread wakes up, it still sees the same memory address and continues as if nothing changed, skipping node B and corrupting the queue order.

### Sources

CON09-C: Avoid the ABA problem when using lock-free algorithms – SEI CERT C Coding Standard  
https://wiki.sei.cmu.edu/confluence/display/c/CON09-C.+Avoid+the+ABA+problem+when+using+lock-free+algorithms  
This source is credible since it is published by the Software Engineering Institute at Carnegie Mellon University and includes real examples and solutions for ABA issues in lock-free queues.

The ABA Problem – Baeldung on Computer Science  
https://www.baeldung.com/cs/aba-concurrency  
This source is credible because Baeldung is a reviewed technical website written by software engineers, explaining in detail how CAS operations can lead to the ABA problem.

Michael and Scott Queue Paper – University of Rochester  
https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf  
This source is credible because it comes from the Department of Computer Science at the University of Rochester and presents the original queue code used in class, describing how CAS must handle the ABA problem.
