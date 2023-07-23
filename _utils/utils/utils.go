package utils

import (
	"log"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func CheckErrAbortProgram(err error, msg string) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}
