package maintenance

import (
	"encoding/json"
	"errors"
)

type MaintenanceFactory struct{}

func NewMaintenanceFactory() *MaintenanceFactory {
	return &MaintenanceFactory{}
}

func (factory *MaintenanceFactory) CreateFromJSON(data []byte) (*Maintenance, error) {
	var maintenance Maintenance
	err := json.Unmarshal(data, &maintenance)
	if err != nil {
		return nil, errors.New("invalid json data")
	}

	if maintenance.TypeMaintenance == "" || maintenance.Date == "" || maintenance.Description == "" || maintenance.OwnerID == "" {
		return nil, errors.New("all fields are required")
	}

	return &maintenance, nil
}
