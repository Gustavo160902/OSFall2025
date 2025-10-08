# Go Design Principle: “Don’t communicate by sharing memory; share memory by communicating.”   

This idea means that instead of letting multiple goroutines or processes work on the same data using locks,
Go programs send data through channels or pipes. Each goroutine owns its data and passes it along when needed. This helps prevent race conditions and makes concurrency easier to understand and manage.  

In our Operating Systems homework #0 (problem #1), we used the same idea in a producer–consumer program. 
The goal was to make two processes communicate through pipes instead of sharing variables directly. My program used Go’s os and os/exec libraries, where the producer sent numbers and the consumer sent back acknowledgments. That follows Go’s main idea of “sharing memory by communicating.” 
This shows that the data moves through a communication channel, not shared memory. Each process handles its own data and only communicates when needed, which is exactly what the Go principle encourages.

### Pros:     

Helps keep the code simple and clear because communication paths are obvious.   
Works well with Go’s lightweight goroutines and automatic scheduling.    
Reduces the need for complicated locking and synchronization.  

### Cons:

Some programs still need shared memory for performance reasons.     
You have to be careful with blocking channels to avoid deadlocks.     
Go doesn’t force you to follow this principle—it’s more of a design style that developers should apply.    

### Sources  

Operating Systems – HW0 (Producer–Consumer Program)  
Why reliable: this class project directly applies Go’s concurrency idea using pipes and inter-process communication.    

Go Blog – “Share Memory by Communicating.”  
Why reliable: official article from the Go team/company.

Effective Go – “Concurrency.”  
Why reliable: official Go documentation

### Links  

https://go.dev/blog/codelab-share   
https://www.digitalocean.com/community/tutorials/understanding-concurrency-in-go
       
AI Disclosure  
AI assistance was used only to help locate reliable Go documentation and to format the citations and sources clearly. Expalnations were based on our own understanding.
