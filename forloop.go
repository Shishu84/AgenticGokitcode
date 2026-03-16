package main

import "fmt"

func main() {

   // for loop syntax  in Golang 
    for i := 1; i <= 5; i++ {
        fmt.Println("Processing task", i)
    }
    // simulate checking logs in real world for authenticating 
    for log := 1; log <= 4; log++ {
        fmt.Println("Checking log entry", log)
    }
}