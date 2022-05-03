package repositery

// interface for the repo/database there would be various types of queries based on your
// usecase I just implimented one.

// DAO stands for Data Access Object, This is what would grant access to the DB to servives and handles

type DAO interface {
	NewMovieQuery() MovieQuery
}

type dao struct{}

func NewDAO() DAO {
	s := &dao{}
	return s
}

func (d *dao) NewMovieQuery() MovieQuery {
	return &movieQuer{}
}
