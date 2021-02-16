package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors)

	// var otherColors map[string]string
	otherColors := make(map[string]string)
	otherColors["white"] = "#ffffff"
	fmt.Println(otherColors)

	otherMap := make(map[int]string)
	otherMap[10] = "#000010"
	fmt.Println(otherMap)

	delete(colors, "green")
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
