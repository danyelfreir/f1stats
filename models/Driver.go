package models

type Driver struct {
	Driverid    int `gorm:"primaryKey"`
	Driverref   string
	Number      int
	Code        string
	Forename    string
	Surname     string
	Dob         string
	Nationality string
	URL         string
}
