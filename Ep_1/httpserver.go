package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
// 		log.Println("Hello World")

// 		// Read request body
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(rw, "Ooops!", http.StatusBadRequest)
// 			return
// 		}
// 		defer r.Body.Close()

// 		fmt.Fprintf(rw, "Hello %s", body)
// 	})

// 	mux.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
// 		log.Println("Goodbye World")
// 		rw.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(rw, "Goodbye!")
// 	})

// 	// Start server
// 	log.Println("Starting server on :9090")
// 	if err := http.ListenAndServe(":9090", mux); err != nil {
// 		log.Fatalf("Server failed: %v", err)
// 	}
// }
