#include <stdio.h>
#include <stdlib.h>
#include <string.h>


using namespace std;

// Function to be used for install command
int install() {
    return (EXIT_SUCCESS);
}

// Function to be used for build command
int build() {
    return (EXIT_SUCCESS);
}

int test() {
    return (EXIT_SUCCESS);
}

int url(char* file) {
    return (EXIT_FAILURE);
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
    }
    else if (strcmp(argv[1], "build") == 0){
        if(build() == EXIT_FAILURE) {
            exit(EXIT_FAILURE);
        }
    }
    else {
        // Validating file name input
        FILE *fileptr; 
        fileptr = fopen(argv[1], "r");

        if (!fileptr) {
            free(fileptr);
            printf("Command did not match any of the following: \n./run build\n./run install\n./run URL_FILE");
            printf("\nor the file could not be found in the given path.");
            exit(EXIT_FAILURE);
        }
        
        // Perform running 
        fclose(fileptr);

        // Running file analysis
        url(argv[1]);    
    }
     
    exit(EXIT_SUCCESS);
}