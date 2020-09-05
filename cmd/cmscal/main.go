package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/lucasbergman/cmscal"
)

func calServer(c string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar; charset=utf-8")
		fmt.Fprint(w, c)
	}
}

func main() {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	bs := cmscal.MakeBuildingSchedule(loc)
	calSixth := cmscal.ICalForSchedule(bs, &cmscal.ScheduleSixth)
	calSeventh := cmscal.ICalForSchedule(bs, &cmscal.ScheduleSeventh)
	calEighth := cmscal.ICalForSchedule(bs, &cmscal.ScheduleEighth)

	http.HandleFunc("/six", calServer(calSixth))
	http.HandleFunc("/6", calServer(calSixth))
	http.HandleFunc("/seven", calServer(calSeventh))
	http.HandleFunc("/7", calServer(calSeventh))
	http.HandleFunc("/eight", calServer(calEighth))
	http.HandleFunc("/8", calServer(calEighth))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://lucasbergman.github.io/cmscal/", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
