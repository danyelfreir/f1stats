package main

import (
	"danyelfreir/f1stats/db"
	"danyelfreir/f1stats/handlers"
	"danyelfreir/f1stats/repositories"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	mux := http.NewServeMux()

	sqlDB, err := db.OpenDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	DB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	driverHandler := handlers.NewDriverHandler(repositories.NewDriverRepository(DB))
	resultHandler := handlers.NewResultHandler(repositories.NewResultRepository(DB))
	fs := http.FileServer(http.Dir("./frontend"))

	mux.Handle("/drivers/", &driverHandler)
	mux.Handle("/results/", &resultHandler)
	mux.Handle("/", fs)

	fmt.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
