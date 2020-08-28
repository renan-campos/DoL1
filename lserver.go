/* Hosts a webserver that generates a gif of a Lissajous figure.

  - GET on <LServer_IP>:1337 returns the gif.
	- GET on <LServer_IP>:1337/Params returns a JSON of the parameters.
	- PUT on <LServer_IP>:1337/Params modifies the curve's parameters.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	lissajous "github.com/renan-campos/DoL1/figures"
)

var params lissajous.Parameters

func main() {
	lissajous.CheckFixParams(&params)

	http.HandleFunc("/", drawer)
	http.HandleFunc("/Params", pHandler)
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func drawer(w http.ResponseWriter, r *http.Request) {
	lissajous.DrawGIF(w, params)
}

func pHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		// We must close r.Body on all execution paths.
		var newParams lissajous.Parameters
		if err := json.NewDecoder(r.Body).Decode(&newParams); err != nil {
			r.Body.Close()
			return
		}
		lissajous.CheckFixParams(&newParams)
		params = newParams
	}
	data, err := json.MarshalIndent(params, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Fprintf(w, "%s\n", data)
}
