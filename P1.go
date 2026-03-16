// package main // every go file starts with a package declaration . using main tells go 
// // that this file should compile into 

// import "fmt"

// func main() {
//     fmt.Println("Hello, world! I am coding in Go right now.")
// }

package main 
import "fmt"
func main(){
	fmt.Println("Hello, World! Iam coding in golang right now  thanku ")

	// 1.creating variables
	name:= "shishu "
	dayslearning :=1

	// 2. doing some basic math 
	hourInday :=24
	totalHours:=daysLearning * hoursInday
	
	//3. Printing the results 
	fmt.Println("Welcome back,", name)
	fmt.Println("I have been coding in Go for", totalHours,"hours.")
 
	// 4. A simple logic chek (if statement ) 
	if totalHours>=24{
		fmt.Prinln("Great job sticking with it for a full day! ")
	}
	else{
		fmt.Println("you are just getting started!")
		
	}
	


}
 