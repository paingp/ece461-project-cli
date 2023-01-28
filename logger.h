#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string>

void log_test();
void log_command(char* command);
void loggerUpdateOne(char *command);
void logError(int errorType, char* command);
int loggerMain(char *command);
