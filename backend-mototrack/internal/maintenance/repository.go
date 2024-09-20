package maintenance

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type MaintenanceRepository struct {
	collection *mongo.Collection
}

func NewMaintenanceRepository(db *mongo.Client) *MaintenanceRepository {
	return &MaintenanceRepository{
		collection: db.Database("Maintenance").Collection("maintenances"),
	}
}

func (repo *MaintenanceRepository) Insert(maintenance *Maintenance) error {
	_, err := repo.collection.InsertOne(context.Background(), maintenance)
	return err
}

func (repo *MaintenanceRepository) FindByOwnerID(ownerId string) ([]Maintenance, error) {
	var maintenances []Maintenance
	cursor, err := repo.collection.Find(context.Background(), bson.M{"ownerId": ownerId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var maintenance Maintenance
		if err := cursor.Decode(&maintenance); err != nil {
			return nil, err
		}
		maintenances = append(maintenances, maintenance)
	}
	return maintenances, nil
}

func (repo *MaintenanceRepository) DeleteByOwnerID(ownerId string) error {
	_, err := repo.collection.DeleteMany(context.Background(), bson.M{"ownerId": ownerId})
	return err
}
