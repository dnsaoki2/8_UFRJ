#include <unistd.h>
#include <sys/stat.h>
#include <fcntl.h>

int main (){
	char c;
	int in, out; /* serão usados como descritores de arquivos */
	in = open("file.in", O_RDONLY);
	out = open("file.out", O_WRONLY|O_CREAT, S_IRGRP|S_IWGRP);
	while (read(in,&c,1) == 1)
		write (out,&c,1);
}