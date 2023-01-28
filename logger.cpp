#include "logger.h"

//////////////////TODO//////////////////
// char mypath[]="LOG_FILE=log.txt"; 
// putenv( mypath ); 

//std::string logFile = getenv("LOG_FILE");
//std::string logLevel = getenv("LOG_LEVEL");
/////////////////////////////////////////

//std::string logFile;
//std::string logLevel;
//int logLevel = "1";

// Logs an individual test: TODO
void log_test() {
    return;
}

// Logs aggregate of tests for command: TODO
void log_command(char* command) {
    return;
}

/*
int loggerUpdateOne(char *command){
    FILE *fileptr; 
    char* logFile = "log.txt";
    fileptr = fopen(logFile, "a");

    if (!fileptr){
        fclose(fileptr);
        return -1;
    }

    fputs("Command Successful:\n" , fileptr);
    fputs(command , fileptr);

    time_t t;
    time(&t);
    fputs(ctime(&t), fileptr);
    
    return 0;
}

int loggerMain(char *command){
    
    if (stoi(logLevel) == 1){
        int updateSuccess = loggerUpdateOne(command);
    } else if (stoi(logLevel) == 2) {
        //int updateSuccess = loggerUpdateTwo();
        printf("s");
    } 
    
    return 0;
}
*/