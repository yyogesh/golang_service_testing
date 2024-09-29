package main

import (
	"job_portal/internal/repository"
	"job_portal/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := gin.Default()
	routes.InitRoutes(r, db)

	r.Run(":8080")
}
