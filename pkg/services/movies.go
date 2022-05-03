package services

// Most the communication better internal app and external functions would occur in the package
// The service folder is the connector between differnet packages both internal and external

// Services need to interact with external or third party providers these can be put in a seperate folder call "adaptors"
// or they can be kept in the service folder but with a name that make it clear that it is an external package

// The service might have a main file to initialize all the services its currently implemented individually

import (
	"github.com/henryikoh/backend-arch/pkg/app"
	"github.com/henryikoh/backend-arch/pkg/repositery"
)

// the service accepts the DAO to allow it communicate with the database
// you can import service specific dependecy like this

// check mailer as an example
type Movies struct {
	db repositery.DAO
	// the mailer service can be an interface that impliments bacis email jobs and it can be used along side the movies service to send emails
	// mailer mailerService
}

// New serve is used to created the server and also assign the routes the the router
func MovieService(doa repositery.DAO) *Movies {

	s := &Movies{
		// other server variables can also be set here.
		db: doa,
	}
	return s
}

func (s *Movies) NumberofMovies() (number int32) {
	// the server is stared here
	mov := s.db.NewMovieQuery().GetMovies()
	return app.NumberofMovies(mov)

}
