package motorcicles

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type MotorcicleRepository struct {
	collection *mongo.Collection
}

func NewMotorcicleRepository(db *mongo.Client, database, collection string) *MotorcicleRepository {
	return &MotorcicleRepository{
		collection: db.Database(database).Collection(collection),
	}
}

func (repo *MotorcicleRepository) Insert(motorcicle *Motorcicle) error {
	_, err := repo.collection.InsertOne(context.Background(), motorcicle)
	return err
}

func (repo *MotorcicleRepository) FindByOwnerID(ownerId string) ([]Motorcicle, error) {
	var motorcicles []Motorcicle
	cursor, err := repo.collection.Find(context.Background(), bson.M{"ownerId": ownerId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var motorcicle Motorcicle
		if err := cursor.Decode(&motorcicle); err != nil {
			return nil, err
		}
		motorcicles = append(motorcicles, motorcicle)
	}
	return motorcicles, nil
}

func (repo *MotorcicleRepository) DeleteByOwnerID(ownerId string) error {
	_, err := repo.collection.DeleteMany(context.Background(), bson.M{"ownerId": ownerId})
	return err
}
