CC = g++
CFLAGS = -Wall -Wextra -Wpedantic
LDFLAGS = #-lraylib -lGL -lm -lpthread -ldl -lrt -lX11 -lglfw -lvulkan -lXxf86vm -lXrandr -lXi
BIN = xxx
ASM = xxx.S
VERS= -std=c++17
PREFIX ?= /usr/local
DBG = debuxxx


all: $(BIN)

#config.h: config.def.h
#	cp config.def.h config.h

SRC = src/main.cxx#config.h

main.o: src/main.cxx
	$(CC) -c $(SRC) $(CFLAGS)

OBJ = main.o

$(BIN): $(OBJ)
	$(CC) $(OBJ) -o $(BIN) $(LDFLAGS) $(VERS)

assembly: xxx
	$(CC) -S $(OBJ) -o $(ASM) $(LDFLAGS)

debug: xxx
	$(CC) -g $(SRC) -o $(DBG) $(LDFLAGS) $(VERS)

#install: xxx
#	mkdir -p ${DESTDIR}${PREFIX}/bin
#	cp -f xxx ${DESTDIR}${PREFIX}/bin

#uninstall:
#	rm -f ${DESTDIR}${PREFIX}/bin/xxx

clean:
	rm -f xxx

.PHONY: all install uninstall clean
