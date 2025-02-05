package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!!")

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d は偶数です\n", i)
		} else {
			fmt.Printf("%d は奇数です\n", i)
		}
	}
}
