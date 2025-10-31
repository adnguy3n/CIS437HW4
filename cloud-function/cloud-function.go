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

const allowedOrigin = "https://storage.googleapis.com"

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
		log.Printf("Failed to create Firestore client for database %s: %v", databaseID, err)
	}

	functions.HTTP("VisitorCounter", visitorCounter)
}

func visitorCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
	}

	ctx := r.Context()
	docRef := firestoreClient.Collection("counts").Doc("visitor_count")

	var currentCount int64 = 0
	doc, err := docRef.Get(ctx)
	if err != nil {
		log.Printf("Failed to get document: %v", err)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	countData, err := doc.DataAt("count")
	if err != nil {
		log.Printf("Failed to get current count: %v", err)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	count, castSuccess := countData.(int64)
	if castSuccess {
		currentCount = count
	} else {
		log.Print("Failed to convert count to int64.")
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	newCount := currentCount + 1

	_, err = docRef.Set(ctx, map[string]any{"count": newCount}, firestore.MergeAll)
	if err != nil {
		log.Printf("Failed to set new count: %v", err)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	response := CounterResponse{Count: newCount}
	json.NewEncoder(w).Encode(response)
}
