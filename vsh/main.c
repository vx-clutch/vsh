#include <stdlib.h>
#include <stdio.h>

char * prompt() {
	char *input = malloc(25);
	puts("(vsh) => ");
	scanf(input);
	puts(input);
	return input;
}

int main() {
	prompt();
	return 0;
}
