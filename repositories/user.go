package repositories

import (
	"ourgym/models"

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

func (ur *UserRepositoryImpl) GetAll() []models.User {
	var users []models.User

	ur.db.Find(&users, "isAdmin = ?", false)

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

	rec := ur.db.Save(&user)

	rec.Last(&user)

	return user
}

func (ur *UserRepositoryImpl) Delete(id uint) bool {
	user := ur.GetOneByFilter("id", id)

	rec := ur.db.Delete(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}
