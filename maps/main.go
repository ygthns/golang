package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf721",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff" //for append elements
	// delete(colors, "white")     //for delete the key values

	printMap(colors)
}

//map[string]int---> means keys will be string, and value will be integer.

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for ", key, " is ", value)
	}
}

//all keys and values must be the same type, unlike structs
//maps are reference type, structs are value type.
//we can add, or delete single elements from maps, but we need to fill all the different fields at compile time for structs.
