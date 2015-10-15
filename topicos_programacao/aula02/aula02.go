package main

import "fmt"

func main() {
	inicioTempo := time.Now()
	lista := []par{{0, 500}, {30, 510}, {190, 400}, {420, 600}}
	k := 100
	intervalo := par{50, 650}
	fmt.Println(montaHistograma(lista, k, intervalo))
	fmt.Printf("levou: %s\n", time.Since(inicioTempo))
}
