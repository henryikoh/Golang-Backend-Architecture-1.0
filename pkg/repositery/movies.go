// This is where all things concercing data would be placed
package repositery

import (
	"github.com/henryikoh/backend-arch/pkg/models"
)

// interface for calling the data. You can add or remove as much interfaces as you like
type MovieQuery interface {
	GetMovies() models.Movies
	GetMoviesByID(id string) *models.Movie
	CreateMovie()
	UpdateMovie()
	DeleteMovie()
}

// dummy data for mov ideally data would be called from a database
var mov = models.Movies{
	{
		ID: "1", Isbn: "23243", Title: "movie 1", Director: &models.Director{FirstName: "john", LastName: "doe"},
	},
	{
		ID: "2", Isbn: "2fd43", Title: "movie 1", Director: &models.Director{FirstName: "Henry", LastName: "west"},
	},
}

type movieQuer struct{}

// there should be different types of query
func NewMovieQuery() MovieQuery {
	// I can pass a value into the p
	return &movieQuer{}
}

func (m *movieQuer) GetMovies() models.Movies {

	return mov
}
func (m *movieQuer) GetMoviesByID(id string) *models.Movie {

	var foundItem *models.Movie
	for _, item := range mov {
		if item.ID == id {
			foundItem = item
			// itm = item
			break

		}
	}

	return foundItem

}
func (m *movieQuer) CreateMovie() {

}
func (m *movieQuer) UpdateMovie() {

}
func (m *movieQuer) DeleteMovie() {

}
