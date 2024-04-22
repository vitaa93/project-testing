package boot

import (
	"log"
	"net/http"
	"project-testing/internal/config"
	projectData "project-testing/internal/data/project"
	projectSvc "project-testing/internal/service/project"

	projectServer "project-testing/internal/delivery/http"
	projectHandler "project-testing/internal/delivery/http/project"

	"github.com/jmoiron/sqlx"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}

	cfg := config.Get()

	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}
	pd := projectData.New(db)
	ps := projectSvc.New(pd)
	ph := projectHandler.New(ps)

	s := projectServer.Server{
		Project: ph,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
