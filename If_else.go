package main

import "fmt"

func main() {
     // usese variable without declaration of data type 
    temperature := 32

    if temperature > 30 {
        fmt.Println("It's hot outside")
    } else {
        fmt.Println("Weather is normal")
    }


			
  // usese variable without declaration of data type  it is ok in go lang
    number := 10

    if number%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // if with short statement
    if x := 20; x > 10 {
        fmt.Println("x is greater than 10")
    }
}
