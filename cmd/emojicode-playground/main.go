package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
	"github.com/theothertomelliott/emojicode-playground/pkg/run/binaryexec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		runner := run.New(binaryexec.New(), "./testdata")

		code, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(code) == 0 {
			http.Error(w, "no code submitted", http.StatusBadRequest)
			return
		}

		err = runner.Run(context.Background(), code, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
