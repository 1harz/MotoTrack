package motorcicles

import (
    "encoding/json"
    "errors"
    "log"
)

type MotorcicleFactory struct{}

func NewMotorcicleFactory() *MotorcicleFactory {
    return &MotorcicleFactory{}
}

func (f *MotorcicleFactory) CreateFromJSON(data []byte) (*Motorcicle, error) {
    var motorcicle Motorcicle
    err := json.Unmarshal(data, &motorcicle)
    if err != nil {
        log.Printf("Error unmarshalling JSON: %v", err)
        return nil, errors.New("invalid JSON format")
    }

    log.Printf("Motorcicle unmarshalled: %v", motorcicle)

    // if motorcicle.Make == "" || motorcicle.Model == "" || motorcicle.Year == 0 || motorcicle.OwnerID == "" {
    //     log.Printf("Missing required fields: %v", motorcicle)
    //     return nil, errors.New("missing required fields")
    // }

    return &motorcicle, nil
}
