package main

import (
	"github.com/eldonaldo/go-project-template/core"
	"github.com/eldonaldo/go-project-template/server"
)

// Runs the app
func main() {

	// Performs migrations, reads the config and starts the server
	core.Migrate()
	config := core.ReadConfig()
	server.StartServer(config.Port)
}
