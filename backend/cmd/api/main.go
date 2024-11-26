package main

import (
	handler "backend/cmd/api/handlers"
	middleware "backend/cmd/api/middleware"
	cardsHandler "backend/internal/cards/handler"
	cardsRepo "backend/internal/cards/repository"
	cardsService "backend/internal/cards/service"
	collectioncardsHandler "backend/internal/collectioncards/handler"
	collectioncardsRepo "backend/internal/collectioncards/repository"
	collectioncardsService "backend/internal/collectioncards/service"
	collectionsHandler "backend/internal/collections/handler"
	collectionsRepo "backend/internal/collections/repository"
	collectionsService "backend/internal/collections/service"
	usersHandler "backend/internal/users/handler"
	usersRepo "backend/internal/users/repository"
	usersService "backend/internal/users/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := handler.New()
	googleHandler := h.NewGoogleHandler(h.Config)

	// Init repositories
	cardsRepo := cardsRepo.NewCardsRepository(h.DB)
	collectionsRepo := collectionsRepo.NewCardsRepository(h.DB)
	collectioncardsRepo := collectioncardsRepo.NewCollectionCardsRepository(h.DB)
	usersRepo := usersRepo.NewUsersRepository(h.DB)

	// Init services
	cardsService := cardsService.NewCardsService(cardsRepo)
	collectionsService := collectionsService.NewCollectionsService(collectionsRepo)
	collectioncardsService := collectioncardsService.NewCollectionCardsService(collectioncardsRepo)
	usersService := usersService.NewUsersService(usersRepo)

	// Init handlers
	cardsHandler := cardsHandler.NewCardsHandler(cardsService)
	collectioncardsHandler := collectioncardsHandler.NewCollectionCardsHandler(collectioncardsService)
	collectionsHandler := collectionsHandler.NewCollectionsHandler(collectionsService)
	usersHandler := usersHandler.NewUsersHandler(usersService)

	router := chi.NewRouter()

	// Sample
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// Authorization
	router.Route("/auth", func(router chi.Router) {
		router.Get("/login/google", googleHandler.GoogleLoginHandler)
		router.Get("/callback/google", googleHandler.GoogleCallbackHandler)
	})

	router.Group(func(r chi.Router) {
		router.Use(middleware.VerifyJWTToken)

		// Cards
		router.Route("/cards", func(router chi.Router) {
			router.With(h.Pagination).Get("/", cardsHandler.Search)
			router.Get("/{cardID:[0-9]+}", cardsHandler.GetByID)
		})

		// Collections
		router.Route("/collections", func(r chi.Router) {
			router.Route("/", func(r chi.Router) {
				router.Post("", collectionsHandler.Create)
				router.Put("", collectionsHandler.Update)
			})
			router.Route("/{collectionID:[0-9]+}", func(r chi.Router) {
				router.Get("/", collectionsHandler.GetByID)
				router.Delete("/", collectionsHandler.GetByID)

				router.Route("/collection-cards", func(r chi.Router) {
					router.Post("/", collectioncardsHandler.AddCardsToCollection)
					router.Get("/", collectioncardsHandler.GetByCollectionID)
					router.Delete("/", collectioncardsHandler.RemoveCardsFromCollection)
				})
			})
		})

		// Users
		router.Route("/users", func(r chi.Router) {
			router.With(h.Pagination).Get("/{userID:[0-9]+}/collections", collectionsHandler.GetAllCollectionsByUserID)
			router.Post("/", usersHandler.Create)
		})
	})
	// Collection Cards Handler

	// User Handler

	// router := mux.NewRouter()
	// // Authorization
	// router.HandleFunc("/auth/login/google", googleHandler.GoogleLoginHandler).Methods("GET")
	// router.HandleFunc("/auth/callback/google", googleHandler.GoogleCallbackHandler).Methods("GET")
	// // Cards Handler
	// router.HandleFunc("/cards", middleware.VerifyJWTToken(cardsHandler.Search)).Methods("GET")
	// // Collections Handler
	// router.HandleFunc("/collections/{collectionId:[0-9]+}", h.GetCollection).Methods("GET")
	// router.HandleFunc("/collections", h.CreateCollection).Methods("POST")
	// router.HandleFunc("/collections/{collectionId:[0-9]+}", h.UpdateCollection).Methods("PUT")
	// router.HandleFunc("/collections/{collectionId:[0-9]+}", h.DeleteCollection).Methods("DELETE")
	// // Collection Cards Handler
	// router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.GetCollectionCards).Methods("GET")
	// router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.AddCardsToCollection).Methods("PUT")
	// router.HandleFunc("/collections/{collectionId:[0-9]+}/collection-cards", h.RemoveCardsFromCollection).Methods("DELETE")
	// // Users Handler
	// router.HandleFunc("/users/{userId:[0-9]+}", h.GetUser).Methods("GET")
	// // router.HandleFunc("/users", h.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{userId:[0-9]+}/collections", h.GetUserCollections).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
