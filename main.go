package main

import (

	"github.com/TanaratSudjai/project-golang-api-shop/config"
	"github.com/TanaratSudjai/project-golang-api-shop/databases"
	"github.com/TanaratSudjai/project-golang-api-shop/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	server := server.NewEchoServer(conf, db.ConnectionGetting())
	server.Start()
}