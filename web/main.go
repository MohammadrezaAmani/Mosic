package main

import (
	"github.com/MohammadrezaAmani/Mosic/api"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbum)
	router.GET("/album/:id", api.GetAlbumByID)
	router.Run("localhost:8000")
}
