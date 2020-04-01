package utils

import (
	"log"
	"net/http"
	"os"
)

//LogWrapper is middleware function for loggin all RESTapi queries
type LogWrapper struct {
	http.Handler
}

func (wr LogWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logFile, err := os.OpenFile("goreactapp.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Printf("%s %s %s\n", r.Method, r.URL, ReadUserIP(r))
	wr.Handler.ServeHTTP(w, r)
}
