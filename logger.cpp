#include "logger.h"

char logFile[10] = "log.txt";
std::string logLevel = "1";

// Log for individual test case
void log_test(){

}

// Log for entire command test
void log_command(char* command){

}

int loggerUpdateOne(char *command){
    FILE *fileptr; 
    fileptr = fopen(logFile, "a");

    if (!fileptr){
        fclose(fileptr);
        return -1;
    }

    time_t t;
    time(&t);
    
    fputs("Command Successful:    " , fileptr); 
    fputs(ctime(&t), fileptr);
    fputs(command , fileptr);

    fclose(fileptr);
    return 0;
}

int loggerMain(char* command){

    if (stoi(logLevel) == 1){
        int updateSuccess = loggerUpdateOne(command);
    } else if (stoi(logLevel) == 2) {
        //int updateSuccess = loggerUpdateTwo();
        printf("s");
    } 
    
    return 0;
}