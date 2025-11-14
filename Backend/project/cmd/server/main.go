package main

import (
	"log"
	"net/http"
	"os"

	"beautiful-minds/backend/project/config"
	"beautiful-minds/backend/project/internal/database"
	"beautiful-minds/backend/project/internal/handlers"
	"beautiful-minds/backend/project/internal/middleware"
	"beautiful-minds/backend/project/internal/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		log.Println("Pas de fichier .env trouvé")
	}

	// Charger la configuration
	cfg := config.Load()

	// Connexion à la base de données
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Erreur connexion DB:", err)
	}
	defer db.Close()

	log.Println("✅ Connexion à PostgreSQL réussie")

	// Initialiser les repositories
	memberRepo := repository.NewMemberRepository(db)
	eventRepo := repository.NewEventRepository(db)
	announcementRepo := repository.NewAnnouncementRepository(db)

	// Initialiser les handlers
	memberHandler := handlers.NewMemberHandler(memberRepo)
	eventHandler := handlers.NewEventHandler(eventRepo)
	announcementHandler := handlers.NewAnnouncementHandler(announcementRepo)

	// Créer le routeur
	router := mux.NewRouter()

	// Middleware CORS
	router.Use(middleware.CORS)

	// Routes API
	api := router.PathPrefix("/api").Subrouter()

	// Routes membres
	api.HandleFunc("/members", memberHandler.GetAll).Methods("GET")
	api.HandleFunc("/members/search", memberHandler.Search).Methods("GET")
	api.HandleFunc("/members", memberHandler.Create).Methods("POST")
	api.HandleFunc("/members/{id}", memberHandler.GetByID).Methods("GET")
	api.HandleFunc("/members/{id}", memberHandler.Update).Methods("PUT")
	api.HandleFunc("/members/{id}", memberHandler.Delete).Methods("DELETE")

	// Routes événements
	api.HandleFunc("/events", eventHandler.GetAll).Methods("GET")
	api.HandleFunc("/events", eventHandler.Create).Methods("POST")
	api.HandleFunc("/events/{id}", eventHandler.GetByID).Methods("GET")
	api.HandleFunc("/events/{id}", eventHandler.Update).Methods("PUT")
	api.HandleFunc("/events/{id}", eventHandler.Delete).Methods("DELETE")
	api.HandleFunc("/events/{id}/register", eventHandler.RegisterMember).Methods("POST")

	// Routes annonces
	api.HandleFunc("/announcements", announcementHandler.GetAll).Methods("GET")
	api.HandleFunc("/announcements", announcementHandler.Create).Methods("POST")
	api.HandleFunc("/announcements/{id}", announcementHandler.GetByID).Methods("GET")
	api.HandleFunc("/announcements/{id}", announcementHandler.Update).Methods("PUT")
	api.HandleFunc("/announcements/{id}", announcementHandler.Delete).Methods("DELETE")

	// Démarrer le serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
