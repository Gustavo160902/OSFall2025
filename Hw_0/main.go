//Created By Deryck Tran
//This program implements a stack using an array

package main

import "fmt"

//push an interger after the last element in the array
func push(arr *[100]int,input int){
    for i := 0; i < len(arr); i++{
        if arr[i] == 0{
            arr[i] = input
            return
        }
    }
    //prints if the stack is full
    fmt.Println("Stack is full")
}

//pops the newest element inputted into the array
func pop(arr *[100]int){
    for i := len(arr) - 1; i >= 0; i--{
        if arr[i] != 0{
            fmt.Println(arr[i])
            arr[i] = 0
            return
        }
    }
    //prints if the stack is empty
    fmt.Println("Stack is empty")
}

func main() {
    //Create an array that has 100 indexes with each element being 0
    var stack [100]int

    //A series of push and pop statement to test the functions
    //Successful test: 3 2 6 5 7 4 9 8 10 1
    push(&stack, 1)
    push(&stack, 2)
    push(&stack, 3)
    pop(&stack)
    pop(&stack)
    push(&stack, 4)
    push(&stack, 5)
    push(&stack, 6)
    pop(&stack)
    pop(&stack)
    push(&stack, 7)
    pop(&stack)
    pop(&stack)
    push(&stack, 8)
    push(&stack, 9)
    pop(&stack)
    pop(&stack)
    push(&stack, 10)
    pop(&stack)
    pop(&stack)
}
