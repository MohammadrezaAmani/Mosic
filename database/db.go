package database

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	// "time"

	// "strconv"
	"strings"

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

var settings = models.Setting{}

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
	music, err := GetMusicByID(id)
	if err != nil {
		return err
	}
	err = os.Remove(music.Path)
	if err != nil {
		return err
	}
	return nil
}

func InitialDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&models.Setting{})
}

func AddPath(path string) error {
	for _, p := range settings.MusicPath {
		if p == path {
			return errors.New("path already exist")
		}
	}
	settings.MusicPath = append(settings.MusicPath, path)
	return nil
}

func GetPath() []string {
	var p []string
	p = append(p, settings.MusicPath...)
	return p
}

func AddStar(id string) {
	music, err := GetMusicByID(id)
	if err != nil {
		return
	}
	for _, m := range settings.Starred {
		if m == music.UniqueID {
			return
		}
	}
	settings.Starred = append(settings.Starred, music.UniqueID)
}
func RemoveStar(id string) {
	music, err := GetMusicByID(id)
	if err != nil {
		return
	}
	for i, m := range settings.Starred {
		if m == music.UniqueID {
			settings.Starred = append(settings.Starred[:i], settings.Starred[i+1:]...)
			return
		}
	}
}

func Starred() []models.Music {
	st := []models.Music{}

	for _, id := range settings.Starred {
		for _, music := range Musics {
			if id == music.UniqueID {
				st = append(st, music)
			}
		}
	}
	return st
}

func Search(s models.Search) []models.Music {
	music := []models.Music{}
	for _,m := range Musics {
		// if s.Time != 0 {
		// 	if 
		// }
		if strings.Contains(strings.ToLower(m.Album+" "+m.Artist+" "+m.Name), strings.ToLower(s.Text)) {
			music = append(music, m)
		}
	}
	return music
}
func Shuffle() []models.Music {
	m := Musics
	rand.Shuffle(len(m), func(i, j int) { m[i], m[j] = m[j], m[i] })
	return m
}