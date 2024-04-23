package http

import (
	"errors"
	"log"
	"net/http"
	"project-testing/pkg/response"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Health Check
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Routes
	project := r.PathPrefix("/project-testing").Subrouter()

	project.HandleFunc("/getalluser", s.Project.GetAllUser).Methods("GET")
	project.HandleFunc("/getuserdatasearchby", s.Project.SearchUserByNameAndttl).Methods("GET")
	project.HandleFunc("/getalluserpagination", s.Project.GetAllUserPagination).Methods("GET")
	project.HandleFunc("/searchuserbyname", s.Project.SearchUserDataByName).Methods("GET")
	project.HandleFunc("/searchuserbykwn", s.Project.SearchUserDataByKwn).Methods("GET")
	project.HandleFunc("/searchuserbykwnorname", s.Project.SearchUserDataByKwnOrName).Methods("GET")

	project.HandleFunc("/insertdatauser", s.Project.InsertDataUser).Methods("POST")

	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
