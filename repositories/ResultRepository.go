package repositories

import (
	"danyelfreir/f1stats/models"
	"fmt"

	"gorm.io/gorm"
)

type ResultList struct {
	Results []models.Result
}

type ResultRepository struct {
	DB *gorm.DB
}

func NewResultRepository(db *gorm.DB) ResultRepository {
	return ResultRepository{db}
}

func (r *ResultRepository) GetLast5Standings(driverId int) ResultList {
	var results []models.Result
	query := r.DB.Where("driverid = ?", driverId).Order("raceid DESC").Limit(5).Find(&results)
	fmt.Printf("%d rows affected\n", query.RowsAffected)
	return ResultList{results}
}
