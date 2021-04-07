package server

import (
	"fmt"
	"github.com/eldonaldo/go-project-template/user"
	"log"
	"net/http"
)

// Starts the server
func StartServer(port int) {

	// Handlers
	http.HandleFunc("/create", user.CreateUser)
	http.HandleFunc("/greet", user.GreetHandler)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println(fmt.Sprintf("Listening: %s ", address))
	log.Fatal(http.ListenAndServe(address, nil))
}
