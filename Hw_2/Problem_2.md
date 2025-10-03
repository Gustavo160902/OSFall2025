## Lock Analysis II
# After changing the unlock function from lock->flag = 0; to lock->flag = lock->flag - 1;.

The new algorithm will not be correct, because the flag can go negative if the lock is not used properly. 
For example, if a thread calls unlock() without holding the lock, or calls it more than once, the 
value of flag can become -1. Once this happens, no other thread will be able to acquire the lock, 
so they will spin forever.

T0 holds the lock  
T0 unlock() // first call: flag = 1 - 1 = 0 (release works)  
T0 unlock() // second call: flag = 0 - 1 = -1  
Other threads spin forever // flag is never 0 again  


