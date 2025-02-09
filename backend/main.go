package main

import (
	"database/sql"
	"container-monitoring/backend/api"
	"container-monitoring/backend/repository"
	"container-monitoring/backend/service"
	"log"
	"os"

	_ "github.com/lib/pq"

)

func main() {
	connStr := getDBConnStr()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инжекция зависимостей
	pingRepo := repository.NewPostgresPingRepository(db)
	pingService := service.NewPingService(pingRepo)

	router := api.NewRouter(pingService)
	router.Run(":" + getPort())
}

func getDBConnStr() string {
	if connStr := os.Getenv("POSTGRES_CONN"); connStr != "" {
		return connStr
	}
	return "host=db port=5432 user=postgres password=postgres dbname=pingdb sslmode=disable"
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}
