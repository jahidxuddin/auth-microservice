package main

import (
	"auth-microservice/ent"
	"context"
	"log"
)

type DBClient struct {
	Client *ent.Client
}

func NewDatabase(dataSourceName string) *Database {
	db, err := NewDBClient(dataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize DB client: %v", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	ctx := context.Background()

	return &Database{
		dbClient: db,
		ctx: ctx,
	}
}

func NewDBClient(dataSourceName string) (*DBClient, error) {
	client, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DBClient{Client: client}, nil
}

func (db *DBClient) Close() {
	if err := db.Client.Close(); err != nil {
		log.Fatalf("failed closing connection to db: %v", err)
	}
}

func (db *DBClient) Migrate() error {
	ctx := context.Background()
	if err := db.Client.Schema.Create(ctx); err != nil {
		return err
	}
	return nil
}
