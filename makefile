P=vsh/main.c
O=bin/

all: build

build:
	@gcc -o $Ovsh-release $P -Wall

debug:
	@gcc -o $Ovsh-debug $P -g -Wall

clean:
	@-rm vsh-*
	@-rm bin/vsh-*

run:
	@gcc -o vsh-release $P
	@-./vsh-release
	@rm vsh-release
