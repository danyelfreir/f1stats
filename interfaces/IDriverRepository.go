package interfaces

import (
	"danyelfreir/f1stats/repositories"

	"gorm.io/gorm"
)

type IDriverRepository interface {
	NewDriverRepository(db *gorm.DB) IDriverRepository
	GetAll() repositories.DriversList
	GetYears() repositories.SeasonList
	GetDriversFromYear(year int) repositories.DriversList
}
