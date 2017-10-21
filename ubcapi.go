package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Course struct {
	Code string `json:"code"`
	Name string `json:"name,omitempty"`
}

// Global courses map
var courseMap map[string]string

func readCourses(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonParser := json.NewDecoder(fp)
	if err = jsonParser.Decode(&courseMap); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	coursesFilePath := os.Getenv("COURSES_FILE_PATH")
	readCourses(coursesFilePath)

	router := NewRouter()

	log.Println("Listening on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, cors.Default().Handler(router)))
}
