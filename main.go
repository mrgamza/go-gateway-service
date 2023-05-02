package main

import (
	"log"
	"net/http"
	
	"gateway-service/route"
	"github.com/julienschmidt/httprouter"
)

func main() {
	route.New()
	
	router := httprouter.New()
	router.GET("/", route.Index)
	router.GET("/php/*path", route.Php)
	router.GET("/java/*path", route.Java)
	router.GET("/node/*path", route.Node)
	router.GET("/ruby/*path", route.Ruby)
	router.GET("/python/*path", route.Python)
	router.GET("/go/*path", route.Go)
	router.GET("/kotlin/*path", route.Kotlin)
	
	log.Fatal(http.ListenAndServe(":80", router))
}
