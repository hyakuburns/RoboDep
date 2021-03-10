CC = gcc
CFLAGS = -Wall -Wextra -Wpedantic
LDFLAGS = -lgit2 -lssh2
BIN = xxx
ASM = xxx.S
VERS= -std=c99
PREFIX ?= /usr/local
DBG= debugxxx

all: $(BIN)

#config.h: config.def.h
#	cp config.def.h config.h

SRC = src/main.c#config.h

main.o: src/main.c
	$(CC) -c $(SRC) $(CFLAGS)

OBJ = main.o

$(BIN): $(OBJ)
	$(CC) $(OBJ) -o $(BIN) $(LDFLAGS) $(VERS)

assembly: xxx
	$(CC) -S $(OBJ) -o $(ASM) $(LDFLAGS)

debug: xxx
	$(CC) -g $(OBJ) -o $(DBG) $(LDFLAGS) $(VERS)

#install: xxx
#	mkdir -p ${DESTDIR}${PREFIX}/bin
#	cp -f xxx ${DESTDIR}${PREFIX}/bin

#uninstall:
#	rm -f ${DESTDIR}${PREFIX}/bin/xxx

clean:
	rm -f xxx

.PHONY: all install uninstall clean

