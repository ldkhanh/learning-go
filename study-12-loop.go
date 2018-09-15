package main 

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	x := 5
	for {
		fmt.Println("Do stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}
	for y := 5; y < 25; y += 3 {
		fmt.Println("Do ", y)
	}
}