package main

import (
	"golang-api-server-template/configs"
	"golang-api-server-template/server"

	_ "golang-api-server-template/docs"
)

func main() {
	configs.Get()
	configs.ConnectDB()
	server.ServerStart()
}
