package models

import "encoding/base64"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Music struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Size     string `json:"size"`
	Genre    string `json:"genre"`
	Year     string `json:"year"`
	Path     string `json:"path"`
	FileName string `json:"filename"`
	UniqueID string `json:"uniqueid"`
}

func (m Music) Unique() string {
	text := m.Name + m.Artist + m.Size + m.Path
	return base64.StdEncoding.EncodeToString([]byte(text))
}

type Setting struct {
	MusicPath []string
	Starred   []string
}

//? other structs

type Path struct {
	PathText string `json:"path"`
}
type Search struct {
	Text  string `json:"text"`
	TimeMax  string `json:"timemax"`
	TimeMin  string `json:"timemin"`
	SizeMin  string `json:"sizemin"`
	SizeMax  string `json:"sizemax"`
	Year  string `json:"year"`
	Genre string `json:"genre"`
}
