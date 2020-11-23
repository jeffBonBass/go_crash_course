package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop thru ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// just print the kev
	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
