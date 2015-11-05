#include <sys/types.h>
#include <sys/stat.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <ctype.h>
#include <unistd.h>

int countRegular = 0;
int countDiretorio = 0;
int countSimbolico = 0;
int countEstruturado = 0;
int countNEstruturado = 0;

int walk_dir (const char *path, void (*func) (const char *)){
	DIR *dirp;
	struct dirent *dp;
	char *p, *full_path;
	int len;
	
	/* abre o diretório */
	if ((dirp = opendir(path)) == NULL)
		return (-1);
	
	len = strlen(path);
	/* aloca uma área na qual, garantidamente, o caminho caberá */
	if ((full_path = malloc(len + MAXNAMLEN + 2)) == NULL) {
		closedir (dirp);
	 	return (-1);
	}
	
	/* copia o prefixo e acrescenta a ‘/’ ao final */
	memcpy (full_path, path, len);
	p = full_path + len; *p++ = '/'; /* deixa “p” no lugar certo! */
	while ((dp = readdir (dirp)) != NULL) {
		/* ignora as entradas “.” e “..” */
	 	if (strcmp (dp->d_name, ".") == 0 || strcmp (dp->d_name, "..") == 0){
	 		continue;
	 	}
	 	strcpy (p, dp->d_name);

	 	/* "full_path" armazena o caminho */
	 	(*func)(full_path);
	}
	 
	free(full_path);
	closedir(dirp);
	return (0);
} 

void func(const char *path) {
	//printf("Verificando o tipo do path: %s\n", path);
	struct stat sb;
	if(stat(path, &sb) == -1){
		perror("stat");
        return;
	}
	switch (sb.st_mode & S_IFMT) {
	    case S_IFBLK:  
	    	countEstruturado++;		
	    	break;
	    case S_IFCHR:  
	    	countNEstruturado++;		
	    	break;
	    case S_IFDIR:  
	    	countDiretorio++;    	
	    	break;
	    case S_IFLNK:  
	    	countSimbolico++;    	
	    	break;
	    case S_IFREG:  
	    	countRegular++;      	
	    	break;
    }
}

int main(int argc, char *argv[]) {
	if(argc == 1) {
		walk_dir(".", &func);
		printf("Para o path: \".\" :\n");	
		printf("...Regular: %d\n",countRegular);
		return 0;
	}
	if(argc == 2 && strcmp(argv[1],"man") == 0){
 		printf("Manual do icount\n\n");
 		printf("USO: icount [-rdlbc] [<dir> ...] \n\n");
 		printf("Conta a quantidade de arquivos do tipo informado \n\n");
 		printf("Flags: \n\n");
 		printf("   -r: arquivo regular\n");
 		printf("   -d: diretório\n");
 		printf("   -l: elo simbólico\n");
 		printf("   -b: dispositivo estruturado\n");
 		printf("   -c: dispositivo não-estruturado\n\n");
 		return 0;
 	}
 	int c;
	int regular, diretorio, simbolico, estruturado, nestruturado,flags = 0;
	opterr = 0;
	while((c = getopt(argc, argv, "rdlbc")) != -1){
		flags = 1;
		switch (c) {
			case 'r':
	        	regular = 1;
	        	break;
	    	case 'd':
	        	diretorio = 1;
	        	break;
	     	case 'l':
	        	simbolico = 1;
	        	break;
	        case 'b':
	        	estruturado = 1;
	        	break;
	        case 'c':
	        	nestruturado = 1;
	        	break;
	        case '?':
	        	if(isprint(optopt)) {
	      			printf("Flag invalida \n\n");
	      			printf("USE: \"icount man\" para mais informações\n\n");
	      		}
	      		return 1;
	      	default:
	      		abort();
	    }
	}
	int inicio = 2;
	if(flags == 0){
	 	inicio = 1;
	 	regular = 1;
	} 
	if(flags == 1 && argc == 2){
		printf("Falta argumentos \n\n");
		printf("USE: \"icount man\" para mais informações\n\n");
	}
	
	for(int i = inicio; i < argc; i++) {
		printf("Para o path: \"%s\"\n", argv[i]);
		countRegular = 0; countDiretorio = 0; countSimbolico = 0; countEstruturado = 0; countNEstruturado = 0;
		walk_dir(argv[i], &func);
		if(regular) {
			printf("...Regular: %d\n", countRegular);
		}
		if(diretorio) {
			printf("...Diretório: %d\n", countDiretorio);
		}	
		if(simbolico) {
			printf("...Simbolico: %d\n", countSimbolico);
		}
		if(estruturado) {
			printf("...Estruturado: %d\n", countEstruturado);	
		}
		if(nestruturado) {
			printf("...Não Estruturado: %d\n", countNEstruturado);	
		}
	}
 	return 0;
}