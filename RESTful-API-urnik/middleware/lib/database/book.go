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

func UpdateBook(id int, newBook *models.Book) (interface{}, int, error) {
	var book models.Book
	tx1 := config.DB.Find(&book, id)

	if tx1.Error != nil {
		return nil, 0, tx1.Error
	}

	if tx1.RowsAffected > 0 {
		tx := config.DB.Model(&book).Updates(newBook)

		if tx.Error != nil {
			return nil, 0, tx.Error
		}

		return book, 1, nil
	}
	return nil, 0, nil
}
