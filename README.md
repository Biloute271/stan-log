# Stan Log

## Purpose

The purpose of this simple module is to write to stdout logs with human readable timestamps
The module proposes 5 log levels :

- DEBUG
- INFO
- WARNING
- ERROR
- CRITICAL

If the loglevel of the event to log is equal or greater than the loglevel defined on the application-level, an entry is printed to stdout

## How to use

### Import
Firs of all, import the package
> Import (
> ...
> stanlog "github.com/Biloute271/stan-log"
> ...
> )
### Define log level

The default log level is set to **WARNING**
THe log level can be defined in several ways :

#### Generic funciton

It is achieved by calling SetLogLevel function and sending a loglevel string as parameter. In this way, the desired logging level can be stored as a string in a configuration file of your Go piece of software.

Examples:

To log events of levels WARNING, ERROR and CRITICAL:
> stanlog.SetLogLevel("WARNING")

To log only events of level CRITICAL
> stanLog.setLogLevel("CRITICAL")

#### SpecificFunctions

The logLevel can also be defined by calling specific functions (and preventing errors resulting from uexpected strings manipulations).

Examples:
> stanlog.SetLogLevelDebug
> stanlog.SetLogLevelInfo
> stanlog.SetLogLevelWarning
> stanlog.SetLogLevelError
> stanlog.SetLogLevelCritical

### Use in your code

After defining the desired log level (or staying at default 'WARNING'), each event that is exepected to be logged (or not) on stdout is logged as parameter string.

Examples:
> stanlog.Debug("This is an event useful for debugging, but not really useful in everyday operations")
> stanlog.Info("This information is not an error, but can be useful in certain circumstances. For example, for reglementary purposes in an application dealing with banking tranfers")
> stanlog.Warning("Oh, something strange happened")
> stanlog.Error("Some error occurred")
> stanlog.Critical("Something is definitely wrong)

## Formatting thoughts

The output to stdout has 3 fixed-width fields

- Timestamp: 26 characters: *yyyy-MM-DD hh.mm.ss.dddddd*
- Log Level: 8 characters (including trailing spaces)
- The published message itself

Fieds are separated by " - "

In this way, if the logs have to be retrieved by an external tool, the importer can be confugured as follows :

- Timestamp : columns 1 to 27
- Message level : columns 30 to 38
- Message itself : column 41 and above
