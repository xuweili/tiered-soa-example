package main

import (
	"fmt"
	"github.com/deferpanic/deferclient/deferstats"
	"io/ioutil"
	"net/http"
)

type blah struct {
	Stuff string
}

func handler(w http.ResponseWriter, r *http.Request) {

	// just pass your spanId w/each request
	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://127.0.0.1:3003/aphrodite", nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Header.Add("X-dpparentspanid", deferstats.GetSpanIdString(w))

	resp, err := client.Do(r)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	r, err = http.NewRequest("POST", "http://127.0.0.1:3004/dionysus", nil)
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

	http.HandleFunc("/athena", dps.HTTPHandlerFunc(handler))
	http.ListenAndServe(":3001", nil)
}
