package core

import (
	"io"
	"log"
)

// Defers the closing of the given closer
func CheckClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Fatal("Error during close")
	}
}
