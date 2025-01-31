//Utilizing Boot Dev Courses to learn

package main

//import

import (
	"fmt" //for formatting i/o operations
)

//Chapter 4 Structs Lessons

//Nested Structures
type car struct{
	brand string
	model string
	doors int
	mileage int
	frontWheel wheel //of type wheel: Showing how to utilize nested structs and dot notation for access
	backWheel wheel //of type wheel
}

//Creating Nested Struct
type wheel struct{
	radius int
	material string
}

//Accessing field using dot notation

func main(){
	//Initialize a new car
	UniverseCar := car{}
	//Create a value for wheel
	UniverseCar.backWheel.radius = 10
	UniverseCar.frontWheel.radius = 12
	backWheelRadius := UniverseCar.backWheel.radius
	frontWheelRadius := UniverseCar.frontWheel.radius


	fmt.Printf("The radius of the front wheel is %d, and the back wheel radius is %d\n", backWheelRadius, frontWheelRadius)
}

// go run bootdev.go