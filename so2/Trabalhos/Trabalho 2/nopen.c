#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <string.h>
#include <errno.h>


int isopen (int fd){
	struct stat *buff;
	buff = malloc(sizeof(struct stat));
	int open = fstat(fd, buff);
	if(open != -1) {
		return 1;
	}
	return 0;
}

int main (void){
 	FILE *pFile;
 	int nopen, fd;
 	pFile = fopen("test.txt","r");
 	if(pFile == NULL){
 		printf ("Error opening file unexist.ent: %s\n",strerror(errno));
		return 0;
 	}
 	for (nopen = fd = 0; fd < getdtablesize(); fd++) {
 		if (isopen(fd))
			nopen++;
 	}
 	
 	printf ("Existem %d descritores abertos\n", nopen);
 	return 0;
} 