package utils

import (
	"fmt"
	"github.com/MohammadrezaAmani/Mosic/database"
	"github.com/MohammadrezaAmani/Mosic/models"
	"github.com/bogem/id3v2/v2"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func ReadFile(path string) {
	pathes := strings.Split(path, ".")
	name := pathes[len(pathes)-1]
	filenames := strings.Split(path, "/")
	filename := filenames[len(filenames)-1]
	formats := [6]string{"mp3", "wav", "flac", "aac", "ogg", "m4a"}
	for _, f := range formats {
		if strings.ToLower(name) == f {
			tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
			check(err)
			database.AddMusic(models.Music{
				Name:     tag.Title(),
				Artist:   tag.Artist(),
				Year:     tag.Year(),
				Album:    tag.Album(),
				Genre:    tag.Genre(),
				Path:     path,
				Size:     fmt.Sprint(tag.Size()),
				FileName: filename,
			})
			defer tag.Close()

		}
	}
}
func readPath(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		check(err)
		{
			if info.Mode().IsRegular() {
				go ReadFile(path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func WalkDir() {
	for _, p := range database.GetPath() {
		go readPath(p)
	}

}

// func ScanMusic(path string, recursive bool) ([]models.Music, error) {
// 		return , nil
// 	}
