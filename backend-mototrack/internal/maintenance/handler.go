package maintenance

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type MaintenanceHandler struct {
	Repository *MaintenanceRepository
	Factory    *MaintenanceFactory
}

// Register implements api.Handler.
func (handler *MaintenanceHandler) Register(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

type response struct {
	Msg string `json:"msg" bson:"msg"`
}

func NewMaintenanceHandler(db *mongo.Client) *MaintenanceHandler {
	return &MaintenanceHandler{
		Factory:    NewMaintenanceFactory(),
		Repository: NewMaintenanceRepository(db),
	}
}

func (handler *MaintenanceHandler) RegisterMaintenance(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var requestData struct {
		TypeMaintenance string  `json:"tm"`
		Date            string  `json:"date"`
		Value           float64 `json:"value"`
		Description     string  `json:"description"`
		OwnerID         string  `json:"ownerId"`
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Invalid json body: "+err.Error(), http.StatusBadRequest)
		return
	}

	maintenance := NewMaintenance(
		requestData.TypeMaintenance,
		requestData.Date,
		requestData.Description,
		requestData.OwnerID,
		requestData.Value,
	)

	err := handler.Repository.Insert(maintenance)
	if err != nil {
		log.Printf("Error inserting maintenance into database: %v", err)
		http.Error(w, "Unable to insert maintenance in the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := response{
		Msg: "Maintenance registered successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (handler *MaintenanceHandler) ListMaintenances(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId parameter", http.StatusBadRequest)
		return
	}

	maintenances, err := handler.Repository.FindByOwnerID(userId)
	if err != nil {
		log.Printf("Error listing maintenances: %v", err)
		http.Error(w, "Unable to list maintenances", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maintenances)
}

func (handler *MaintenanceHandler) DeleteMaintenances(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId parameter", http.StatusBadRequest)
		return
	}

	err := handler.Repository.DeleteByOwnerID(userId)
	if err != nil {
		log.Printf("Error deleting maintenances: %v", err)
		http.Error(w, "Unable to delete maintenances", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
