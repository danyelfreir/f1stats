package models

type Season struct {
	Year int `gorm:"primaryKey"`
	Url  string
}
