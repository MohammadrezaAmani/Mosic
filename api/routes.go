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
