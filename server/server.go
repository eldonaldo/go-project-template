package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	empty = ""
	tab   = "  "
)

// Extracts the argument from the request
func getRequiredArgument(argumentName string, r *http.Request) string {
	arguments, ok := r.URL.Query()[argumentName]
	if !ok || len(arguments[0]) < 1 {
		panic(fmt.Sprintf("Parameter %s is missing", argumentName))
	}

	return arguments[0]
}

// Pretty prints the data to the stream
func PrettyPrint(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent(empty, tab)
	_ = encoder.Encode(data)
}

// Starts the server
func StartServer(port int) {
	// Handlers
	http.HandleFunc("/", MatchHandler)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Println(fmt.Sprintf("Listening: %s ", address))
	log.Fatal(http.ListenAndServe(address, nil))
}
