package dtos

type Video struct {
	ID          string
	Title       string
	UserHash    string
	Description string
	YouTubeID   string
	EditKey     string
	Notations   []Notation
}
