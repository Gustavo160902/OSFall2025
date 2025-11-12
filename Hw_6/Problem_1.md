# Research to see Improvement or not nowadays   
## Format of the answer is to what is it referring to and then the value found in the Jeff Dean's talk.   

### L1 cache reference - 0.5 ns  
Found: 0.5ms  
I found the same value in multiple recent sources, meaning there has been no noticeable change since 2009.  
Although the Intel website lists a slightly different value (around 1 ns), that difference comes from specific hardware tests and cannot be used as a general reference.

### Branch mispredict – 5 ns  
Found: 5 ns  
Recent references report about the same value, with small variations between 4 and 5 ns.  
This shows that branch misprediction latency has remained mostly unchanged in modern CPUs.  

### L2 cache reference – 7 ns  
Found: 7 ns  
Modern results are also very similar, and I found several sources reporting values of 7 ns.  
Intel’s page gives a slightly lower number for its Xeon CPUs, but that is again a hardware-specific case, not a general update.  

### Mutex lock/unlock – 25 ns  
Found: 100 ns  
Recent measurements show slightly higher times around 100 ns.  
This difference depends on the operating system and CPU architecture, so the value varies from source to source.  

### Main memory reference – 100 ns  
Found: 100 ns  
Modern memory systems such as DDR4 and DDR5 continue to show about 100 ns latency.  
Intel’s data confirms similar values, meaning there has been little overall change.  

### Read 1 MB from memory – 250 000 ns  
Found: 250 000 ns  
Recent data also list about 250 000 ns for this operation.  
The difference across sources is minimal and mainly due to different memory bandwidths.  

### Round trip within same datacenter – 500 000 ns  
Found: 500 000 ns  
Modern data-center networks show similar latency around 0.5 ms on average.  
Some newer facilities can achieve slightly lower results depending on network size and configuration.  

### Disk seek – 10 000 000 ns  
Found: 10 000 000 ns  
The same value appears in current references for mechanical hard drives.  
Solid-state drives are now much faster, but this value remains valid for HDD performance.  

### Read 1 MB sequentially from disk – 20 000 000 ns  
Found: 30 000 000 ns  
New data list about 30 000 000 ns for HDDs, which is close to Jeff Dean’s value.  
The latency still varies depending on the drive type and speed.  

### Send packet CA ↔ Netherlands – 150 000 000 ns  
Found: 150 000 000 ns  
Recent measurements confirm similar latency values around 150 ms.  
Network improvements reduced variability, but the overall delay across continents remains about the same.  

### Note on Accuracy and Sources  
The values differ slightly among references because each source measures latency under different hardware and software conditions.  
Intel’s *“Memory Performance in a Nutshell”* article provides reliable but CPU-specific numbers, which were used only to confirm trends.  
Overall, cache and memory latency have stayed similar since 2009, while storage and network performance have improved most "significantly".  


### References  
- Jeff Dean, *“Numbers Everyone Should Know,”* LADIS 2009  
- Intel Corporation - https://www.intel.com/content/www/us/en/developer/articles/technical/memory-performance-in-a-nutshell.html   
- Forum https://stackoverflow.com/questions/59505783/latency-of-accessing-main-memory-is-almost-the-same-order-of-sending-a-packet
