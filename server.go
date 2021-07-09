package main

import (
	//"flag"
	"fmt"
	"net/http"

	//"net/http/httptest"
	"os"
	"time"
	//"cache_example/cache"
	//"cache_example/cache/memory"
)

//var storage cache.Storage

func indexo(w http.ResponseWriter, r *http.Request) {

	time.Sleep(2 * time.Second)

	content := fmt.Sprintf(`
		<h1>Hello World! You are on: %s</h1> 
		<p>Current time is: %s</p>

		<ul>
			<li><a href="/">Home</a></li>
			<li><a href="/?page=1">Page 1</a></li>
			<li><a href="/?page=2">Page 2</a></li>
			<li><a href="/about">About (not cached!)</a></li>
		</ul>
	`, r.RequestURI, time.Now())
	w.Write([]byte(content))

}

func about(w http.ResponseWriter, r *http.Request) {

	time.Sleep(2 * time.Second)

	content := fmt.Sprintf(`
		<h1>About!</h1> 
		<p>Current time is: %s</p>

		<ul>
			<li><a href="/">Home</a></li>
			<li><a href="/about">About (not cached!)</a></li>
		</ul>
	`, time.Now())
	w.Write([]byte(content))
}

func main() {

	http.HandleFunc("/", indexo)
	http.HandleFunc("/about", about)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is up and listening on port %s.\n", port)
	http.ListenAndServe(":"+port, nil)
}
