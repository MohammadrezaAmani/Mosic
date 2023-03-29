package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/MohammadrezaAmani/Mosic/models"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

var musics []models.Music

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) {
	wg.Done()
	pathes := strings.Split(path, ".")
	name := pathes[len(pathes)-1]
	formats := [6]string{"mp3", "wav", "flac", "aac", "ogg", "m4a"}
	for _, f := range formats {
		if strings.ToLower(name) == f {
			log.Println(path)
		}

	}
}

func WalkDir(path string, recursive bool) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		check(err)
		if info.IsDir() && !recursive && path != "je" {
			return filepath.SkipDir
		}
		{
			if info.Mode().IsRegular() {
				wg.Add(1)
				go ReadFile(path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

// func ScanMusic(path string, recursive bool) ([]models.Music, error) {
// []string{"MP3", "WAV", "FLAC", "AAC", "OGG"}
//
//		return []models.Music{}, nil
//	}
func main() {
	WalkDir("/mnt/D/personal/Musics", true)
}
