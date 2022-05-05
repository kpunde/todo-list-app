package service

import (
	"sampleAppDemo/entity"
	"sampleAppDemo/repository"
)

type PersonService interface {
	Save(person entity.User) entity.User
	Update(person entity.User) entity.User
	Delete(person entity.User)
	FindAll() []entity.User
	FindById(id int64) entity.User
}

type personService struct {
	personRepository repository.PersonRepository
}

func (i *personService) FindById(id int64) entity.User {
	return i.personRepository.FindById(id)
}

func (i *personService) Update(person entity.User) entity.User {
	return i.personRepository.Update(person)
}

func (i *personService) Delete(person entity.User) {
	i.personRepository.Delete(person)
}

func (i *personService) Save(person entity.User) entity.User {
	return i.personRepository.Save(person)
}

func (i *personService) FindAll() []entity.User {
	return i.personRepository.FindAll()
}

func NewPersonService() PersonService {
	return &personService{personRepository: repository.NewPersonRepository()}
}
