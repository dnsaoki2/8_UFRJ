package main

import "testing"

var (
	k         = 100
	intervalo = par{50, 650}
	lista     = []par{{0, 500}, {30, 510}, {190, 400}, {420, 600}}
)

func TestIniciaMapa(t *testing.T) {
	result := iniciaMapa(intervalo, k)
	for _, valor := range result {
		if valor != 0 {
			t.Error("Esperado 0, obteve ", valor)
		}
	}
}

func TestContaJogoresIngenuo(t *testing.T) {
	result := montaHistogramaIngenuo(lista, k, intervalo)
	if result[50] != 2 {
		t.Error("Esperado 2, obteve ", result[50])
	}
	if result[150] != 2 {
		t.Error("Esperado 2, obteve ", result[150])
	}
	if result[250] != 3 {
		t.Error("Esperado 3, obteve ", result[250])
	}
	if result[350] != 3 {
		t.Error("Esperado 3, obteve ", result[350])
	}
	if result[450] != 3 {
		t.Error("Esperado 3, obteve ", result[450])
	}
	if result[550] != 1 {
		t.Error("Esperado 1, obteve ", result[550])
	}
	if result[650] != 0 {
		t.Error("Esperado 0, obteve ", result[650])
	}
}

func TestContaJogoresOrdenado(t *testing.T) {
	result := montaHistogramaOrdenado(lista, k, intervalo)
	if result[50] != 2 {
		t.Error("Esperado 2, obteve ", result[50])
	}
	if result[150] != 2 {
		t.Error("Esperado 2, obteve ", result[150])
	}
	if result[250] != 3 {
		t.Error("Esperado 3, obteve ", result[250])
	}
	if result[350] != 3 {
		t.Error("Esperado 3, obteve ", result[350])
	}
	if result[450] != 3 {
		t.Error("Esperado 3, obteve ", result[450])
	}
	if result[550] != 1 {
		t.Error("Esperado 1, obteve ", result[550])
	}
	if result[650] != 0 {
		t.Error("Esperado 0, obteve ", result[650])
	}
}
