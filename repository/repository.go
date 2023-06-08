package repository

//no ORM has been used

import (
	"errors"
	"gin-crud/entity"
)

type Repository struct {
	DB map[string]entity.Book
}

func NewRepository() *Repository {
	example := entity.Book{
		ID:     "1",
		Tittle: "Man of Peace : The Silent Warrior",
		Year:   2027,
		Author: "_Hrishabh",
	}

	return &Repository{
		DB: map[string]entity.Book{example.ID: example},
	}
}

func (repo *Repository) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book

	for _, book := range repo.DB {
		books = append(books, book)
	}

	return books, nil
}

func (repo *Repository) GetBook(id string) (entity.Book, error) {
	if book, exist := repo.DB[id]; exist {
		return book, nil
	}

	return entity.Book{}, errors.New("not found")
}

func (repo *Repository) CreateBook(book entity.Book) (entity.Book, error) {
	if _, exist := repo.DB[book.ID]; exist {
		return entity.Book{}, errors.New("already exist")
	}

	repo.DB[book.ID] = book

	return book, nil
}

func (repo *Repository) UpdateBook(book entity.Book) error {
	if _, exist := repo.DB[book.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[book.ID] = book

	return nil
}

func (repo *Repository) DeleteBook(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
