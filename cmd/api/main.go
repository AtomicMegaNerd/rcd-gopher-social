package main

import (
	"log"

	env "github.com/atomicmeganerd/rcd-gopher-social/internal"
	"github.com/atomicmeganerd/rcd-gopher-social/internal/db"
	"github.com/atomicmeganerd/rcd-gopher-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{
			addr: env.GetString(
				"DB_ADDR",
				"postgres://admin:adminpassword@localhost/social?sslmode=disable",
			),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdeConns:  env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdeConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panicf("could not connect to db: %v", err)
	}

	defer db.Close()
	log.Println("Connected to database connection pool")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
