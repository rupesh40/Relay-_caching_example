package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"cache_example/cache"
	"cache_example/cache/memory"
	"io/ioutil"
	"log"
	//"mime/multipart"
)

type whoiam struct {
	Addr string
}

var storage cache.Storage

func init() {
	strategy := flag.String("s", "memory", "Cache strategy (memory or unitdb)")
	flag.Parse()

	if *strategy == "memory" {
		storage = memory.NewStorage()
	} else {
		panic(fmt.Sprintf("Invalid cache strategy %s.", *strategy))
	}
}

func main() {

	http.Handle("/", cached("10s", index))
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//riter := multipart.NewWriter(body)
	url := "http://host.docker.internal:8080"
	log.Printf("Target %s.", url)
	//r.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Content-Type", "image/jpeg; charset=utf-8")
	w.Write(body1)
	//fmt.Fprintf(w, "hi")

}

func cached(duration string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		content := storage.Get(r.RequestURI)
		if content != nil {
			fmt.Print("Cache Hit!\n")
			w.Write(content)
		} else {
			c := httptest.NewRecorder()
			handler(c, r)

			for k, v := range c.HeaderMap {
				w.Header()[k] = v
			}

			w.WriteHeader(c.Code)
			content := c.Body.Bytes()

			if d, err := time.ParseDuration(duration); err == nil {
				fmt.Printf("New page cached: %s for %s\n", r.RequestURI, duration)
				storage.Set(r.RequestURI, content, d)
			} else {
				fmt.Printf("Page not cached. err: %s\n", err)
			}

			w.Write(content)
		}

	})
}
