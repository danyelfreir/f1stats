package internal

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler interface {
	HandleSeasons() http.HandlerFunc
	HandleCircuits() http.HandlerFunc
	HandleDrivers() http.HandlerFunc
	HandleLapsPits() http.HandlerFunc
}

type TemplateHandler struct {
	templates *template.Template
	logger    *slog.Logger
	service   Service
}

type ApiHandler struct {
	logger  *slog.Logger
	service Service
}

func NewTemplateHandler(templates *template.Template, logger *slog.Logger, service Service) Handler {
	return TemplateHandler{templates, logger, service}
}

func (h TemplateHandler) HandleSeasons() http.HandlerFunc {
	h.logger.Info("Calling GetYears")
	return func(w http.ResponseWriter, r *http.Request) {
		years, err := h.service.GetSeasons(r.Context())
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.WriteHeader(http.StatusOK)
		h.templates.ExecuteTemplate(w, "index", years)
	}
}

func (h TemplateHandler) HandleCircuits() http.HandlerFunc {
	h.logger.Info("Calling GetCircuits")
	return func(w http.ResponseWriter, r *http.Request) {
		year, err := strconv.Atoi(r.PathValue("year"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		circuits, err := h.service.GetCircuitsOfYear(r.Context(), year)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.WriteHeader(http.StatusOK)
		h.templates.ExecuteTemplate(w, "circuits_list", circuits)
	}
}

func (h TemplateHandler) HandleDrivers() http.HandlerFunc {
	h.logger.Info("Calling GetDrivers")
	return func(w http.ResponseWriter, r *http.Request) {
		raceId, err := strconv.Atoi(r.PathValue("raceId"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		drivers, err := h.service.GetDriversOfRace(r.Context(), raceId)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.WriteHeader(http.StatusOK)
		err = h.templates.ExecuteTemplate(w, "drivers_list", map[string]any{
			"Drivers": drivers.Drivers,
			"Raceid":  raceId,
		})
		if err != nil {
			h.logger.Error(err.Error())
		}
	}
}

func (h TemplateHandler) HandleLapsPits() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		// raceId, err := strconv.Atoi(r.PathValue("raceId"))
		// if err != nil {
		// 	h.logger.Error(err.Error())
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// driverId, err := strconv.Atoi(r.PathValue("driverId"))
		// if err != nil {
		// 	h.logger.Error(err.Error())
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// lapsPits, err := h.service.GetLapsAndPits(r.Context(), raceId, driverId)
		// if err != nil {
		// 	h.logger.Error(err.Error())
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		// w.WriteHeader(http.StatusOK)
		// h.templates.ExecuteTemplate(w, "laps_pits_chart", lapsPits)
	}
}

func NewApiHandler(logger *slog.Logger, service Service) Handler {
	return ApiHandler{logger, service}
}

func (h ApiHandler) HandleSeasons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		years, err := h.service.GetSeasons(r.Context())
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(&years)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
func (h ApiHandler) HandleCircuits() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		year, err := strconv.Atoi(r.PathValue("year"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		circuits, err := h.service.GetCircuitsOfYear(r.Context(), year)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(&circuits)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
func (h ApiHandler) HandleDrivers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raceId, err := strconv.Atoi(r.PathValue("raceId"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		drivers, err := h.service.GetDriversOfRace(r.Context(), raceId)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(&drivers)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}

func (h ApiHandler) HandleLapsPits() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raceId, err := strconv.Atoi(r.PathValue("raceId"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		driverId, err := strconv.Atoi(r.PathValue("driverId"))
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lapsPits, err := h.service.GetLapsAndPits(r.Context(), raceId, driverId)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.logger.Info(fmt.Sprintf("GET Request to %v from %v", r.RequestURI, r.RemoteAddr))
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(&lapsPits)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
}
