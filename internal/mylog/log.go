package mylog

import (
	"log"
)

func LogRequest(remoteAddr, method, url, proto, status string) {
	if status == "500" || status == "404" {
		log.Printf("ERROR %s - - \"%s %s %s\" +%v -\n", remoteAddr, method, url, proto, status)
	}

	log.Printf("INFO %s - - \"%s %s %s\" +%v -\n", remoteAddr, method, url, proto, status)

}

func LogMessage(message string) {
	log.Println(message)
}
