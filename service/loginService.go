package service

import (
	"sampleAppDemo/repository"
	"sampleAppDemo/utility"
)

type LoginService interface {
	Login(email string, password string) bool
}

type loginService struct {
	personRepository repository.PersonRepository
}

func (l loginService) Login(email string, password string) bool {
	result, err := l.personRepository.FindByEmail(email)
	if err != nil {
		return false
	}

	return utility.IsPasswordMatch(result.PasswordHash, password)
}

func NewLoginService() LoginService {
	return &loginService{personRepository: repository.NewPersonRepository()}
}
