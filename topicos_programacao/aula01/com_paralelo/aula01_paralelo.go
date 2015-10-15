package main

import (
	"fmt"
	"sync"
	"time"
)

type par struct {
	inicio int
	fim    int
}

var wait sync.WaitGroup

func iniciaMapa(intervalo par, k int) map[int]int {
	mapa := make(map[int]int)
	for n := intervalo.inicio; n <= intervalo.fim; n += k {
		mapa[n] = 0
	}
	return mapa
}

func contaJogadores(tempo int, lista []par) int {
	contador := 0
	for _, par := range lista {
		if tempo >= par.inicio && tempo <= par.fim {
			contador++
		}
	}
	return contador
}

func montaHistograma(lista []par, mapa map[int]int) map[int]int {
	for tempo, _ := range mapa {
		copia := tempo
		wait.Add(1)
		go func() {
			mapa[copia] = contaJogadores(copia, lista)
			wait.Done()
		}()
	}
	wait.Wait()
	return mapa
}

func main() {
	inicioTempo := time.Now()
	lista := []par{{0, 500}, {30, 510}, {190, 400}, {420, 600}}
	k := 100
	intervalo := par{50, 650}
	mapa := iniciaMapa(intervalo, k)
	fmt.Println(montaHistograma(lista, mapa))
	fmt.Printf("levou: %s\n", time.Since(inicioTempo))
}
