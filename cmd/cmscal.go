package main

import (
	"flag"
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

	var buildID string
	flag.StringVar(&buildID, "buildid", "", "ID for this build, used for unique event IDs")
	flag.Parse()
	if buildID == "" {
		log.Fatal("buildid flag value is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/six", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar")
		fmt.Fprint(w, cmscal.ICalForSchedule(buildID, loc, cmscal.ScheduleSixth))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
