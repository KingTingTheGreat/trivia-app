package dlog

import "log"

const DEBUG = false

func DLog(v ...any) {
	if DEBUG {
		log.Println(v...)
	}
}
