package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	empty = ""
	space = "  "
)

// Extracts the required argument from the HTTP request
func GetRequiredArgument(argumentName string, r *http.Request) string {
	arguments, ok := r.URL.Query()[argumentName]
	if !ok || len(arguments[0]) < 1 {
		panic(fmt.Sprintf("Parameter %s is missing", argumentName))
	}

	return arguments[0]
}

// Pretty prints the data as json to the stream
func PrettyPrint(w http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent(empty, space)
	_ = encoder.Encode(data)
}

// Defers the closing of the given closer
func CheckClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Fatal("Error during close")
	}
}
