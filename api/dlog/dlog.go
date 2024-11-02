package dlog

import "log"

const DEBUG = false

func DLog(v ...interface{}) {
	if DEBUG {
		log.Println(v...)
	}
}
