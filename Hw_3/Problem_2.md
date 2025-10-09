# Design:   Implements two Concurrent Linked Lists with either a lock for the whole list or a lock per node, hand-over-hand.  The benchmark measures the average wait time under different workloads, ratios of inserts to lookups. Insert-Heavy, 90% insert and 10% lookup, Lookup-Heavy, 10% insert and 90% lookup, Balanced, 50% insert and 50% lookup. These different workloads simulate the each extremes, lookup or insert heavy, and when it is in the middle, balanced.   
## Instructions:
Open up the file named _
Run the code
## Analysis:
The simple concurrent list is better than the hand-over-hand method, unless there is more than 99% insert, where there is a chance that hand-over-hand has better times, as hand-over-hand has many locks that seem to add over time.
### Dependencies/library:
fmt
sync
time
