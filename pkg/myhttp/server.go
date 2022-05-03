/*
	This page is the main server of the project. It would contain all the routers and handlers
	on this sever. They may be a need for multiple servers if this page grows too large.

	All the handler functions are attached to the server this makes it easier to share dependecies and also
	keeps evethings in one place. Like middlewares and helpers. see example below

	could also be named rest handlers or myhttp or routes
*/

// this was orginially named "handlers" but I changed the name to "myhttp" as I feel this might be more descriptive
package myhttp

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henryikoh/backend-arch/pkg/repositery"
)

/*
	The server is a struct and would contain the dependencies its need within the struct also server varieables
	can be definded here.

*/

type Server struct {
	// Router is currently the default server mux, This could also be gorlaar mux or and servemux struct
	// This type would change but ideally this should be a type that is a router
	mux *mux.Router

	router *http.ServeMux

	// db would contain the current database ( its currently being used to instansiate the Movies repo )
	db repositery.DAO
}

// New serve is used to created the server and also assign the routes the the router
func NewServer(doa repositery.DAO) *Server {
	// http.NewServeMux().router.
	s := &Server{
		// other server variables can also be set here. Check serve struck docs for more info
		// if the variable is capital it would be exported.

		// I implemented the defaul server mux which works
		router: http.NewServeMux(),
		db:     doa,

		// I also tried mux which is why it is exported
		mux: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// My server can be used has an http handlers because it impliments the serverHTTP interface
	// the server is stared here

	// Once the serve is called it is rirected to your router of choice
	// I implemented with but MUX and the default router. ( its your picks )

	// s.router.ServeHTTP(w, r)

	s.mux.ServeHTTP(w, r)
}

// here us a list of all routes
func (s *Server) routes() {
	// the routes a are unique because they are not handler functions directly
	// but ther are fucntions that return handler functions. this results in a
	// closure that can be used to do other cool things.

	// Here is the mux example

	s.mux.HandleFunc("/movies", s.GetMovies()).Methods("GET")

	// Default httpmux router here

	s.router.HandleFunc("/movies", s.GetMovies())
	s.router.HandleFunc("/movie", s.GetMiddleware(s.GetMovie()))
	// s.router.HandleFunc("/movie/{id}", getMovie)
	// s.router.HandleFunc("/movies", createMovie)
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
}

// Middleware
// Middleware in go is just a functions that takes in a handler function and return a handler function

// This functions just chceks to see if the request is a get request
func (s *Server) GetMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusBadRequest)
			return
		}
		h(w, r)
	}
}

// helpers are fucntions that can help handle repeared tasks like writting a json object to the reposone writer
func (s *Server) ToJSON(w http.ResponseWriter, mov interface{}) error {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	return e.Encode(mov)
}
