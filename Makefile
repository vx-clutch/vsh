P=vsh/main.c

all: build

build:
	@gcc -o vsh-release $P -Wall

debug:
	@gcc -o vsh-debug $P -g -Wall

clean:
	@rm vsh-*

run:
	@gcc -o vsh-release $P
	@-./vsh-release
	@rm vsh-release
