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
	router.GET("/rescan", api.Rescan)
	router.GET("/remove/:id", api.Remove)
	router.GET("/get/:id", api.GetMusic)
	router.POST("/path",api.AddPath)
	router.GET("/star/:id", api.AddStar)
	router.GET("/unstar/:id", api.RemoveStar)
	router.GET("/stars/", api.Stars)
	router.POST("/search",api.Search)



	utils.WalkDir()
	router.Run("0.0.0.0:8000")
}
