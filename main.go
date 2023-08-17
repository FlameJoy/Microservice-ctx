package main

import "log"

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	router := NewApiServer(svc)

	log.Fatal(router.Start(":8080"))
}
