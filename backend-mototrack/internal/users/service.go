package users 

import ("log"
	    "net/http"
	    "gitlab.com/prof-caio-de-melo/user-service/api"
	    "go.mongodb.org/mongo-driver/mongo")

type UserService struct {
	Router *api.Router
	Database *mongo.Client
}

func NewUserService(db *mongo.Client) *UserService {
	
	handler := NewUserHandler(db)

	paths := map[api.Route]api.HandlerFunc{
		api.Route{"POST", "/user"}:      handler.Register,
		api.Route{"POST", "/user/auth"}: handler.AuthenticateUser,
	}

	return &UserService{
		Database: db,
		Router: api.NewRouter(handler, paths),
	}

}

func (s *UserService) Start() {

	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

