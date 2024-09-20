package maintenance

type Maintenance struct {
	TypeMaintenance string  `json:"tm" bson:"tm"`
	Date            string  `json:"date" bson:"date"`
	Value           float64 `json:"value" bson:"value"`
	Description     string  `json:"description" bson:"description"`
	OwnerID         string  `json:"ownerId" bson:"ownerId"`
}

func NewMaintenance(tm, date, description, ownerId string, value float64) *Maintenance {
	return &Maintenance{
		TypeMaintenance: tm,
		Date:            date,
		Value:           value,
		Description:     description,
		OwnerID:         ownerId,
	}
}
