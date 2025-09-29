###Lock Analysis   
After taking out the code of guard from the algorithm. And assuming thre iis not issue wakeup/waiting race.  
The new algorithm will not be correct, because concurrent queue operations can be interleaved, leading to multiple calls and breaking mutual exclusion. 
This assumes only what is stated in the problem, without taking into account any possibility that the queue functions are thread-safe.   
T0 holds the lock // line 8    
T1 queque_add // line 20 goes to else as it doesn't satisfy the IF stament   
T2 queque_add // line 20 goes to else as it doesn't satisfy the IF stament
T1 and T2 are call interleave // this leads to the queque to be corrupted which drives to T2 to be let in wait/ sleep forever.
