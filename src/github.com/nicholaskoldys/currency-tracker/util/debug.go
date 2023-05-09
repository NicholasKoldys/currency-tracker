package util

import "log"

// Function to print to the console only if the "global variable" is set 3 through 1, else if the verbose-mode is set to 0, send all logs to an external file.

// Enum constants for logging messages where 3 is verbose and 1 is important messages, 0 is disabled and will print only to the external log file.
const (
	DEBUG_MODE = 1
)

const (
	LOG_LEVEL_VERBOSE = 3
	LOG_LEVEL_DEBUG   = 2
	LOG_LEVEL_INFO    = 1
	LOG_LEVEL_WARNING = 0
	LOG_LEVEL_ERROR   = 0
	LOG_LEVEL_FATAL   = 0
	LOG_LEVEL_NONE    = 0
	LOG_LEVEL_ALL     = 0
	LOG_LEVEL_DEFAULT = 0
	LOG_LEVEL_MAX     = 0
	LOG_LEVEL_MIN     = 0
)

func Debug(level int, messages ...string) {
	if level >= DEBUG_MODE {
		for i, message := range messages {
			log.Printf("[DEBUG] %d:%s", i, message)
		}
	}
}
