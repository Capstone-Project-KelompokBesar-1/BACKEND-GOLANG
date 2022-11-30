package repositories

import (
	"ourgym/models"
	"strings"

	"gorm.io/gorm"
)

func NewUserRepository(gormDB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: gormDB,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (ur *UserRepositoryImpl) GetAll(name string) []models.User {
	var users []models.User

	ur.db.Find(&users, "is_admin = ? && name LIKE ?", false, "%"+name+"%")

	return users
}

func (ur *UserRepositoryImpl) GetOneByFilter(key string, value any) models.User {
	var user models.User

	ur.db.First(&user, key, value)

	return user
}

func (ur *UserRepositoryImpl) Create(userRequest models.User) models.User {
	var user models.User

	rec := ur.db.Create(&userRequest)

	rec.Last(&user)

	return user
}

func (ur *UserRepositoryImpl) Update(id string, userRequest models.User) models.User {
	user := ur.GetOneByFilter("id", id)

	user.Name = userRequest.Name
	user.Password = userRequest.Password
	user.Phone = userRequest.Phone
	user.Address = userRequest.Address
	user.Gender = userRequest.Gender
	user.BirthDate = userRequest.BirthDate
	user.Photo = userRequest.Photo

	rec := ur.db.Save(&user)

	rec.Last(&user)

	return user
}

func (ur *UserRepositoryImpl) Delete(id string) bool {
	user := ur.GetOneByFilter("id", models.User{}.Name)

	rec := ur.db.Delete(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (ur *UserRepositoryImpl) DeleteMany(ids string) bool {
	userIds := strings.Split(ids, ",")

	rec := ur.db.Delete(&models.User{}, "id IN (?)", userIds)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
