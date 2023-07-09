package api

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/.well-known/webfinger", webFingerHandler)
	http.HandleFunc("/v1/inbox", inboxHandler)
	http.HandleFunc("/v1/outbox", outboxHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func webFingerHandler(w http.ResponseWriter, r *http.Request) {
}

func inboxHandler(w http.ResponseWriter, r *http.Request) {

}

func outboxHandler(w http.ResponseWriter, r *http.Request) {

}
