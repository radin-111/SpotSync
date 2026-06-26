package main

import (
	server "SpotSync/internal"
	"SpotSync/internal/config"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDB(cfg)
	server.Start(db, cfg)
}
