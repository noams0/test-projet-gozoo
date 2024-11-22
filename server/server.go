package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"zoo-panic/simulation"
)

func NewServerRestAgent(addr string, simulation *simulation.SimulationMA) *ServerRestAgent {
	return &ServerRestAgent{
		addr:       addr,
		simulation: simulation,
	}
}

// Middleware pour gérer les en-têtes CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si la méthode est OPTIONS, on retourne immédiatement une réponse OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *ServerRestAgent) handleGetAnimals(w http.ResponseWriter, r *http.Request) {
	animals := s.simulation.Env.GetAnimals()
	var animalDTOs []AnimalDTO
	for _, animal := range animals {
		animalDTOs = append(animalDTOs, ToAnimalDTO(animal))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animalDTOs)
}

func (s *ServerRestAgent) handleGetVisitors(w http.ResponseWriter, r *http.Request) {
	visitors := s.simulation.Env.GetVisitors()
	log.Println(visitors)
	// Transformer les visiteurs en DTO
	var visitorDTOs []VisitorDTO
	for _, visitor := range visitors {
		visitorDTOs = append(visitorDTOs, ToVisitorDTO(*visitor))
	}

	// Retourner en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visitorDTOs)
}

func (s *ServerRestAgent) handleOpenCage(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CageID int `json:"cage_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	s.simulation.Env.OpenCage(request.CageID)
	w.WriteHeader(http.StatusOK)
}

func (s *ServerRestAgent) UpdateAnimalEscape(w http.ResponseWriter, r *http.Request) {
	// Parse la requête JSON
	var input struct {
		ID      int  `json:"id"`
		Escaped bool `json:"escaped"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Logique pour trouver et mettre à jour l'animal
	for _, animal := range s.simulation.Env.GetAnimals() {
		if animal.Id == input.ID {
			animal.Escape(s.simulation.Env)
			fmt.Printf("Animal %d marqué comme échappé: %v\n", input.ID, input.Escaped)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Animal non trouvé
	http.Error(w, "Animal not found", http.StatusNotFound)
}

func (rsa *ServerRestAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/api/get_animals", rsa.handleGetAnimals)
	mux.HandleFunc("/api/get_visitors", rsa.handleGetVisitors)
	mux.HandleFunc("/api/update_animal_escape", rsa.UpdateAnimalEscape)

	handlerWithCORS := enableCORS(mux)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        handlerWithCORS,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
