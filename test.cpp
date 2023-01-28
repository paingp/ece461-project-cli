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
    if(install() == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test();
}

// Test Suite for URL command
void url_tests(int* tests_total, int* tests_passed, char* filename) {
    if(url(filename) == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test();
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

    // Testing install command
    int install_total = 0; // Total number of build tests executed
    int install_passed = 0; // Total number of build tests passed
    build_tests(&install_total, &install_passed);
    log_command((char*)"install");
    tests_total += install_total;
    tests_passed += install_passed;

    // Testing url command
    int url_total = 0; // Total number of build tests executed
    int url_passed = 0; // Total number of build tests passed
    build_tests(&url_total, &url_passed);
    log_command((char*)"url");
    tests_total += url_total;
    tests_passed += url_passed;


    // Outputting to stdout
    fprintf(stdout, "%d/%d test cases passed. %d%% line coverage achieved.", tests_passed, tests_total, 0);

    return EXIT_SUCCESS;
}