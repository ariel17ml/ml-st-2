package main

import (
	"fmt"
	"net/http"
)

type Planet struct {
	name      string
	speed     float32 // in grades/day
	distance  int     // in km
	clockwise bool
}

// Handles the requests to the weather API. It receives a "day" parameter,
// which works as query index for data.
func weather(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", r.URL.Query().Get("day"))
}

func main() {

	// ferengi, betasoide, vulcano :=
	// 	Planet{"Ferengi", 1, 500, true},
	// 	Planet{"Betasoide", 3, 2000, true},
	// 	Planet{"Vulcano", 5, 1000, false}

	http.HandleFunc("/weather", weather)
	http.ListenAndServe(":8080", nil)
}
