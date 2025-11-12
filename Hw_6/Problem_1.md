# Research to see Improvement or not nowadays   
## Format of the answer is to what is it referring to and then the value found in the Jeff Dean's talk.   


### L1 cache reference - 0.5 ns
I found the same value in multiple recent sources, meaning there has been no noticeable change since 2009.  
Although the Intel website lists a slightly different value (around 1 ns), that difference comes from specific hardware tests and cannot be used as a general reference.
### Branch mispredict
### L2 cache reference - 7ns
Modern results are also very similar, and I found several sources reporting values of 7ns.
Intel’s page gives a slightly lower number for its Xeon CPUs, but that is again a hardware-specific case, not a general update.
### Mutex lock/unlock
### Main memory reference
### Read 1 MB sequentially from memory
### Round trip within same datacenter
### Disk seek
### Read 1 MB sequentially from disk
### Send packet CA->Netherlands->CA
