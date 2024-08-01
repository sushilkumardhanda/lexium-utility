package main

import (
	"lexium-utility/config"
	"lexium-utility/middlewares"
	"lexium-utility/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.GetMongoClient()
	config.GetRedisClient()
}
func main() {
	//seeding.Seed()
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	public := r.Group("/api")

	routes.RoutesPublic(public)

	protected := r.Group("api/common")
	routes.RoutesProtected(protected)

	r.Run(":8080")
}
