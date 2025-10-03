### Design: Implements a Ticket Lock and a CAS Spinlock. The Ticket Lock uses atomic.AddUint32 as a FAA substitute. The SpinLock uses atomic.CompareAndSwapUint32. The benchmark measure the average wait time under different contention, different amount of threads/goroutines, over 1000 iterations.
# Instructions:
    Open up the file named _
    Run the code
# Analysis:
    The Ticket Lock average times seems to be increasing linearly and take more time than the spin lock.
# Dependencies/library:
    sync/atomic
    runtime
    time
    fmt
    sync
