package main

import (
	"encoding/json"
	"github.com/deferpanic/deferclient/deferstats"
	"net/http"
	"time"
)

type blah struct {
	Stuff string
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(300 * time.Millisecond)

	stuff := blah{
		Stuff: "reply from aphrodite",
	}

	js, err := json.Marshal(stuff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {

	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")

	go dps.CaptureStats()

	http.HandleFunc("/aphrodite", dps.HTTPHandlerFunc(handler))
	http.ListenAndServe(":3003", nil)
}
