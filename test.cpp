#include "test.h"

// Test Suite for Build command
void build_tests(int* tests_total, int* tests_passed) {
    if(build() == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test();
}

// Test Suite for Install command
void install_tests(int* tests_total, int* tests_passed) {

}

// Test Suite for URL command
void url_tests(int* tests_total, int* tests_passed, char* filename) {
    (*tests_total)++;
    (*tests_passed)++;
}

int master_test() {
    int tests_total = 0; // Total number of tests executed
    int tests_passed = 0; // Total number of tests passed

    // Testing build command
    int build_total = 0; // Total number of build tests executed
    int build_passed = 0; // Total number of build tests passed
    build_tests(&build_total, &build_passed);
    log_command((char*)"build");
    tests_total += build_total;
    tests_passed += build_passed;


    // Outputting to stdout
    fprintf(stdout, "%d/%d test cases passed. %d%% line coverage achieved.", tests_passed, tests_total, 0);

    return EXIT_SUCCESS;
}