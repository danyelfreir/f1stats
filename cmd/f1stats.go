package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"

	d "danyelfreir/f1stats/internal"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
		port = "8000"
	}
	connectionString := os.Getenv("CONN_STR")
	mux := http.NewServeMux()

	// To log files to json
	// logfile, err := os.OpenFile("logfile.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// logger := slog.New(slog.NewJSONHandler(logfile, nil))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := d.OpenDB(connectionString)
	// err = db.Ping()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	repository := d.NewPostgresRepository(db, logger)
	service := d.NewService(logger, repository)

	templHandler := d.NewTemplateHandler(
		template.Must(template.ParseGlob("frontend/*.html")),
		logger,
		service,
	)
	apiHandler := d.NewApiHandler(logger, service)

	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.Handle("GET /circuits/{year}", templHandler.HandleCircuits())
	mux.Handle("GET /races/{raceId}/drivers", templHandler.HandleDrivers())
	mux.Handle("GET /races/{raceId}/drivers/{driverId}", templHandler.HandleLapsPits())
	mux.Handle("GET /", templHandler.HandleSeasons())
	mux.Handle("GET /api/circuits/{year}", apiHandler.HandleCircuits())
	mux.Handle("GET /api/races/{raceId}/drivers", apiHandler.HandleDrivers())
	mux.Handle("GET /api/races/{raceId}/drivers/{driverId}", apiHandler.HandleLapsPits())
	mux.Handle("GET /api/", apiHandler.HandleSeasons())

	fmt.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
