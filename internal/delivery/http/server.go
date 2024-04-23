package http

import (
	"net/http"
	"project-testing/pkg/grace"

	"github.com/rs/cors"
)

type ProjectHandler interface {
	GetAllUser(w http.ResponseWriter, r *http.Request)
	SearchUserByNameAndttl(w http.ResponseWriter, r *http.Request)
	GetAllUserPagination(w http.ResponseWriter, r *http.Request)
	GetUserByKwn(w http.ResponseWriter, r *http.Request)
	SearchUserDataByName(w http.ResponseWriter, r *http.Request)
	SearchUserDataByKwn(w http.ResponseWriter, r *http.Request)
	SearchUserDataByKwnOrName(w http.ResponseWriter, r *http.Request)
	InsertDataUser(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Project ProjectHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
