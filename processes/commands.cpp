#include "../headers/commands.h"

// Function to be used for build command
int build() {
    return (EXIT_SUCCESS);
}
// Function to be used for install command
int install() {
    return (EXIT_SUCCESS);
}

// Function to be used for url command
int url(char* file) {
    std::string str1 = (std::string) "go run main.go ";
    system((str1 + (std::string) file).c_str());

    return (EXIT_SUCCESS);
}