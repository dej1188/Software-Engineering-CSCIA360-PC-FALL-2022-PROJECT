package models

type Video struct {
	ID          string `gorm:"primary_key"`
	Title       string
	UserHash    string
	Description string
	YouTubeID   string
	EditKey     string
	Notations   []Notation `gorm:"foreignKey:VideoID"`
}
