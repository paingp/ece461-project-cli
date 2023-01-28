#include "../headers/cli.h"

using namespace std;

int test() {
    return master_test();
}

int main(int argc, char *argv[]) {
    // Input must have at least 2 parameters
    if(argc <= 1) {
        exit(EXIT_FAILURE);
    }

    if (strcmp(argv[1], "install") == 0){
        // Install command 
        if(install() == EXIT_FAILURE) {
            exit(EXIT_FAILURE);
        }

        loggerMain(argv[1]);
    }
    else if (strcmp(argv[1], "build") == 0){
        if(build() == EXIT_FAILURE) {
            exit(EXIT_FAILURE);
        }

        loggerMain(argv[1]);
    }
    else if(strcmp(argv[1], "test") == 0) {
        if(test() == EXIT_FAILURE) {
            exit(EXIT_FAILURE);
        }

        loggerMain(argv[1]);
    }
    else {
        // Validating file name input
        FILE *fileptr; 
        fileptr = fopen(argv[1], "r");

        if (!fileptr) {
            free(fileptr);
            fprintf(stdout, "Command did not match any of the following: \n./run build\n./run install\n./run URL_FILE");
            fprintf(stdout, "\nor the file could not be found in the given path.");  
            logError(1, argv[1]);
            exit(EXIT_FAILURE);

        }

        // Perform running 
        fclose(fileptr);

        // Running file analysis
        url(argv[1]);    

        // Calling logger to log the command that was just run
        char *runURL = (char *)"RUN_URL";
        loggerMain(runURL);
    }
     
    exit(EXIT_SUCCESS);
}