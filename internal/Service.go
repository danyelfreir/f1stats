package internal

import (
	"context"
	"log/slog"
)

type Service interface {
	GetSeasons(context.Context) (SeasonList, error)
	GetCircuitsOfYear(context.Context, int) (CircuitList, error)
	GetDriversOfRace(context.Context, int) (DriverList, error)
	GetLapsAndPits(context.Context, int, int) (LapsAndPits, error)
}

type ServiceImpl struct {
	logger     *slog.Logger
	repository Repository
}

func NewService(logger *slog.Logger, repo Repository) Service {
	return ServiceImpl{logger: logger, repository: repo}
}

func (s ServiceImpl) GetSeasons(ctx context.Context) (SeasonList, error) {
	var seasons []Season
	seasons, err := s.repository.FetchAllSeasons(ctx)
	return SeasonList{Seasons: seasons}, err
}

func (s ServiceImpl) GetCircuitsOfYear(ctx context.Context, year int) (CircuitList, error) {
	var circuits []Circuit
	circuits, err := s.repository.FetchCircuitsByYear(ctx, year)
	return CircuitList{Circuits: circuits}, err
}

func (s ServiceImpl) GetDriversOfRace(ctx context.Context, raceId int) (DriverList, error) {
	var drivers []Driver
	drivers, err := s.repository.FetchDriversByRaceId(ctx, raceId)
	return DriverList{Drivers: drivers}, err
}

func (s ServiceImpl) GetLapsAndPits(ctx context.Context, raceId, driverId int) (LapsAndPits, error) {
	var laps []LapTime
	laps, err := s.repository.FetchLaptimesByRaceIdAndDriverId(ctx, raceId, driverId)
	var pits []PitStop
	// TODO
	return LapsAndPits{Laps: laps, Pits: pits}, err
}
