package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	x, y := 2, 3

	x, y = troca(x, y)
	fmt.Println(x, y)

}
