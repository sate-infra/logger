package logger

import "log"

func Errorf(format string, v ...any) {
	format = string(COLOR_RED) + "ERROR: " + format + string(COLOR_RESET)
	log.Printf(format, v...)
}
