Goroutines are lightweight units of work managed by the Go runtime, while 
threads are heavier and managed by the operating system. 
This matters because OS threads have kernel bookkeeping and larger default stacks, 
so creating many threads uses more memory and CPU. Goroutines communicate with channels
in built into the language, while threads usually use locks and condition variables from
libraries. The main trade-off is that goroutines help maximize CPU use with less overhead
but give you less low-level control; threads give you full OS control, but they consume 
more system resources when you create a lot of them. 

Pros: 

Lightweight — cheaper to run many of them. 

Built-in communication with channels/select — simpler, cleaner code. 

Go runtime schedules across cores — less manual thread management. 

Cons: 

Less direct, low-level control than threads. 

You must avoid goroutine leaks and watch for blocking calls that can stall work. 

Performance still depends on runtime settings and your design; goroutines aren’t a magic speed-up. 

 

Citations / Sources 

Medium — Geison F. G. “Unleashing Concurrent Power…”  
Why reliable : practitioner write-up; good as an example. 

Medium — Damini Bansal. “OS Threads vs Goroutines…”  
Why reliable: clear explanation of models; useful secondary user experience view. 

Go Concurrency Patterns (official Go talk slides, go.dev). 
 Why reliable: official material from the Go team (Google). It states goroutines are “very cheap,” practical to have thousands, and that they’re multiplexed onto threads 

 

Links 

https://medium.com/@geisonfgfg/unleashing-concurrent-power-exploring-the-differences-pros-and-cons-of-go-routines-and-channels-c04ffa8a917a 

 

https://daminibansal.medium.com/os-threads-vs-goroutines-understanding-the-concurrency-model-in-go-bad187372c89 

 

https://go.dev/doc/faq#goroutines 
