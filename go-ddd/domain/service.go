package domain

// Service Public interface
type Service interface {
	Create(e *Entity)
	RetrieveAll() []*Entity
}

type service struct {
	repository Repository
}

// NewService Public factory method
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(e *Entity) {
	s.repository.Save(e)
}

func (s *service) RetrieveAll() []*Entity {
	return s.repository.List()
}
