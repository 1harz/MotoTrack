package motorcicles

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type MotorcicleHandler struct {
	Repository *MotorcicleRepository
	Factory    *MotorcicleFactory
}

// Register implements api.Handler.
func (handler *MotorcicleHandler) Register(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

type response struct {
	Msg string `json:"msg" bson:"msg"`
	Id  string `json:"id"  bson:"id"`
}

func NewMotorcicleHandler(db *mongo.Client) *MotorcicleHandler {
	return &MotorcicleHandler{
		Factory:    NewMotorcicleFactory(),
		Repository: NewMotorcicleRepository(db, "your-database-name", "motorcicles"),
	}
}

func (handler *MotorcicleHandler) RegisterMotorcicle(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var requestData struct {
		Make           string `json:"make"`
		Model          string `json:"model"`
		Year           int    `json:"year"`
		Plate          string `json:"plate"`
		Color          string `json:"color"`
		EngineCapacity string `json:"engineCapacity"`
		Mileage        int    `json:"mileage"`
		OwnerID        string `json:"ownerId"`
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		http.Error(w, "Invalid json body: "+err.Error(), http.StatusBadRequest)
		return
	}

	motorcicle := NewMoto(
		requestData.Make,
		requestData.Model,
		requestData.Plate,
		requestData.Color,
		requestData.EngineCapacity,
		requestData.OwnerID,
		requestData.Mileage,
		requestData.Year,
	)

	err := handler.Repository.Insert(motorcicle)
	if err != nil {
		log.Printf("Error inserting motorcicle into database: %v", err)
		http.Error(w, "Unable to insert motorcicle in the database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := response{
		Msg: "Registration successful",
		Id:  motorcicle.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (handler *MotorcicleHandler) ListMotorcicles(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId parameter", http.StatusBadRequest)
		return
	}

	motorcicles, err := handler.Repository.FindByOwnerID(userId)
	if err != nil {
		log.Printf("Error listing motorcicles: %v", err)
		http.Error(w, "Unable to list motorcicles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(motorcicles)
}

func (handler *MotorcicleHandler) DeleteMotorcicles(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId parameter", http.StatusBadRequest)
		return
	}

	err := handler.Repository.DeleteByOwnerID(userId)
	if err != nil {
		log.Printf("Error deleting motorcicles: %v", err)
		http.Error(w, "Unable to delete motorcicles", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
