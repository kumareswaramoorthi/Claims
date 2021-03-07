package main

import (
	"claims/database"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	_, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	fmt.Println("starting server...")
	r.Run() // listen and serve on 0.0.0.0:8080
}

