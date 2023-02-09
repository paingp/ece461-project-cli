#include "../headers/logger.h"

// char* logFile = getenv("LOG_FILE");
// char* logLevel = getenv("LOG_LEVEL");

//char logFile[18] = "Logs/log.txt";
char* logFile = getenv("LOG_FILE");
char* logLevel = getenv("LOG_LEVEL");

///////////////////////////////////////////
//
//  log_test
//
//  takes in function name as string
//  and status (pass/fail) as int 
//  (0 = pass, 1 = fail)
//
///////////////////////////////////////////
int log_test(std::string function, int status){

    return EXIT_SUCCESS;
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
    // fputs("\n\n" , fileptr);

    fclose(fileptr);
}

void loggerUpdateTwo(char *command, std::vector<std::string> lineNumbers){
    std::string lineN = std::to_string(__LINE__ - 1);

    FILE *fileptr; 
    fileptr = fopen(logFile, "a");
    
    lineN = "Function: loggerUpdateTwo in logger.cpp logging information on line " + lineN;
    lineNumbers.push_back(lineN);

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
    fputs("________________Debugging Information________________\n", fileptr);

    for (int i = 0; i < lineNumbers.size(); i++){
        const char* fileStr = lineNumbers[i].c_str();
        fputs(fileStr, fileptr);
        fputs("\n" , fileptr);
    }

    //fputs("\n\n" , fileptr);

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

int loggerMain(char* command, std::vector<std::string> lineNumbers){

    if (atoi(logLevel) == 1){
        loggerUpdateOne(command);
    } else if (atoi(logLevel) == 2) {
        std::string lineN = std::to_string(__LINE__ + 3);
        lineN = "Function: loggerMain  in logger.cpp creating log (verbosity 2) with function loggerUpdateTwo on line " + lineN;
        lineNumbers.push_back(lineN);
        loggerUpdateTwo(command, lineNumbers);
    } 

    return 0;
}