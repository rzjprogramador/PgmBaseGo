package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rzjprogramador/PgmBaseGo/controller/routes"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
