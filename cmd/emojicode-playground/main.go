package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
	"github.com/theothertomelliott/emojicode-playground/pkg/run/binaryexec"
)

func main() {
	workingDir := os.Args[1]

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		runner := run.New(binaryexec.New(), path.Join(workingDir, "testdata"))

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
