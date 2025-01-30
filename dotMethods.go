package main

import "fmt"

//Hello World Function: Returns "Hello, World!"
func HelloWorld() string{
    return "Hello, World!"
}

//Sum Two Numbers Function: Returns the Sum of Two Numbers
func SumTwo(num1, num2 int) int{
    return num1 + num2
}

//Greet by Name Function: Greets by name
func greetName(name string){
    fmt.Printf("Hello, %s!",name)
}

//Area of Rectangle: Returns area of a rectangle
func rectArea(length, width int) int{
    return length * width
}

//Is Even Function: Determines whether number given is even or odd
func isEven(num int) string{
    if num % 2 == 0{
        return "Even"
    } else{
        return "Odd"
    }

}

//Max of Two Numbers Function: Returns the maximum number between any two given numbers 
func MaxTwo(num1, num2 int) int{
    return max(num1, num2)
}

//Square of Number Function: Returns a number times itself, resulting in number squared
func squaredNumber(num1 int) int{
    return num1*num1
}

//Convert Celcius to Farenheit Function: Converts A Given Number into Farenheit / e.g. 0 degree Celcius = 32 degrees Farenheit
func celciusToFarenheit(num int) int{
    return (num * 9/5) + 32
}

//String Length Function: Returns the string length of any given string
func stringLength(str string) int{
    return len(str)
}

//Swap Two Variables Function: Returns two variables in reverse order
func reverseVariables(varOne, varTwo string) string{
    varOne, varTwo = varTwo, varOne
    return varOne+varTwo
}


func main() {
    //Testing Functions
    //Test Hello World
    fmt.Println(HelloWorld())
    fmt.Println()

    //Test Sum Two Function
    fmt.Println(SumTwo(2,2))
    fmt.Println()

    //Test Square Number Function
    fmt.Printf("The Square of the Given Number is %d\n",squaredNumber(8))

    //Test Greet Name Function
    greetName("Nick")
    fmt.Println()

    //Test Max Two Function
    fmt.Printf("The max of the two numbers is: %d\n",MaxTwo(200,400))


    //Test is Even Function
    fmt.Println()
    fmt.Println("The number given is even: ", isEven(2))
    fmt.Println()

    //Test Area of Rectangle Function
    fmt.Println("The area of the rectangle is:", rectArea(20,240))
    fmt.Println()

    //Test Celcius To Farenheit Function
    fmt.Println(celciusToFarenheit(200))
    fmt.Println()

    //Test length of the Given Word Function
    fmt.Printf("The length of the given word is: %d",stringLength("Bob"))
    fmt.Println()

    //Test Reverse Variables Function
    fmt.Println(reverseVariables("Nick","Alvarez"))
    fmt.Println()


    
}

//go run dotMethods.go