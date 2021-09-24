package main

import "fmt"

func main() {
	calcAdd := calc(2, 2, "add")
	calcAddition := calc(3, 6, "add")
	calcMult := calc(7, 7, "mult")
	calcDiv := calc(9, 3, "div")
	fmt.Printf("2 + 2 = %d, 3 + 6 = %d, 7 * 7 = %d, 9 / 3 = %d", calcAdd, calcAddition, calcMult, calcDiv)
}

func calc(a int, b int, t string) int {
	ch := make(chan int)
	switch t {
	case "add":
		go addition(a, b, ch)
	case "subs":
		go substraction(a, b, ch)
	case "mult":
		go multiplication(a, b, ch)
	case "div":
		go division(a, b, ch)
	}

	return <-ch
}

func addition(a, b int, ch chan int) {
	ch <- a + b
}

func substraction(a, b int, ch chan int) {
	ch <- a - b
}

func multiplication(a, b int, ch chan int) {
	ch <- a * b
}

func division(a, b int, ch chan int) {
	ch <- a / b
}
