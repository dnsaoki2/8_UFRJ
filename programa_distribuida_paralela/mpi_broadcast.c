#include <stdio.h>
#include <string.h>
#include "mpi.h"

int main(int argc, char** argv) {
	int ranque, finalizado;
	int valor[4];
	MPI_Init(&argc, &argv);
	MPI_Comm_rank(MPI_COMM_WORLD, &ranque);
	valor[0] = ranque+2;
	valor[1] = ranque+2;
	valor[2] = ranque+2;
	valor[3] = ranque+2;
	
	if (ranque == 0) {
		printf("Entre um valor: \n");
		scanf("%d", &valor[0]);
		scanf("%d", &valor[1]);
		scanf("%d", &valor[2]);
		scanf("%d", &valor[3]);
	}
	
	MPI_Bcast(&valor,4,MPI_INT,0,MPI_COMM_WORLD);
	printf("O processo com ranque %d recebeu o valor: %d,%d,%d,%d\n",ranque, valor[0],valor[1],valor[2],valor[3]);

	MPI_Finalized(&finalizado);
	if (!finalizado)
		MPI_Finalize();
}