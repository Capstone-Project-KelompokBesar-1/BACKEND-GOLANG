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
	user.Phone = userRequest.Phone
	user.Address = userRequest.Address
	user.Gender = userRequest.Gender
	user.BirthDate = userRequest.BirthDate
	user.Photo = userRequest.Photo

	rec := ur.db.Save(&user)

	rec.Last(&user)

	return user
}

func (ur *UserRepositoryImpl) ChangePassword(id string, newPassword string) bool {
	user := ur.GetOneByFilter("id", id)

	user.Password = newPassword

	rec := ur.db.Save(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (ur *UserRepositoryImpl) Delete(id string) bool {
	user := ur.GetOneByFilter("id", id)

	rec := ur.db.Select("Transactions").Delete(&user)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (ur *UserRepositoryImpl) DeleteMany(ids string) bool {
	userIds := strings.Split(ids, ",")

	var users []models.User

	ur.db.Find(&users, "id IN (?)", userIds)

	rec := ur.db.Select("Transactions").Delete(&users, "id IN (?)", userIds)

	if rec.RowsAffected == 0 {
		return false
	}

	return true
}

func (ur *UserRepositoryImpl) CountUser() int64 {
	var total int64

	ur.db.Find(&models.User{}, "is_admin = ?", false).Count(&total)

	return total
}
