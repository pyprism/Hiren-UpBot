package main

import (
	"log"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "login")
}

func main() {
	router := httprouter.New()
	router.GET("/", Login)

	PORT := "127.0.0.1:8000"
	fmt.Println("Server running on:", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}


