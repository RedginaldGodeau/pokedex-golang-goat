package datastore

import (
	"backend/pkg/env"
	"context"
	"database/sql"
	"fmt"
	"log"

	"backend/ent/gen"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDatabase(env *env.Env) *gen.Client {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName)

	// if cfg.Environment == config.Test {
	// 	applyMigrations(dataSourceName)
	// }

	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := gen.NewClient(gen.Driver(drv))

	// if cfg.Environment == config.Development {
	if err := client.Schema.Create(context.TODO()); err != nil {
		defer client.Close()
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	return client
}
