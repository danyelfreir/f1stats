package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Repository interface {
	FetchAllSeasons(context.Context) ([]Season, error)
	FetchCircuitsByYear(context.Context, int) ([]Circuit, error)
	FetchDriversByRaceId(context.Context, int) ([]Driver, error)
	FetchLaptimesByRaceIdAndDriverId(context.Context, int, int) ([]LapTime, error)
	FetchPitstopsByRaceIdAndDriverId(context.Context, int, int) ([]PitStop, error)
}

type PostgresRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewPostgresRepository(db *sql.DB, logger *slog.Logger) Repository {
	return PostgresRepository{db: db, logger: logger}
}

func (r PostgresRepository) FetchAllSeasons(ctx context.Context) ([]Season, error) {
	q := "SELECT * FROM seasons s ORDER BY s.year DESC;"
	var seasonList []Season
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		r.logger.Error(err.Error())
		return seasonList, err
	}
	r.logger.Info(q)
	for rows.Next() {
		var s Season
		err = rows.Scan(
			&s.Year,
			&s.Url,
		)
		if err != nil {
			r.logger.Error(err.Error())
		} else {
			seasonList = append(seasonList, s)
		}
	}
	r.logger.Info(fmt.Sprintf("%d rows fetched successfully", len(seasonList)))
	return seasonList, nil
}

func (r PostgresRepository) FetchCircuitsByYear(ctx context.Context, year int) ([]Circuit, error) {
	q := "SELECT c.circuitid, r.raceid, c.name, c.location, c.country FROM circuits c JOIN races r ON r.circuitid = c.circuitid WHERE r.year = $1 ORDER BY r.round ASC;"
	r.logger.Info(q)
	var circuitList []Circuit
	rows, err := r.db.QueryContext(ctx, q, year)
	if err != nil {
		r.logger.Error(err.Error())
		return circuitList, err
	}
	for rows.Next() {
		var c Circuit
		err = rows.Scan(
			&c.Circuitid,
			&c.Raceid,
			&c.Name,
			&c.Location,
			&c.Country,
		)
		if err != nil {
			r.logger.Error(err.Error())
		} else {
			circuitList = append(circuitList, c)
		}
	}
	r.logger.Info(fmt.Sprintf("%d rows fetched successfully\n", len(circuitList)))
	return circuitList, nil
}

func (r PostgresRepository) FetchDriversByRaceId(ctx context.Context, raceId int) ([]Driver, error) {
	q := "SELECT d.driverid, d.driverref, d.number, d.code, d.forename, d.surname, d.dob, d.nationality, d.url FROM drivers d JOIN results re ON d.driverid = re.driverid JOIN races ra ON re.raceid = ra.raceid WHERE ra.raceid = $1 ORDER BY re.statusid, re.position ASC;"
	r.logger.Info(q)
	var driverList []Driver
	rows, err := r.db.QueryContext(ctx, q, raceId)
	if err != nil {
		r.logger.Error(err.Error())
		return driverList, err
	}
	for rows.Next() {
		var d Driver
		err = rows.Scan(
			&d.Driverid,
			&d.Driverref,
			&d.Number,
			&d.Code,
			&d.Forename,
			&d.Surname,
			&d.Dob,
			&d.Nationality,
			&d.URL,
		)
		if err != nil {
			r.logger.Error(err.Error())
		} else {
			driverList = append(driverList, d)
		}
	}
	r.logger.Info(fmt.Sprintf("%d rows fetched successfully\n", len(driverList)))
	return driverList, nil
}

func (r PostgresRepository) FetchLaptimesByRaceIdAndDriverId(ctx context.Context, raceId int, driverId int) ([]LapTime, error) {
	q := "SELECT l.raceid, l.driverid, l.lap, l.position, l.time, l.milliseconds FROM lap_times l WHERE l.raceid = $1 AND l.driverid = $2 ORDER BY l.lap ASC;"
	r.logger.Info("Executing SQL", "Query: ", q)
	var lapTimes []LapTime
	rows, err := r.db.QueryContext(ctx, q, raceId, driverId)
	if err != nil {
		r.logger.Error(err.Error())
	} else {
		for rows.Next() {
			var l LapTime
			err = rows.Scan(
				&l.Raceid,
				&l.Driverid,
				&l.Lap,
				&l.Position,
				&l.Time,
				&l.Milliseconds,
			)
			if err != nil {
				r.logger.Error(err.Error())
			} else {
				lapTimes = append(lapTimes, l)
			}
		}
	}
	r.logger.Info(fmt.Sprintf("%d rows fetched successfully\n", len(lapTimes)))
	return lapTimes, err
}

func (r PostgresRepository) FetchPitstopsByRaceIdAndDriverId(ctx context.Context, raceId int, driverId int) ([]PitStop, error) {
	var laps []PitStop
	var err error
	return laps, err
}
