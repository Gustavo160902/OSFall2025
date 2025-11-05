## Problem 1 – Priority Inversion and Priority Inheritance   
### Question 1    
Priority Inversion -    
The 3 threads and their priority: 
T1 = high priority = 1   
T2 = low priority = 2   
T3 = medium priority = 3    

T2 starts first and gets a lock. Then T1 arrives and needs the same lock, so T1 waits. After that T3 arrives and since it has   higher priority than T2, T3 runs and stops T2 from finishing. Because of this, T1 (the highest priority) has to wait behind   both T2 and T3. That is priority inversion since a lower-priority task ends up delaying a high-priority one.   

Timeline of events:  
Time 0: T2 starts first and gets the lock    
Time 1: T1 arrives and waits because T2 still has the lock    
Time 2: T3 arrives but can't run once priority rules apply   
Time 0–5: T2 keeps running and finishes its critical section   
Time 5–9: T1 runs after T2 releases the lock   
Time 9–12: T3 runs last after T1 finishes

### Question 2     
1 - T2 runs 5 time units in both PIP and PPP   
2 - Both finish T2 at time 5    
3 - Both finish at time 12   
4 - T1 sees inversion in PIP, but not in PPP    
5 - In both, T3 waits and runs last, but PPP prevents inversion and PIP fixes it by boosting T2     




### Sources  
Difference Between Priority Inversion and Priority Inheritance   
https://www.geeksforgeeks.org/operating-systems/difference-between-priority-inversion-and-priority-inheritance/   
Textbook - Operating Systems: Three Easy Pieces   
https://pages.cs.wisc.edu/~remzi/OSTEP/   
