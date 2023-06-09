package book

type BookCore struct {
	ID					uint
	Title				string
	PublishedYear		string
	ISBN				int
}

type IusecaseInterface interface {
	GetAllBook() (data []BookCore, err error)
	GetBookById(id int) (data BookCore, err error)
	CreateBook(data BookCore) (row int, err error)
	UpdateBook(id int, update BookCore) (int, error)
	DeleteBook(id int, delete BookCore) (int, error)
}

type IdataInterface interface {
	SelectAllBook() (data []BookCore, err error)
	SelectBookById(id int) (data BookCore, err error)
	InsertBook(data BookCore) (row int, err error)
	PutBook(id int, update BookCore) (int , error)
	DelBook(id int, delete BookCore) (int, error)
}