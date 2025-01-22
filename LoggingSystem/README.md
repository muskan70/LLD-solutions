## Problem Statement:
You have to design and implement a logger library that applications can use to log messages.
Client/application make use of your logger library to log messages to a sink

## Platform Capabilities:
#### Message
1. has content which is of type string
2. has a level associated with it
3. has namespace associated with it to identify the part of application that sent the message
#### Sink
1. This is the destination for a message (e g text file, database, console, etc)
2. Sink is tied to one or more message level
#### Logger library
1. Requires configuration during sink setup
2. Accepts messages from client(s)
3. Routes messages to appropriate sink based on the namespace
4. Supports following message level in the order of priority: FATAL, ERROR WARN, INFO,
DEBUG.
    a. Message levels with higher priority above a given message level should be logged.
    Ex: If INFO is configured as a message level, FATAL, ERROR, WARN and INFO should be logged

5. Enriches message with additional information (like timestamp) while directing message to a sink
6. Should not impact the application flow.
#### Sending messages
1. Sink need not be mentioned while sending a message to the logger library.
3. Client specifies message content, level and namespace while sending a message
#### Logger configuration (see sample below)
1. Specifies all the details required to use the logger library.
2. One configuration per association of message level and sink
3. You may consider logger configuration as a key-value pair
4. Example:

- logging level
- sink type
- details required for sink (eg file location)

#### Bonus Capabilities:
- No information loss. Logger library should log all the messages before the client application
shuts down.
- Log Level for the application can be updated dynamically.

#### Commands:
1. Command to set logging level for application -
Command: SET_LOG_LEVEL <log-level>
2. Command to add a config -
Command: ADD_CONFIG <log-level> <sink-type> <sink-location [optional]>
3. Command to log a message
Command : LOG_MESSAGE <log-level> <message> <namespace>
Output : <sink-type> <level> <timestamp> <message> <namespace>

#### Test Cases:
This doesn’t have to be a command line interface, you have to expose an equivalent function to execute
the following scenarios. Eg - func info(args...), func debug(args...)
1. Command : ADD_CONFIG INFO CONSOLE
Output : NA
2. Command : ADD_CONFIG ERROR FILE
Output : NA
3. Command : ADD_CONFIG DEBUG FILE
Output : NA
4. Command : SET_LOG_LEVEL INFO

5. Command : LOG_MESSAGE INFO “this is a info message” “classroomhandler.go”
Output :
CONSOLE INFO “this is a info message” “classroomhandler.go”
6. Command : LOG_MESSAGE ERROR “this is an error message” “classroomhandler.go”
Output : FILE ERROR “this is an error message” “classroomhandler.go”
7. Command : SET_LOG_LEVEL FATAL
8. Command : LOG_MESSAGE INFO “this is a info message” “classroomhandler.go”
Output : no-output
9. Command : LOG_MESSAGE FATAL “this is a fatal message” “classroomhandler.go”
Output : CONSOLE FATAL “this is a fatal message” “classroomhandler.go”