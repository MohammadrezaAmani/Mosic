package api

import (
	"net/http"

	"github.com/MohammadrezaAmani/Mosic/database"
	"github.com/MohammadrezaAmani/Mosic/models"
	"github.com/gin-gonic/gin"
)


func PostAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	go database.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := database.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,album)
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, *database.GetAllAlbums())
}

func GetMusicByID(c *gin.Context) {
	id := c.Param("id")
	music, err := database.GetMusicByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "music not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,music)
}

func GetMusics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, *database.GetAllMusics())
}
func PostMusic(c *gin.Context) {
	var newMusic models.Music
	if err := c.BindJSON(&newMusic); err != nil {
		return
	}
	go database.AddMusic(newMusic)
	c.IndentedJSON(http.StatusOK, newMusic)
}