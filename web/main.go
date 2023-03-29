package main

import (
	"github.com/MohammadrezaAmani/Mosic/api"
	"github.com/MohammadrezaAmani/Mosic/utils"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.POST("/albums", api.PostAlbum)
	router.GET("/album/:id", api.GetAlbumByID)
	router.GET("/musics", api.GetMusics)
	router.POST("/musics", api.PostMusic)
	router.GET("/music/:id", api.GetMusicByID)
	utils.WalkDir("/mnt/D/personal/Musics", true)
	router.Run("localhost:8000")
}
