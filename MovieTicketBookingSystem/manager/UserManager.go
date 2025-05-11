package manager

import (
	"ticketBooking/models"
)

var userManager *UserManager

type UserManager struct {
	users map[uint64]*models.User
}

func NewUserManager() *UserManager {
	userManager = &UserManager{
		users: make(map[uint64]*models.User),
	}
	return userManager
}

func (m *UserManager) GetUserById(userId uint64) *models.User {
	return m.users[userId]
}

func (m *UserManager) CreateUser(name, email string, phone int) {
	user := models.NewUser(name, email, phone)
	m.users[user.UserId] = user
}
