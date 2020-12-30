package entity

// Movie entity
type Movie struct {
	ID          int      `json:"id,string"`
	Title       string   `json:"title"`
	Poster      string   `json:"poster"`
	Overview    string   `json:"overview"`
	ReleaseDate int      `json:"release_date"`
	Genres      []string `json:"genres"`
}
