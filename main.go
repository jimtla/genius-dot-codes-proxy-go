package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const head = "<head>"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upstreamUrl := "http:/" + r.URL.Path
		resp, err := http.Get(upstreamUrl)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		res := strings.Replace(string(buf), "<head>",
			fmt.Sprintf(`<head><base href="%s"><script async src="//genius.codes"></script>`,
				upstreamUrl), 1)

		w.Write([]byte(res))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
