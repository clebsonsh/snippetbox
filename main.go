package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice
// containing "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches
	// "/". If it doesn't, use the http.NotFound() function
	// to send a 404 respponse to the client. Importantly,
	// we then return form the handler. If we don't return
	// the handler would keep executing and also write the
	// "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the resquest is using Post or not
	if r.Method != http.MethodPost {
		// If it's not, use the Header().Set() method to add an 'Allow: POST' headar
		// to the response header map. The first parameter is the
		// header name, and the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)

		// Set a new cache-control header. If an existing "Cache-Control" header exists
		// it will be overwritten.
		w.Header().Set("Cache-Control", "public, max-age=31536000")

		// In contrast, the Add() method appends a new "Cache-Control" header and can
		// be called multiple times.
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")

		// Delete all values for the "Cache-Control" header.
		w.Header().Del("Cache-Control")

		// Retrieve the first value for the "Cache-Control" header.
		w.Header().Get("Cache-Control")

		// Retrieve a slice of all values for the "Cache-Control" header.
		w.Header().Values("Cache-Control")

		/// Suppress system-generated header
		w.Header()["Date"] = nil

		// Use the http.Error() method to send a 405 status code with
		// "Method Not Allowed" response body. We then return from
		// the function so that the subsequest code is no executed.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a
	// new servemux, then register the home function as the
	// handlesr for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Register the two new handle functions and
	// corresponding URL patterns with the servemux,
	// in exactly the same way that we did before.
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new
	// web server. We pass in two parameters: the TCP network
	// address to listen on (in this case ":4000") and the
	// servermux we just created. If http.ListenAndServe() return
	// an error we use the log.Fatal() function to log the error
	// message and exit.
	// Note that any error returned by  http.ListenAndServe()
	// is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
