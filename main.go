package main

import (
	"fmt"
	"github.com/RianIhsan/go-chi-crud/config"
	"github.com/RianIhsan/go-chi-crud/db"
	"github.com/RianIhsan/go-chi-crud/feature/movie/repository"
	"github.com/RianIhsan/go-chi-crud/feature/movie/service"
	"github.com/RianIhsan/go-chi-crud/routes"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	dbConn, err := db.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	movieRepository := repository.NewMovieRepository(dbConn)
	movieService := service.NewMovieService(movieRepository)
	router := chi.NewRouter()
	routes.MovieRoutes(router, movieService)

	appConfig := config.GetConfig()
	port := appConfig.AppPort
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
