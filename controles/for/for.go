package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("Misturando...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("par")
		} else {
			fmt.Println("impar")
		}
	}
	for {
		//laco infinito
	}
}
