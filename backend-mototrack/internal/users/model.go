package users

import "gitlab.com/prof-caio-de-melo/user-service/internal/profiles"

type Metadata map[string]string

type User struct {
  ID       string	`json:"id"      bson:"id"`
  //Name     string           `json:"name"    bson:"name"`
  Email    string	`json:"email"   bson:"email"`
  Password string	`json:"pass"    bson:"password"`
  Bike     []MotorCicle	`json:"bike"	bson:"bike"`
  //Metadata Metadata         `json:"meta"    bson:"metadata"`
  //Profile  profiles.Profile `json:"profile" bson:"profile"` 
}
type MotorCicle struct{
	MtID             string		`json:"id"	bson:"id"`
	Name             string         `json:"name"    bson:"name"`
	CylinderCapacity int		`json:"cylinders"	bson:"cylinders"`
	LicensePlate     int 		`json:"plate"	bson:"plate"`
}


func newUser(id, name, email, password string, profile profiles.Profile, meta Metadata) *User {
  return &User{
    ID: id,
    //Name: name,
    Email: email,
    Password: password,
    //Profile: profile, 
    Metadata: meta,
  }
}

func (u *User) Authenticate(password string) bool {
  return u.Password == password
}

func (u *User) SetProfile(p profiles.Profile) {
  u.Profile = p
}

func (u *User) SetMetadataElement(key, value string) {
	u.Metadata[key] = value
}

func (u *User) GetMetadataElement(key string) (string, bool) {
  value, exists :=  u.Metadata[key]
  return value, exists
}
/* Json exemplo: 
{
    "id": "user123",
    "email": "user@example.com",
    "pass": "securepassword",
    "bike": [
        {
            "id": "bike1",
            "name": "Motorcycle A",
            "cylinders": 2,
            "plate": 12345
        },
        {
            "id": "bike2",
            "name": "Motorcycle B",
            "cylinders": 4,
            "plate": 67890
        }
    ]
}*/

