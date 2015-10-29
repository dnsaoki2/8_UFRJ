#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

int main (){
	// if ((write(1, "Alo Mundo\n", 18)) != 18){
	// 	write (2, "write error on file descriptor 1\n", 33);
	// }
	char buffer[128];
	int nread;
	nread = read(0, buffer, 128);
	if (nread == -1)
		write (2, "a read error has occurred\n", 26);
	if ((write(1, buffer, nread)) != nread)
		write (2, "a write error has occurred\n", 27);
	
	return 0;
}
