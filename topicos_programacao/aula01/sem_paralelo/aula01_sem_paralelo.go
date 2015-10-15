//Entrada: -Lista de n pares(inicio, fim)
//         -início global e fim global (do histograma)
//         -granularidade(tamanho da "coluna")
//Saída:   -mapa{horário: quantidade de jogadores online}
//EX : [(0,500),(30,510),(190,400),(420,600)]
//    (inicio global, fim global) : (50, 650)
//    granularidade: 100
//Saida: 50: 2 ; 150: 2 ; 250: 3 ; 350: 3 ; 450: 3; 550: 1 ; 650: 0

package main

import (
	"fmt"
	"time"
)

type par struct {
	inicio int
	fim    int
}

func iniciaMapa(intervalo par, k int) map[int]int {
	mapa := make(map[int]int)
	for n := intervalo.inicio; n <= intervalo.fim; n += k {
		mapa[n] = 0
	}
	return mapa
}

func montaHistogramaIngenuo(lista []par, k int, intervalo par) map[int]int {
	mapa := iniciaMapa(intervalo, k)
	for _, parLista := range lista {
		for x := intervalo.inicio; x <= intervalo.fim; x += k {
			if x >= parLista.inicio && x <= parLista.fim {
				mapa[x]++
			}
		}
	}
	return mapa
}

func montaHistogramaOrdenado(lista []par, k int, intervalo par) map[int]int {
	mapa := iniciaMapa(intervalo, k)
	for x := intervalo.inicio; x <= intervalo.fim; x += k {
		for _, parLista := range lista {
			if x < parLista.inicio {
				break
			}
			if x <= parLista.fim {
				mapa[x]++
			}
		}
	}
	return mapa
}

func main() {
	inicioTempo := time.Now()
	lista := []par{{0, 500}, {30, 510}, {190, 400}, {420, 600}}
	k := 100
	intervalo := par{50, 650}
	fmt.Println(montaHistogramaIngenuo(lista, k, intervalo))
	fmt.Printf("levou: %s\n", time.Since(inicioTempo))
}
