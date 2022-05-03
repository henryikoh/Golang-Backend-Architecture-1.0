package myhttp

import (
	"net/http"

	"github.com/henryikoh/backend-arch/pkg/services"
)

// its important realise this is a closure pattern, what this means is anything decleared before the turn is persisted
// in memory as long as the function is still active
func (s *Server) GetMovies() http.HandlerFunc {
	/*

		**prepare a thing**

		any data declared here is persisted in memory for all function calls.
		This can be used to call func specific dependecies


	*/

	// how to interact with services example note that the DAO is passed into the service

	num := services.MovieService(s.db).NumberofMovies()

	return func(w http.ResponseWriter, r *http.Request) {
		println(num)
		movies := s.db.NewMovieQuery().GetMovies()
		s.ToJSON(w, movies)

	}
}

func (s *Server) GetMovie() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "json")
		id := r.URL.Query().Get("id")
		movie := s.db.NewMovieQuery().GetMoviesByID(id)

		s.ToJSON(w, movie)

	}
}
