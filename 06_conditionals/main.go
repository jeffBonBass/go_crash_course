package main

import "fmt"

func main() {
	x := 15
	y := 15

	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("the color is red")
	} else if color == "blue" {
		fmt.Println("the color is blue")
	} else {
		fmt.Println("the color is Not blue or red")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("the color is red")
	case "blue":
		fmt.Println("the color is blue")
	default:
		fmt.Println("the color is Not blue or red")
	}

}
