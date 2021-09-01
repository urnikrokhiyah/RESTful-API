package database

import (
	"mvc/config"
	"mvc/models"
)

func GetBook() (interface{}, error) {
	var books []models.Book
	err := config.DB.Find(&books).Error

	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetOneBook(id int) (interface{}, int, error) {
	var book models.Book
	err := config.DB.Find(&book, id).Error

	if err != nil {
		return nil, 0, err
	}

	return book, 1, nil
}

func CreateNewBook(book *models.Book) (interface{}, error) {
	err := config.DB.Create(&book).Error

	if err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book []models.Book
	err := config.DB.Delete(&book, id).Error
	if err != nil {
		return nil, err
	}
	return "deleted", nil
}

func UpdateBook(book interface{}, newBook models.Book) (interface{}, error) {
	err := config.DB.Model(book).Updates(newBook).Error

	if err != nil {
		return nil, err
	}

	return "updated", nil
}
