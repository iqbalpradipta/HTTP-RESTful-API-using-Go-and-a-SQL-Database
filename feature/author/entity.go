package author

type AuthorCore struct {
	ID			uint
	Name		string
	Country		string
}

type IusecaseInterface interface{
	GetAllAuthor() (data []AuthorCore, err error)
	GetAuthorByName(name string) (data AuthorCore, err error)
	CreateAuthor(data AuthorCore) (row int, err error)
	UpdateAuthor(id int, update AuthorCore) (int, error)
	DeleteAuthor(id int, delete AuthorCore) (int, error)
}

type IdataInterface interface {
	SelectAllAuthor() (data []AuthorCore, err error)
	SelectAuthorByName(name string) (data AuthorCore, err error)
	InsertAuthor(data AuthorCore) (row int, err error)
	PutAuthor(id int, update AuthorCore) (int , error)
	DelAuthor(id int, delete AuthorCore) (int, error)
}

