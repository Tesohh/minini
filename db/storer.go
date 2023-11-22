package db

type Query map[string]any

type Storer[T any] interface {
	One(Query) (*T, error)
	Many(Query) ([]T, error)
	Put(T) error
	Update(Query, T) error
	Delete(Query) error
}
