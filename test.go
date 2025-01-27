// Package declaration
package main

//Import packages
import (
	"fmt"
)

//Warm Up Exercises

//Exercise 1: Define helloWorld func, returns Hello, World!
func helloWorld() string{
	return "Hello, World!"
}

//Exercise 2: Define addSum func, returns sum of two parameters
func addSum(num1, num2 int) int{
	return num1+num2
}

//Exercise 3: Define even or odd func, returns "Even" or "Odd"
func evenOrOdd(number int) string{
	if number % 2 == 0{
		return "Even"
	}
	return "Odd"
}

//Exercise 4: Define Max of Two func, returns larger of two numbers
func largestNumber(num1, num2 int) int{
	if num1 > num2{
		return num1
	} else {
		return num2
	}
}

//Exercise 5: Define area of square func, returns the area ,  length * length
func areaOfSquare(length int) int{
	return length*length
}

//Exercise 6: Define the length of any given string, returns the length of a given string
func lengthString(str string) int{
	//Edge Case check if str is empty
	if str == ""{
		return 0
	}
	//Calculate string length
	return len(str)
}

//Exercise 7: Define a Swap Values func, swaps the values of two intergers passed as pointers
func swapValues(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

//Exercise 8: Define a func to repeat a string n times, n is parameter, returns repeated string
func repeatString(str string, n int) string{
	result :=""
	for i := 0; i<n; i++{
		result += str + " "
	}
	return result
}

//Exercise 9: Define a function that returns the first character of a string, returns string
func firstCharacter(str string) string{
	//Edge Case: If str is length 0, return empty string
	if len(str) == 0{
		return ""
	}
	return string(str[0])
}

//Exercise 10: Define a function to check if a number is positive, returns boolean
func positiveNumber(num int) bool{
	return num >= 0
}



func main(){
	//First
	fmt.Println("The first exercise defines a helloWorld function:",helloWorld())
	fmt.Println("-------------------------------")

	//Second
	fmt.Printf("The second exercise defines a sum function.\nThe sum of %v and %v is: %v\n", 1,3,addSum(1,3))
	fmt.Println("-------------------------------")

	//Third
	fmt.Printf("The third exercise defines an even or odd function.\n The result of passing %v into the function yields: %s\n",3,evenOrOdd(3))
	fmt.Println("-------------------------------")

	//Fourth
	fmt.Println("The fourth exercise defines a largest number function that compares two numbers.\n The largest number is:", largestNumber(40,5))
	fmt.Println("-------------------------------")

	//Fifth
	fmt.Printf("The fifth exercise defines the area of a square calculation.\n The area of the square for %d length is: %d\n",10, areaOfSquare(10))
	fmt.Println("-------------------------------")

	//Sixth
	fmt.Printf("The sixth exercise defines a string length function.\n The total length of the given string is: %d\n", lengthString("Total"))
	fmt.Println("-------------------------------")

	//Seventh
	fmt.Println("The seventh exercise defines a swap values function that swaps the values of two integers passed as pointers")
	x, y := 5, 10
	fmt.Printf("Before swap: x = %d, y = %d\n",x, y)
	swapValues(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n",x,y)
	fmt.Println("-------------------------------")

	//Eighth
	fmt.Println("The eighth exercise repeats any given string N amounts of times")
	fmt.Println(repeatString("LiftingCodes", 5))
	fmt.Printf("The string %s is repeated %d times.\n","LiftingCodes", 5)
	fmt.Println("-------------------------------")
	
	//Ninth Exercise
	fmt.Println("The ninth exercise returns the first character of any given string")
	fmt.Printf("The first character of %s is: %s\n", "Gold", firstCharacter("Gold"))
	fmt.Println("-------------------------------")

	//Tenth Exercise
	fmt.Println("The tenth exercise returns a boolean value, true or false, if the number is positive.")
	fmt.Printf("The number entered into the function was %d, and this is a positive number? %v\n", 10, positiveNumber(10))
	
	}
	



//go run test.go