package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11352.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}
	funcsESalarios["Rafael Filho"] = 1350.0

	fmt.Println(funcsESalarios)
	delete(funcsESalarios, "Inexistente")

	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}
}
