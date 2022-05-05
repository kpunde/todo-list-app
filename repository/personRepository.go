package repository

import (
	"errors"
	"gorm.io/gorm"
	"sampleAppDemo/entity"
	"sampleAppDemo/utility"
)

type PersonRepository interface {
	Save(person entity.User) entity.User
	Update(person entity.User) entity.User
	Delete(person entity.User)
	FindAll() []entity.User
	FindById(id int64) entity.User
	FindByEmail(email string) (*entity.User, error)
}

type personDb struct {
}

func (d personDb) FindByEmail(email string) (*entity.User, error) {
	var person entity.User
	dbResult := utility.DB.Where("email =?", email).First(&person)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &person, nil
}

func (d personDb) Save(person entity.User) entity.User {
	utility.DB.Create(&person)
	return person
}

func (d personDb) Update(person entity.User) entity.User {
	utility.DB.Save(&person)
	return person
}

func (d personDb) Delete(person entity.User) {
	utility.DB.Delete(&person)
}

func (d personDb) FindAll() []entity.User {
	var persons []entity.User
	utility.DB.Table("users").Select("id", "first_name", "last_name", "email").Scan(&persons)
	return persons
}

func (d personDb) FindById(id int64) entity.User {
	var person entity.User
	utility.DB.Where("id =?", id).First(&person)
	return person
}

func NewPersonRepository() PersonRepository {
	return &personDb{}
}
