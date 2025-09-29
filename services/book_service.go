package services

import (
	"github.com/emmanuelYohore/api-golang/database"
	"github.com/emmanuelYohore/api-golang/models"
)

func GetAllBooks(books *[]models.Book) error {
	result := database.DB.Find(books)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBookByID(book *models.Book, id string) error {
	result := database.DB.Where("id = ?", id).First(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateBook(book *models.Book) error {
	createBook := database.DB.Create(book)
	if createBook.Error != nil {
		return createBook.Error
	}
	return nil
}

func UpdateBook(book *models.Book, id string) error {
	updateBook := database.DB.Model(book).Where("id = ?", id).Updates(book)
	if updateBook.Error != nil {
		return updateBook.Error
	}
	return nil
}

func DeleteBook(book *models.Book, id string) error {
	deleteBook := database.DB.Where("id = ?", id).Delete(book)
	if deleteBook != nil {
		return deleteBook.Error
	}
	return nil
}
