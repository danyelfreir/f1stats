package repositories

import (
	"danyelfreir/f1stats/models"
	"fmt"

	"gorm.io/gorm"
)

type DriversList struct {
	Drivers []models.Driver
}

type SeasonList struct {
	Seasons []models.Season
}

type DriverRepository struct {
	DB *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return DriverRepository{db}
}

func (r DriverRepository) GetAll() DriversList {
	var drivers []models.Driver
	result := r.DB.Find(&drivers)
	fmt.Println(result.RowsAffected)
	return DriversList{drivers}
}

func (r DriverRepository) GetYears() SeasonList {
	var seasons []models.Season
	result := r.DB.Order("year DESC").Find(&seasons)
	fmt.Println(result.RowsAffected)
	return SeasonList{seasons}
}

func (r DriverRepository) GetDriversFromYear(year int) DriversList {
	var drivers []models.Driver
	result := r.DB.Joins("JOIN driver_season ON drivers.driverid = driver_season.driverid").Where("driver_season.year = ?", year).Find(&drivers)
	fmt.Println(result.RowsAffected)
	return DriversList{drivers}
}
