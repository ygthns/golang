package main

import "fmt"

func evenodd(n []int) {
	for i, is := range n {
		if is%2 != 0 {
			fmt.Println(i, " is odd")
		} else {
			fmt.Println(i, " is even")
		}
	}
}

func main() {
	numbersList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenodd(numbersList)
}
