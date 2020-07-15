package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"djuta.software/record-roulette-api/controllers/channelscontroller"
	"djuta.software/record-roulette-api/controllers/videoscontroller"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/videos/random", videoscontroller.GetRandomVideos)
	router.HandleFunc("/v1/channels", channelscontroller.GetChannel)
	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	allowedOrigins := strings.Split(
		os.Getenv("ALLOWED_ORIGINS"),
		",",
	)

	c := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
	})

	log.Fatal(http.ListenAndServe(port, c.Handler(router)))
}
