package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		log.Println("start request")
		_, err := http.Get("http://localhost:3000/ping")
		if err != nil {
			log.Println("request error: ", err)
			break
		}
		// time.Sleep(1 * time.Second)
		log.Println("resuest succesfull")
		time.Sleep(2 * time.Second)
	}
}
