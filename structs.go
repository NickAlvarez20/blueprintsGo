// Project Structs: Understanding how to utilize structs, methods, and pointers
package main

//import packages
import (
	"fmt" //for formatting I/O operations
)

//Define coffee struct for creating cups of coffee

type cupOfCOffee struct{
	color string
	taste string
	ounces int
}

//Define CoffeeCollection that organizes a group of different cups of coffeeL: Light Collection
type CoffeeCollection struct{
	coffees []cupOfCOffee
}

//Define Methods for light coffee collection: Add a cup to the collection
func (l *CoffeeCollection) addLightCoffee(color, taste string, ounces int){
	newCoffee := cupOfCOffee{color: color, taste: taste, ounces: ounces}
	l.coffees = append(l.coffees, newCoffee)
}


//func main for creating program
func main(){
	//Initialize a new collection / light coffee collection
	lightCoffeeCollection := CoffeeCollection{}

	//Append new cup of coffee
	lightCoffeeCollection.addLightCoffee("Light","Bitter",12)

	//Test new collection
	if len(lightCoffeeCollection.coffees) > 0{
		fmt.Printf("The color of the first cup of coffee is %s, the taste is %s, and the total amount is %d ounces.\n", lightCoffeeCollection.coffees[0].color, lightCoffeeCollection.coffees[0].taste, lightCoffeeCollection.coffees[0].ounces)
	} else{
		fmt.Println("There are no coffees in the collection")
	}

}

//go run structs.go