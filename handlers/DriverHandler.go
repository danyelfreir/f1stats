package handlers

import (
	"danyelfreir/f1stats/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DriverHandler struct {
	repository repositories.DriverRepository
}

func NewDriverHandler(repository repositories.DriverRepository) DriverHandler {
	return DriverHandler{repository}
}

func (h *DriverHandler) getYears(w http.ResponseWriter) {
	seasons := h.repository.GetYears()
	res, err := json.Marshal(&seasons)
	if err != nil {
		fmt.Printf("Error while marshalling YEARS\n%s\n", err)
		http.Error(w, "Error while marshalling YEARS", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *DriverHandler) getDrivers(w http.ResponseWriter, year int) {
	// Returns an array of all drivers that participated in a given year/season
	drivers := h.repository.GetDriversFromYear(year)
	res, err := json.Marshal(&drivers)
	if err != nil {
		fmt.Printf("Error while marshalling data\n%s\n", err)
		http.Error(w, "Error while marshalling data", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *DriverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s request to %s from %s\n", r.Method, r.RequestURI, r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(r.Form) > 0 {
		if r.Form.Has("year") {
			year, err := strconv.Atoi(r.Form["year"][0])
			if err != nil {
				log.Fatal(err.Error())
			}
			h.getDrivers(w, year)
		}
	} else {
		h.getYears(w)
	}

}
