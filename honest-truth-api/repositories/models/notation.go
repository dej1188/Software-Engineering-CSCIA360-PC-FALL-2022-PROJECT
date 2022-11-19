package models

type Notation struct {
	ID          string `gorm:"primary_key"`
	VideoID     string
	Time        uint64
	Description string
}
