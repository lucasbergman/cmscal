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

	bs := cmscal.MakeBuildingSchedule(loc)
	calSixth := cmscal.ICalForSchedule(bs, &cmscal.ScheduleSixth)
	calSeventh := cmscal.ICalForSchedule(bs, &cmscal.ScheduleSeventh)
	calEighth := cmscal.ICalForSchedule(bs, &cmscal.ScheduleEighth)
	http.HandleFunc("/six", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar; charset=utf-8")
		fmt.Fprint(w, calSixth)
	})
	http.HandleFunc("/seven", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar; charset=utf-8")
		fmt.Fprint(w, calSeventh)
	})
	http.HandleFunc("/eight", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar; charset=utf-8")
		fmt.Fprint(w, calEighth)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
