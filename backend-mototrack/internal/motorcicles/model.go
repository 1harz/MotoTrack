package motorcicles

type Motorcicle struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	Make           string `json:"make" bson:"make"`
	Model          string `json:"model" bson:"model"`
	Year           int    `json:"year" bson:"year"`
	Plate          string `json:"plate" bson:"plate"`
	Color          string `json:"color" bson:"color"`
	EngineCapacity string `json:"engineCapacity" bson:"engineCapacity"`
	Mileage        int    `json:"mileage" bson:"mileage"`
	OwnerID        string `json:"ownerId" bson:"ownerId"`
}

func NewMoto(make, model, plate, color, engineCapacity, ownerID string, mileage, year int) *Motorcicle {
	return &Motorcicle{
		Make:           make,
		Model:          model,
		Year:           year,
		Plate:          plate,
		Color:          color,
		EngineCapacity: engineCapacity,
		Mileage:        mileage,
		OwnerID:        ownerID,
	}
}
