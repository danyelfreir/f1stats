package models

type Result struct {
	Resultid        int32 `gorm:"primaryKey"`
	Raceid          int32
	Driverid        int
	Constructorid   int32
	Number          int32
	Grid            int32
	Position        int32
	Positiontext    string
	Positionorder   int32
	Points          int32
	Laps            int32
	Time            string
	Milliseconds    int32
	Fastestlap      string
	Rank            string
	Fastestlaptime  string
	Fastestlapspeed float32
	Statusid        int32
}
