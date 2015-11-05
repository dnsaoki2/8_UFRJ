#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "mpi.h"

int main(int argc, char** argv) {
	int ranque, i, finalizado;
	int tam_grupo, *buf_envia;
	int raiz, buf_recebe[10];
	raiz = 0;
	MPI_Comm com = MPI_COMM_WORLD;
	MPI_Init(&argc, &argv);
	
	MPI_Comm_size(com, &tam_grupo);
	MPI_Comm_rank(MPI_COMM_WORLD, &ranque);
	
	buf_envia = (int *)malloc(tam_grupo*10*sizeof(int));
	
	for(i = 0; i < tam_grupo*10; i++) {
		buf_envia[i] = i;
	}

	MPI_Scatter(buf_envia, 10, MPI_INT, buf_recebe, 10, MPI_INT, raiz, com);

	printf("Meu ranque=%d ", ranque);
	for(i = 0; i < 10; i++){
		printf("%d ",buf_recebe[i]);
	}
	printf("\n");

	MPI_Finalized(&finalizado);
	if (!finalizado)
		MPI_Finalize();
}