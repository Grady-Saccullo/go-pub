package api

import (
	activityStream "github.com/Grady-Saccullo/activity-pub-go/internal/activity_stream"
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
	ex := activityStream.Note{
		Object: activityStream.Object{
			Type: activityStream.ObjectTypeNote,
			ObjectCore: activityStream.ObjectCore{
				JsonLDContext: []string{
					"https://www.w3.org/ns/activitystreams",
					"https://w3id.org/security/v1",
				},
				ID: "https://hackerman-laboratories.com/v1/notes/123",
			},
		},
	}
}
