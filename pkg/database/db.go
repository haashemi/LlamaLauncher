package database

import (
	"context"
	"fmt"
	"llamalauncher/pkg/database/ent"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath string

func AutoMigrateDB(path string) error {
	dbPath = path + ":ent?mode=memory&cache=shared&_fk=1"

	client, err := ent.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %s", err.Error())
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %s", err.Error())
	}
	return nil
}

func GetConnection() (*ent.Client, error) {
	return ent.Open("sqlite3", dbPath)
}
