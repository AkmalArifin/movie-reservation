package main

import (
	"log"

	"github.com/AkmalArifin/movie-reservation/internal/db"
	"github.com/AkmalArifin/movie-reservation/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// TODO: setting env file to absolute path
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load env file")
	}

	r := gin.Default()
	db.InitDB()

	// SEEDING
	// err = seed.Seeder()
	// if err != nil {
	// 	log.Fatal("seed not working", err.Error())
	// }

	routes.RegisterRoutes(r)
	r.Run(":8080")

	log.Println("database connected")
}
