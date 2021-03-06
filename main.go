package main

import (
	"github.com/AkatsaZXC/web-service-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", routes.GetAlbums)
	router.POST("/albums", routes.PostAlbums)
	router.GET("/albums/:id", routes.GetAlbumByID)
	router.Run("localhost:8080")
}
