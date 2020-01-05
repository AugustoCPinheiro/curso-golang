package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando endereco da variavel
	*p++   // desreferenciacao (pega o valor)
	i++

	fmt.Println(p, *p, i)

}
