package main

import "fmt"

func main() {
	var a [6]int
	fmt.Println("emp ", a)

	a[5] = 400
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4, 4}
	fmt.Println(b)

	c := [10]int{0, 1, 3, 4, 4, 5, 8}
	fmt.Println(c)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	d := [...]int{100, 3: 20, 4: 80, 10: 90, 30}
	fmt.Println(d)
}
