package cloudFunction

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var firestoreClient *firestore.Client

type CounterResponse struct {
	Count int64 `json:"count"`
}

func init() {
	ctx := context.Background()
	projectID := "cis437-hw4-476803"
	databaseID := "visitor-count-db"

	var err error
	firestoreClient, err = firestore.NewClientWithDatabase(ctx, projectID, databaseID)
	if err != nil {
		log.Fatalf("firestore init error: %v", err)
	}

	functions.HTTP("VisitorCounter", visitorCounter)
}

func visitorCounter(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ctx := r.Context()
	doc := firestoreClient.Collection("counts").Doc("visitor_count")

	firestoreSnapshot, err := doc.Get(ctx)
	if err != nil {
		respondError(w, err)
		return
	}

	firestoreCount, err := firestoreSnapshot.DataAt("count")
	if err != nil {
		respondError(w, err)
		return
	}

	count, ok := firestoreCount.(int64)
	if !ok {
		log.Print("Count cast to int64 failed.")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	newCount := count + 1

	if _, err := doc.Set(ctx, map[string]any{"count": newCount}, firestore.MergeAll); err != nil {
		respondError(w, err)
		return
	}

	json.NewEncoder(w).Encode(CounterResponse{Count: newCount})
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://storage.googleapis.com")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func respondError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
