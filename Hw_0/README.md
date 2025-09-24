# OSFall2025_hw0 // EECE.4811
Operating Systems Fall 2025 Class – Homework #0

---

## Names of each code file:
Problem 1 is hw0_OS.go  
Problem 2 is main.go

---

## How to Run Both Problems
Make sure you have Go installed (version 1.25.x or newer).  
It does not matter if you are using Linux, macOS, Windows, or WSL.  

Check the dependencies/libraries for each program.  

Open a terminal inside the `hw0/` folder and run:
go run hw0.go

**Note:** Both programs were only tested in VS Code (Windows and WSL).  

---

## Designs
Problem #1
- The program has two process Producer and Consumer.
- The Producer makes numbers 1 to 5 and sends them to the Consumer.
- The Consumer prints the numbers and then sends back an "ack" to the Producer.
- Pipes allow comunication between process to each other.
- The Producer waits for the ack before sending the next number.
- This repeats until both reach 5 cycles.

Problem #2
-There is already a test in the code to check if the program works correctly.   
-The output of the test should be: 3 2 6 5 7 4 9 8 10 1, which should show up in the terminal.    
-To modify the test, alter the series of push and pop functions in the main function.   
-Works by creating a 100-sized array and using the push and pop functions to modify the elements of the array to work as a stack.

---

## Dependencies / Libraries Used
Problem #1
- The 3 libraries used were os, os/exec and fmt.
- They made it easier to use the concepts learned in class about how the output from one process can be sent to the second one and then back.

Problem #2
- The only library used was fmt.
