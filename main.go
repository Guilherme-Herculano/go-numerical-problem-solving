package main

import "fmt"

func main() {
	// part 1
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	println("-----")
	// part 2
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d Pin Pan\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Pin\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Pan\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
