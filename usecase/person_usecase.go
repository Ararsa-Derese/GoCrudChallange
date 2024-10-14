package usecase

import (
	"person/domain"
	"person/repository"
	"time"
)

// PersonUsecase represent the person's usecases

type PersonUsecase struct {
	PersonRepository *repository.PersonRepository
	contextTimeout   time.Duration
}

// NewPersonUsecase will create new an personUsecase object representation of personUsecase uint32erface
func NewPersonUsecase(p *repository.PersonRepository, timeout time.Duration) PersonUsecase {
	return PersonUsecase{
		PersonRepository: p,
		contextTimeout:   timeout,
	}
}

// GetAll is a usecase function to get all person
func (p *PersonUsecase) GetAll() ([]*domain.Person, error) {
	persons, err := p.PersonRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return persons, nil
}

// GetById is a usecase function to get person by id
func (p *PersonUsecase) GetById(id uint32) (*domain.Person, error) {
	person, err := p.PersonRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

// Create is a usecase function to create new person
func (p *PersonUsecase) Create(person *domain.Person) (*domain.Person, error) {
	createdPerson, err := p.PersonRepository.Save(person)
	if err != nil {
		return nil, err
	}
	return createdPerson, nil
}

// Update is a usecase function to update person
func (p *PersonUsecase) Update(id uint32, person *domain.Person) (*domain.Person, error) {
	person.Id = id
	updatedPerson, err := p.PersonRepository.Update(person)
	if err != nil {
		return nil, err
	}
	return updatedPerson, nil
}

// Remove is a usecase function to remove person by id
func (p *PersonUsecase) Remove(id uint32)  error {
	_, err := p.PersonRepository.Delete(id)
	if err != nil {
		return  err
	}
	return  nil
}
