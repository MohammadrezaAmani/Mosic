package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type Music struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
	Path   string `json:"path"`
}
type Settings struct {

}
