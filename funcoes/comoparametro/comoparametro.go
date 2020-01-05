package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(fun func(int, int) int, p1, p2 int) int {
	return fun(p1, p2)
}

func main() {
	fmt.Println(exec(multiplicacao, 2, 4))
}
