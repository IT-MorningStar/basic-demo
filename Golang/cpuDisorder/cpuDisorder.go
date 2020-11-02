package main

import "fmt"

func main() {
	var i int
	for {
		var a, b int = 5, 5
		var x, y int
		go func() {
			a = 12
			y = b
		}()

		go func() {
			b = 13
			x = a
		}()

		if x == 5 && y == 5 {
			fmt.Printf("cpu乱序存在:执行了%d次", i)
			break
		}
		i++
	}
}
