package main

import (
	"SpotSync/internal/config"
	"SpotSync/internal/server"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDB(cfg)
	server.Start(db, cfg)
}
