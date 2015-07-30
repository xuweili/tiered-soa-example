package main

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// just pass your spanId w/each request
	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://127.0.0.1:3001/athena", nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("X-dpparentspanid", deferstats.GetSpanIdString(w))

	resp, err := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	r, err = http.NewRequest("POST", "http://127.0.0.1:3002/hercules", nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("X-dpparentspanid", deferstats.GetSpanIdString(w))

	resp, err = client.Do(r)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, string(body))
}

func main() {

	dps := deferstats.NewClient("v00L0K6CdKjE4QwX5DL1iiODxovAHUfo")

	go dps.CaptureStats()

	http.HandleFunc("/zeus", dps.HTTPHandlerFunc(handler))
	http.ListenAndServe(":3000", nil)
}
