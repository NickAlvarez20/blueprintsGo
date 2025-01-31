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
	UniverseCar := car{brand: "Tesla", model: "Roaster", doors: 4, mileage: 100000, frontWheel: wheel{radius: 100, material: "Plywood"}, backWheel: wheel{radius: 200, material: "Glue"}}
	//Create a value for wheel
	backWheelRadius := UniverseCar.backWheel.radius
	frontWheelRadius := UniverseCar.frontWheel.radius


	fmt.Printf("The radius of the front wheel is %d, and the back wheel radius is %d\n", backWheelRadius, frontWheelRadius)
	fmt.Printf("The material made in front wheel of the %s is %s. The material used in making the back wheel of the %s is %s.\n", UniverseCar.brand, UniverseCar.frontWheel.material, UniverseCar.brand, UniverseCar.backWheel.material)
}

// go run bootdev.go