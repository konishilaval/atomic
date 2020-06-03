package atomic

import (
	"log"
)

func logFatal(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func logError(msg string, err error) bool {
	if err != nil {
		log.Println(msg, err)
		return true
	}
	return false
}
