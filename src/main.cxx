#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>

#define COLOR1 "\033[38;5;214m"
#define COLOR2 "\033[38;5;147m"
#define COLOR3 "\033[38;5;44m"
#define RED "\033[38;5;196m"
#define ORANGE "\033[38;5;208m"
#define RESET "\033[0m"

void checkErr() {
    if (EACCES == errno) {
        printf("No permissions cuz\n");
        exit(EXIT_FAILURE);
    }
}

void makeRBDIR() {
    struct stat fakeSt = {0};
    if (-1 == stat("./robodep", &fakeSt)) {
        mkdir("./robodep", 0700);
        checkErr();
    } else {
        printf(COLOR1 "robodep" RESET " directory already exists.\n");
    }
}

void exitWRetVal(int retVal) {
    if (0 != retVal) {
        exit(EXIT_FAILURE);
    }
}

void initRBD(int argc, char **argv) {
    if (1 == argc) {
        printf("You literally didnt even try.\n");
        exit(EXIT_FAILURE);
    }
    if (1 < argc && 0 != strcmp(argv[1], "get")) {
        printf("kys\n");
    } else {
        printf("yous hould still kys \n");
    }
    printf("\n%s\n\n", argv[1]);
    char cloneroo[] = "git -C ./robodep clone https://github.com/hyakuburns/milfs";
    makeRBDIR();
    int retVal = system(cloneroo);
    exitWRetVal(retVal);
}

int main(int argc, char **argv) {
    initRBD(argc, argv);
}
