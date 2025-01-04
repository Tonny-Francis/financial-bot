package main

import "financial-bot/config"

func main() {
	// Load Container
	context, container, err := config.LoadContainer()

	if err != nil {
		panic(err)
	}

	// HTTP Adapter
	// External HTTP Adapter Loader
	router := config.LoadRouter(context)

	// Load HTTP Server
	config.LoadHTTP(context, container, router)
}
