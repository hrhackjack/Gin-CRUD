package services

import (
	"gin-crud/entity"
	"gin-crud/repository"

	"github.com/google/uuid"
)

type Services struct {
	Repository *repository.Repository
}


func NewServices() *Services {
	return &Services{
		Repository: repository.NewRepository(),
	}
}

func (s *Services) GetAllBooks() ([]entity.Book, error) {
	return s.Repository.GetAllBooks()
}

func (s *Services) GetBook(id string) (entity.Book, error) {
	return s.Repository.GetBook(id)
}

func (s *Services) CreateBook(book entity.Book) (entity.Book, int, error) {
	book.ID = uuid.New().String()

	book, err := s.Repository.CreateBook(book)
	if err != nil {
		return entity.Book{}, 400, err
	}

	return book, 201, nil
}

func (s *Services) UpdateBook(book entity.Book) (entity.Book, error) {
	if err := s.Repository.UpdateBook(book); err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (s *Services) DeleteBook(id string) error {
	if err := s.Repository.DeleteBook(id); err != nil {
		return err
	}

	return nil
}
