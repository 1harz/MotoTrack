package maintenance

import (
	"log"
	"net/http"

	"gitlab.com/prof-caio-de-melo/user-service/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type MaintenanceService struct {
	Router   *api.Router
	Database *mongo.Client
}

func NewMaintenanceService(db *mongo.Client) *MaintenanceService {

	handler := NewMaintenanceHandler(db)

	paths := map[api.Route]api.HandlerFunc{
		{Method: "POST", Path: "/maintenance/register"}: handler.RegisterMaintenance,
		{Method: "GET", Path: "/maintenances"}:           handler.ListMaintenances,
		{Method: "DELETE", Path: "/maintenances"}:        handler.DeleteMaintenances,
	}

	return &MaintenanceService{
		Database: db,
		Router:   api.NewRouter(handler, paths),
	}
}

func (s *MaintenanceService) Start() {
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
