package controllers

import (
	"mvc/lib/database"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBookController(c echo.Context) error {
	books, err := database.GetBook()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes get all books",
		"books":    books,
	})
}

func GetOneBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, _, err := database.GetOneBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get one book",
		"book":     book,
	})
}

func CreateBookController(c echo.Context) error {
	var books models.Book
	c.Bind(&books)

	book, err := database.CreateNewBook(&books)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new book",
		"book":     book,
	})

}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, err := database.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
	})
}

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	book, row, err := database.GetOneBook(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if row > 0 {
		newBook := models.Book{}
		c.Bind(&newBook)

		_, err := database.UpdateBook(book, newBook)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success updated",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "failed to update",
	})
}
