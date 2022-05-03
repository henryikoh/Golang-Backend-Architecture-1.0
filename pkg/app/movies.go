package app

// this would contain applcation core functions
// These fucntions would not directly interact with any of the external services that is the work of /servciess
// the goal for the app package is simple to handle basic business logic using the core models
// no external pakages should be important here

import (
	"github.com/henryikoh/backend-arch/pkg/models"
)

func NumberofMovies(movies models.Movies) (number int32) {
	number = 0
	for range movies {
		number++
	}
	return
}
