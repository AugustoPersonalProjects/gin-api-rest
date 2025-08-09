package main

import (
	"github.com/AugustoPersonalProjects/gin-api-rest/database"
	"github.com/AugustoPersonalProjects/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
