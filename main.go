package main 

import (
	"log"
	"net/http"
)


func main() {
	const filepathRoot = "."
	const port = "8080"

	fileServer := http.FileServer(http.Dir(filepathRoot))

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))
	mux.Handle("/assets", http.FileServer(http.Dir(filepathRoot + "/assets")))
	mux.HandleFunc("/healthz", readinessHandler)

	server := &http.Server {
		Handler: mux,  
		Addr: ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	// set the Content-type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Write the status code
	w.WriteHeader(http.StatusOK)
	// Write "OK" to the response body
	w.Write([]byte("OK"))
}