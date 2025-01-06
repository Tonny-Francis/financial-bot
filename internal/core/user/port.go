package user

type UserRepository interface {
	SaveUser(user User)	error
	GetUserByEmail(email string) (User, error)
	GetUserById(id string) (User, error)
	DeleteUser(id string) error
}