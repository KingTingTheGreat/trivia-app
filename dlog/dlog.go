package dlog

import "log"

const DEBUG = true

func DLog(v ...interface{}) {
	if DEBUG {
		log.Println(v...)
	}
}
