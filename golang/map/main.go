package main

import "fmt"

func main() {

	colors := map[string]string {
		"red": "#ff0000",
		"blue": "#gg0000",
		"white": "ffffff",
	}
	printMap(colors)
}

func printMap (c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hexcode for ", color, "is ", hex )
	} 
}