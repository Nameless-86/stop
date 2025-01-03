package main

import (
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

func main() {
	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server %q", html.EscapeString(r.URL.Path))
		})
		err := http.ListenAndServe(":3000", mux)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
		} else if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}

	}()
	go func() {
		resp, err := http.Get("https://google.com")
		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		str_body := string(body)
		log.Printf("%s", str_body)
	}()
	select {}
}
