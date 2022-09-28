package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/single/:id", getSingleAlbum)
	router.POST("/albums/create", createAlbum)

	router.Run("localhost:8081")

}
