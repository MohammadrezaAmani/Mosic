package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MohammadrezaAmani/Mosic/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var Musics = []models.Music{}

func AddAlbum(a models.Album) {
	Albums = append(Albums, a)
}

func GetAllAlbums() *[]models.Album {
	return &Albums
}

func GetAlbumByID(id string) (*models.Album, error) {
	for _, a := range Albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return &models.Album{}, errors.New("not found")
}
func AddMusic(m models.Music) {
	m.ID = fmt.Sprint(len(Musics))
	Musics = append(Musics, m)
}

func GetAllMusics() *[]models.Music {
	return &Musics
}

func GetMusicByID(id string) (*models.Music, error) {
	for _, a := range Musics {
		if a.ID == id {
			return &a, nil
		}
	}
	return &models.Music{}, errors.New("not found")
}
func EmptyMusic() {
	Musics = []models.Music{}
}
func RemoveMusic(id string) error {
	music,err := GetMusicByID(id)
	if err != nil{
		return err
	}
	err = os.Remove(music.Path)
	if err != nil {
		return err
	}
	return nil
}

func InitialDB() {
	db, err := gorm.Open(sqlite.Open("test.db"),&gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&models.Settings{})
	
}