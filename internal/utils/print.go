package utils

import (
	"fmt"
	"log"
)

func FailOnError(err error, msg string) {
	if err == nil {
		return
	}

	log.Panicf("%s: %s", msg, err)
}

func PrintOnDebug(debug bool, msg string) {
	if !debug {
		return
	}

	fmt.Print(msg)
}
