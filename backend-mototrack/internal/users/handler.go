package users 

import ("net/http"
	    //"log"
	    "io"
	    "go.mongodb.org/mongo-driver/mongo")


type UserHandler struct {
	Repository *UserRepository
	Factory *UserFactory
}

func NewUserHandler(db *mongo.Client) *UserHandler {

	return &UserHandler{
		Factory:    NewUserFactory(),
		Repository: NewUserRepository(db, "user-auth", "users"),
	}
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	user, err := handler.Factory.CreateFromJSON(body)
	
	if err != nil {
		http.Error(w, "Invalid json body: " + err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Repository.Insert(user) 

	if err != nil {
		http.Error(w, "Unable to insert user in the database", http.StatusInternalServerError)
		return
	}

	

	return

}

func (*UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {

	// Será implementado na versão 0.3v
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}