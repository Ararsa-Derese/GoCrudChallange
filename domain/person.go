package domain

import (
	"github.com/google/uuid"
)

type Person struct {
	Id      uint32
	Name    string
	Age     uint
	Hobbies []string
}

type PersonConfig struct {
	Name    string
	Age     uint
	Hobbies []string
}

// NewPerson will create new person object
func NewPerson(config PersonConfig) *Person {
	return &Person{
		Id:      uuid.New().ID(),
		Name:    config.Name,
		Age:     config.Age,
		Hobbies: config.Hobbies,
	}
}

type PersonRepository interface {
	FindAll() ([]*Person, error)
	FindById(id int) (*Person, error)
	Save(person *Person) (*Person, error)
	Update(person *Person) (*Person, error)
	Delete(id int) (*Person, error)
}

type PersonUseCase interface {
	GetAll() ([]*Person, error)
	GetById(id int) (*Person, error)
	Create(person *Person) (*Person, error)
	Update(person *Person) (*Person, error)
	Remove(id int) (*Person, error)
}
