package main

import "fmt"

func main() {
	i := 10
	if i >= 5 {
		fmt.Println("5以上")
	} else {
		fmt.Println("5未満")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
