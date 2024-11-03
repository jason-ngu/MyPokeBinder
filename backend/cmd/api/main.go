package main

import (
	handler "backend/cmd/api/handlers"
	middleware "backend/cmd/api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	h := handler.New()

	googleAuth := h.NewGoogleAuth()

	router := mux.NewRouter()
	// Authorization
	router.HandleFunc("/auth/login/google", googleAuth.GoogleLoginHandler).Methods("GET")
	router.HandleFunc("/auth/callback/google", googleAuth.GoogleCallbackHandler).Methods("GET")
	// Cards Handler
	router.HandleFunc("/cards", middleware.VerifyJWTToken(h.SearchCards)).Methods("GET")
	// Collections Handler
	router.HandleFunc("/collections/{collectionId:[0-9]+}", h.GetCollection).Methods("GET")
	router.HandleFunc("/collections", h.CreateCollection).Methods("POST")
	router.HandleFunc("/collections/{collectionId:[0-9]+}", h.UpdateCollection).Methods("PUT")
	router.HandleFunc("/collections/{collectionId:[0-9]+}", h.DeleteCollection).Methods("DELETE")
	// Collection Cards Handler
	router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.GetCollectionCards).Methods("GET")
	router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.AddCardsToCollection).Methods("PUT")
	router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.RemoveCardsFromCollection).Methods("DELETE")
	// Users Handler
	router.HandleFunc("/users/{userId:[0-9]+}", h.GetUser).Methods("GET")
	// router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{userId:[0-9]+}/collections", h.GetUserCollections).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
