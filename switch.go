package main

import "fmt"

func main() {

	//calender week's day switch
    day := 6
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
				case 4:
        fmt.Println("Thrustday")
				case 5:
        fmt.Println("Friday")
				case 6:
        fmt.Println("Saturday")
								

    default:
        fmt.Println(" sunday is Holi day") 
    }

				// a simple arithmatic calculator 
				var a, b float64 //  for 15–16 decimal digits of precision.
    var op string

    fmt.Println("Enter first number: ")
    fmt.Scanln(&a)

    fmt.Println("Enter operator (+, -, *, /): ")
    fmt.Scanln(&op)

    fmt.Println("Enter second number: ")
    fmt.Scanln(&b)

    switch op {
    case "+":
        fmt.Println("Result:", a+b)
    case "-":
        fmt.Println("Result:", a-b)
    case "*":
        fmt.Println("Result:", a*b)
    case "/":
        if b != 0 {
            fmt.Println("Result:", a/b)
        } else {
            fmt.Println("Error: Division by zero")
        }
    default:
        fmt.Println("Unknown operator")
    }
}
