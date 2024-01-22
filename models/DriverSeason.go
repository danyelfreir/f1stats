package models

type DriverSeason struct {
	DriverId int `gorm:"primaryKey"`
	Year     int `gorm:"primaryKey"`
}
