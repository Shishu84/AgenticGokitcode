// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
package main

import "fmt"

func main() {

	var name string = "shishu kumar"
	fmt.Println(name)

	var age int = 22
	fmt.Println(age)

	// Go can infer type
	var city = "Gaya"
	fmt.Println(city)

	// short declaration
	country := "India"
	fmt.Println(country)

	// real world user define variable name 
	var projectName string = "inventory-service"
    var port int = 8080

    fmt.Println("Project:", projectName)
    fmt.Println("Port:", port)

    // using short syntax
    environment := "development"

    fmt.Println("Environment:", environment)
}

