package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"white": "ffffff",
	}

	// colors := make(map[string]string)
	colors["blue"] = "000000"
	// delete(colors, "green")
	// fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("key:", color, "value:", hex)
	}
}
