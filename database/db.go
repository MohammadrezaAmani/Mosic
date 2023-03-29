package database

import (
	"errors"

	"github.com/MohammadrezaAmani/Mosic/models"
)

var Albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

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