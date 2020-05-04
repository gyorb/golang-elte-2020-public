package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	app "github.com/gerifield/golang-elte-2020-public/httpserver/ex8/app"
)

func main() {
	addr := flag.String("listen", ":8080", "Listening address")
	sqliteConfig := flag.String("sqlitedb", "file:httpserver/ex8/sqlite.db?cache=shared", "SQLite database connection string")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync() // flushes buffer, if any

	// Setup the DB connection
	var sqlxConn *sqlx.DB
	sqlxConn, err = sqlx.Connect("sqlite3", *sqliteConfig) // Connect also sends a Ping to check the server Open won't do that!
	if err != nil {
		logger.Fatal("database connection failure", zap.Error(err))
	}
	defer sqlxConn.Close()
	sqlxConn.SetMaxOpenConns(1)

	// Setup the app part
	application := app.NewApp(logger, sqlxConn)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello there!")
	})

	r.Mount("/api", application.Routes(chi.NewRouter()))

	// Start the HTTP server
	logger.Info("Listening", zap.String("addr", *addr))
	err = http.ListenAndServe(*addr, r)
	if err != nil {
		logger.Fatal("server failure", zap.Error(err))
	}
}
