package user

type UserCore struct {
	ID       int
	Name     string
	Email	 string 
	Password string
	BookId	 int
	AuthorId int
}

type IusecaseInterface interface {
	CreateUser(data UserCore) (row int, err error)
	LoginUser(data UserCore) (string, error)
}

type IdataInterface interface {
	RegisterUser(data UserCore) (row int, err error)
	AuthUser(email string) (userEntity UserCore, err error)
}

