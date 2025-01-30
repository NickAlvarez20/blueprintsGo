//Project Structs: Understanding how to utilize structs, methods, and pointers
package main

//import packages
import (
	"fmt" //for formatting I/O operations
)

//Define structs

//Coffee Struct
type Coffee struct{
	color string
	flavor string
	ounces int
}

//Create a collection of Coffee
type CoffeeCollection struct{
	coffees []Coffee
}

//Append new coffee to collection
func (c *CoffeeCollection) newCoffee(color, flavor string, ounces int) {
	newCoffee := Coffee{
		color: color,
		flavor: flavor,
		ounces: ounces,
	}
	c.coffees = append(c.coffees, newCoffee)
}


//func main for creating program
func main(){
	//Initialize new coffee collection
	lightCollection := []CoffeeCollection{}
	//Append new coffe into light collection
	
}

