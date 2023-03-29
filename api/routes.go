package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MohammadrezaAmani/Mosic/database"
	"github.com/MohammadrezaAmani/Mosic/models"
	"github.com/MohammadrezaAmani/Mosic/utils"
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
	c.IndentedJSON(http.StatusOK, album)
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
	c.IndentedJSON(http.StatusOK, music)
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
func Rescan(c *gin.Context) {
	database.EmptyMusic()
	utils.WalkDir()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "scanned successfully"})
}

func Remove(c *gin.Context) {
	err:= database.RemoveMusic(c.Param("id"))
	if err!=nil {
		c.IndentedJSON(http.StatusBadRequest,"err")
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"music remove sucessfully"})
}
func GetMusic(c *gin.Context) {
	id := c.Param("id")
	music,err := database.GetMusicByID(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound,"not found")
		return

	}
	content, err2 := os.ReadFile(music.Path)
	if err2 != nil {
		log.Println(err2)
		c.IndentedJSON(http.StatusNotFound,"not found")
		return
	}
	log.Printf(music.FileName)
	c.Header("Content-Type","attachment; filename="+ music.FileName)
	c.Header("Accept-Length",fmt.Sprintf("%d",len(content)))
	c.Writer.Write(content)
}