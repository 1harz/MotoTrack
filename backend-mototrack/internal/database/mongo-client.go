package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoClient struct {
	URI string
	IdleTime time.Duration
	client *mongo.Client
}

func (m *MongoClient) Start() error {

	// 0. Verifica se os atributos estão setados, 
	// caso não estejam, define o valor default. 
	m.validate()
	
	// 1. Define os parametros do cliente mongo, nesse caso a URI e Idle Time
	clientOptions := options.Client().ApplyURI(m.URI).SetMaxConnIdleTime(m.IdleTime * time.Minute) 

	// 2. Tenta conectar com o mongo utilizando os parametros definidos
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// 3. Faz um ping para verificar se a conexao esta viva
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	// 4. Guarda o cliente do mongo na variavel mongoClient
	m.client = client
	return nil
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.client
}

func (m *MongoClient) CloseClient() error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(context.TODO())
}

func (m *MongoClient) validate() {
	
	if m.URI == "" {
		m.URI = "mongodb://127.0.0.1:27017/"
	}

	if m.IdleTime <= 0 {
		m.IdleTime = 5
	} 

	return
}