package models

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
}
type Settings struct {
}
