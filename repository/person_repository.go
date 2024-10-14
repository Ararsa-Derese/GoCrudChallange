package repository

import (
	"errors"
	"sync"
	"person/domain"
)

type PersonRepository struct {
	persons []*domain.Person
	mu      sync.Mutex
}

var (
	ErrPersonNotFound = errors.New("person not found")
)

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		persons: []*domain.Person{},
	}
}

func (r *PersonRepository) FindAll() ([]*domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.persons, nil
}

func (r *PersonRepository) FindById(id uint32) (*domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, person := range r.persons {
		if person.Id == id {
			return person, nil
		}
	}
	return nil, ErrPersonNotFound
}

func (r *PersonRepository) Save(person *domain.Person) (*domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.persons = append(r.persons, person)
	return person, nil
}

func (r *PersonRepository) Update(person *domain.Person) (*domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, p := range r.persons {
		if p.Id == person.Id {
			r.persons[i] = person
			return person, nil
		}
	}
	return nil, ErrPersonNotFound
}

func (r *PersonRepository) Delete(id uint32) (*domain.Person, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, person := range r.persons {
		if person.Id == id {
			r.persons = append(r.persons[:i], r.persons[i+1:]...)
			return person, nil
		}
	}
	return nil, ErrPersonNotFound
}