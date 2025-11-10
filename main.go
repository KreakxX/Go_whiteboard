package main

import "fmt"

func main() {
	fmt.Println("Hello Go")

	var x int = 10
	x += 10

	fmt.Println(x)

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

	go func() {
		fmt.Println("hello")
	}()

}

func add(a, b int) int {
	return a + b
}
