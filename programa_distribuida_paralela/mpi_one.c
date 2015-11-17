#include <stdio.h>
#include <string.h>
#include "mpi.h"

int main(int argc, char** argv) {
	int meu_ranque, np, origem, destino, etiq=0;
	char mensagem[200];
	MPI_Status info;
	MPI_Init(&argc, &argv);
	int iniciado, finalizado;
	double tempo_inicial, tempo_final;
	int versao, subversao;
	char maquina[MPI_MAX_PROCESSOR_NAME];
	int aux, ret;

	MPI_Initialized(&iniciado);
	if (!iniciado)
		MPI_Init(NULL, NULL);
	
	tempo_inicial = MPI_Wtime();
	MPI_Get_version (&versao,&subversao);
	//printf ("Versão do MPI = %d Subversão = %d \n", versao, subversao);
	
	MPI_Comm_rank(MPI_COMM_WORLD, &meu_ranque);
	MPI_Comm_size(MPI_COMM_WORLD,&np);
	if (meu_ranque != 0) {
		sprintf(mensagem, "Processo %d está vivo!", meu_ranque);
		destino = 0;
		//MPI_Send(mensagem,strlen(mensagem)+1,MPI_CHAR,destino,etiq,MPI_COMM_WORLD);
		MPI_Send(&meu_ranque,1,MPI_INT,destino,etiq,MPI_COMM_WORLD);
	} else {
		for (origem=1; origem<np; origem++) {
			MPI_Recv(&meu_ranque, 200, MPI_INT, origem, etiq, MPI_COMM_WORLD, &info);
			printf("%d\n",meu_ranque);
		}
	}
	MPI_Get_processor_name(maquina, &aux);
	//printf ("Número de tarefas = %d Meu ranque = %d Executando em %s\n", np, meu_ranque, maquina);
	tempo_final = MPI_Wtime();
	//printf("Foram gastos %f segundos id: %d\n",tempo_final-tempo_inicial, meu_ranque);
	
	MPI_Finalized(&finalizado);
	if (!finalizado)
		MPI_Finalize();
}