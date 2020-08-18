package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/lucasbergman/cmscal"
)

func main() {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/six", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar")
		fmt.Fprint(w, cmscal.ICalForSchedule(loc, cmscal.ScheduleSixth))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
