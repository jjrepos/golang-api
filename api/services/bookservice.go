package services

import (
	"jjrepos/gonang/api/database"
	"jjrepos/gonang/api/database/models"
	"jjrepos/gonang/api/dtos"
)

func FindBooks() []models.Book {
	var books []models.Book
	database.Db.Find(&books)
	return books
}

func CreateBook(input dtos.CreateBookDto) models.Book {
	book := models.Book{Title: input.Title, Author: input.Author}
	database.Db.Create(&book)
	return book
}

func UpdateBook(id uint, input dtos.UpdateBookDto) (models.Book, error) {
	var book models.Book
	if err := database.Db.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	database.Db.Model(&book).Updates(input)
	return FindBook(id)
}

func FindBook(id uint) (models.Book, error) {
	var book models.Book
	err := database.Db.Where("id = ?", id).First(&book).Error
	return book, err
}

func DeleteBook(id uint) error {
	book := models.Book{Id: id}
	return database.Db.Delete(&book).Error
}
