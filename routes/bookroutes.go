package routes

import (
	"net/http"

	"bookcollection.com/rest-api/db"
	"github.com/labstack/echo/v4"
)

func getTradBooks(c echo.Context) error {
	books, err := db.GetAllTradBooks()
	if err != nil {
		return c.Redirect(http.StatusInternalServerError, "/books")
	}

	return c.JSON(http.StatusOK, books)
}

func getTradBooksByRef(c echo.Context) error {
	books, err := db.GetTradBooksByReference(c.Param("ref"))
	if err != nil {
		return c.Redirect(http.StatusInternalServerError, "/books/"+c.Param("ref"))
	}

	return c.JSON(http.StatusOK, books)
}

func saveBook(c echo.Context) error {
	book := new(db.Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	//TODO: Validation coming soon

	book.Save()
	return nil
}
