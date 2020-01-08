package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/theothertomelliott/emojicode-playground/pkg/run"
	"github.com/theothertomelliott/emojicode-playground/pkg/run/binaryexec"
)

func main() {
	workingDir := os.Args[1]

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			fmt.Printf("Request completed in %v\n", time.Since(start))
		}()

		runner := run.New(binaryexec.New(), path.Join(workingDir, "testdata"), 2*time.Second)

		code, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Returning error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(code) == 0 {
			http.Error(w, "no code submitted", http.StatusBadRequest)
			return
		}

		err = runner.Run(context.Background(), code, w)
		if err != nil {
			fmt.Printf("Returning error: %v", err)
			// TODO: Ensure this adds an appropriate status code
			// Currently results in a "http: superfluous response.WriteHeader call" warning
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
