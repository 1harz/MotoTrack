package motorcicles

import (
	"log"
	"net/http"

	"gitlab.com/prof-caio-de-melo/user-service/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type MotorcicleService struct {
	Router   *api.Router
	Database *mongo.Client
}

func NewMotorcicleService(db *mongo.Client) *MotorcicleService {

	handler := NewMotorcicleHandler(db)

	paths := map[api.Route]api.HandlerFunc{
		{Method: "POST", Path: "/motorcicles/register"}: handler.RegisterMotorcicle,
		{Method: "GET", Path: "/motorcicles"}:           handler.ListMotorcicles,
		{Method: "DELETE", Path: "/motorcicles"}:        handler.DeleteMotorcicles,
	}

	return &MotorcicleService{
		Database: db,
		Router:   api.NewRouter(handler, paths),
	}
}

func (s *MotorcicleService) Start() {
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
