package handlers

import (
	"danyelfreir/f1stats/repositories"
	"danyelfreir/f1stats/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ResultHandler struct {
	repository repositories.ResultRepository
}

func NewResultHandler(repository repositories.ResultRepository) ResultHandler {
	return ResultHandler{repository}
}

func (h *ResultHandler) sendLastFiveResults(w http.ResponseWriter, driverId int) {
	results := h.repository.GetLast5Standings(driverId)
	res, err := json.Marshal(&results)
	if err != nil {
		fmt.Printf("Error while marshalling DRIVERS\n%s\n", err)
		http.Error(w, "Error while marshalling DRIVERS", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s request to %s from %s\n", r.Method, r.RequestURI, r.RemoteAddr)
	pathStubs := util.CleanAndSplitURL(r.RequestURI)
	if len(pathStubs) <= 1 {
		http.Error(w, "Incomplete URL", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.HasPrefix(pathStubs[1], "last_five") {
		if len(r.Form) > 0 {
			if r.Form.Has("driverid") {
				driverid, err := strconv.Atoi(r.Form["driverid"][0])
				if err != nil {
					http.Error(w, "Not a valid driver ID", http.StatusBadRequest)
					return
				}
				h.sendLastFiveResults(w, driverid)
			}
		} else {
			http.Error(w, "No arguments in GET call", http.StatusNotAcceptable)
		}
	}

}
