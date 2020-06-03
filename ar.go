package atomic

import (
	"log"
)

func LogFatal(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func LogError(msg string, err error) bool {
	if err != nil {
		log.Println(msg, err)
		return true
	}
	return false
}
