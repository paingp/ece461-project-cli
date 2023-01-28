#include "../headers/logger.h"

char logFile[18] = "Logs/log.txt";
std::string logLevel = "1";

// Log for individual test case
void log_test(){

}

// Log for entire command test
void log_command(char* command){

}

//////////////////////////////////////
// 
// Logger Verbosity Level 1
//
// Outputs:
//  * Command success or failure at timestamp
//  * Command Ran
//     
//////////////////////////////////////
void loggerUpdateOne(char *command){
    FILE *fileptr; 
    fileptr = fopen(logFile, "a");

    if (!fileptr){
        fclose(fileptr);
        printf("\nError logging, log file pointer was NULL.");
        exit(EXIT_FAILURE);
    }

    time_t t;
    time(&t);
    
    fputs("Command Successful:    " , fileptr); 
    fputs(ctime(&t), fileptr);
    fputs("Command ran was:       ./run ", fileptr);
    fputs(command , fileptr);
    fputs("\n\n" , fileptr);

    fclose(fileptr);
}

void logError(int errorType, char* command) {

    FILE *fileptr; 
    fileptr = fopen(logFile, "a");

     if (!fileptr){
        fclose(fileptr);
        printf("Error logging, log file pointer was NULL");
        exit(EXIT_FAILURE);
    }

    time_t t;
    time(&t);

    switch(errorType){
        case 1: // Invalid run command 
            fputs("Command Failed:        " , fileptr); 
            fputs(ctime(&t), fileptr);
            fputs("Command ran was:       ./run ", fileptr);
            fputs(command , fileptr);
            fputs("\n\n" , fileptr);
            break;
        default:
            fputs("Unknown error in execution " , fileptr);
            fputs("\n\n" , fileptr);
    }

    fclose(fileptr);
}

int loggerMain(char* command){

    if (stoi(logLevel) == 1){
        loggerUpdateOne(command);
    } else if (stoi(logLevel) == 2) {
        //int updateSuccess = loggerUpdateTwo();
        printf("s");
    } 

    return 0;
}