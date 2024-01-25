package main

import (
	"github.com/Monteiro712/go-gin-api/db"
	"github.com/Monteiro712/go-gin-api/routes"
)

func main() {
	db.ConectaComBancoDeDados()
	routes.HandlerRequests()
}