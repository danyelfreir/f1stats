package internal

type Season struct {
	Year int
	Url  string
}

type SeasonList struct {
	Seasons []Season
}

type Circuit struct {
	Circuitid int
	Raceid    int
	Name      string
	Location  string
	Country   string
}

type CircuitList struct {
	Circuits []Circuit
}

type Driver struct {
	Driverid    int
	Driverref   string
	Number      int
	Code        string
	Forename    string
	Surname     string
	Dob         string
	Nationality string
	URL         string
}

type DriverList struct {
	Drivers []Driver
}

type LapTime struct {
	Raceid       int
	Driverid     int
	Lap          int
	Position     int
	Time         string
	Milliseconds int
}

type PitStop struct {
}

type LapsAndPits struct {
	Laps    []LapTime
	Pits    []PitStop
	Average int
}
