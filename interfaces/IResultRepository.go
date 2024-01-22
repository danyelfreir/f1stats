package interfaces

import "danyelfreir/f1stats/repositories"

type IResaultRepository interface {
	GetLast5Standings(driverId int) repositories.ResultList
}
