package main

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"io/ioutil"
	"net/http"
	"net/url"
)

type blah struct {
	Stuff string
}

func handler(w http.ResponseWriter, r *http.Request) {

	// just pass your spanId w/each request
	resp, err := http.PostForm("http://127.0.0.1:3003/aphrodite",
		url.Values{"defer_parent_span_id": {deferstats.GetSpanIdString(w)},
			"blah": {"2"}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	resp, err = http.PostForm("http://127.0.0.1:3004/dionysus",
		url.Values{"defer_parent_span_id": {deferstats.GetSpanIdString(w)},
			"blah": {"2"}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))

}

func main() {

	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")

	go dps.CaptureStats()

	http.HandleFunc("/athena", dps.HTTPHandlerFunc(handler))
	http.ListenAndServe(":3001", nil)
}
