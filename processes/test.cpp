#include "../headers/test.h"

// Test Suite for Build command
void build_tests(int* tests_total, int* tests_passed) {
    int status;
    if((status = build()) == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test("build()", status);
}

// Test Suite for Install command
void install_tests(int* tests_total, int* tests_passed) {
    int status;
    if((status = install()) == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test("install()", status);
}

// Test Suite for URL command
void url_tests(int* tests_total, int* tests_passed, char* filename) {
    int status;
    if((status = url(filename)) == EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test("url()", status);
}

// Test Suit for Logger
void logger_tests(int* tests_total, int* tests_passed) {
    int status;
    if((status = log_test("log_test()", PASS))== EXIT_SUCCESS) {
        (*tests_passed)++;
    }
    (*tests_total)++;
    log_test("log_test()", status);
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
    install_tests(&install_total, &install_passed);
    log_command((char*)"install");
    tests_total += install_total;
    tests_passed += install_passed;

    // Testing url command
    int url_total = 0; // Total number of build tests executed
    int url_passed = 0; // Total number of build tests passed
    //url_tests(&url_total, &url_passed, (char*)"sample.txt");
    log_command((char*)"url");
    tests_total += url_total;
    tests_passed += url_passed;

    // Testing log command
    int log_total = 0; // Total number of build tests executed
    int log_passed = 0; // Total number of build tests passed
    logger_tests(&log_total, &log_passed);
    log_command((char*)"log");
    tests_total += log_total;
    tests_passed += log_passed;


    system("go test ./ratom > temp.txt");
    //system("rm temp.txt");

    int coverage = 0;
    // Outputting to stdout;
    fprintf(stdout, "Total: %d\n", tests_total);
    fprintf(stdout, "Passed: %d\n", tests_passed);
    fprintf(stdout, "Coverage %d%%\n", coverage);
    fprintf(stdout, "%d/%d test cases passed. %d%% line coverage achieved.\n", tests_passed, tests_total, coverage);

    return EXIT_SUCCESS;
}