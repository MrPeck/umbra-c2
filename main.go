package main

import "umbra-c2/api"

func main() {
	apiConfig := &api.APIConfig{
		Host: "0.0.0.0",
		Port: "8080",
	}

	api.Run(apiConfig)
}
