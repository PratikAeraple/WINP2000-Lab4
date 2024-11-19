package main

import (
	"net/http"
)

func main() {
	// Handle `/` route by serving files from the "templates" directory
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)

	// Start the HTTP server on PORT 8080
	port := "8080"
	println("Server running on port", port)
	panic(http.ListenAndServe(":"+port, nil))

}
