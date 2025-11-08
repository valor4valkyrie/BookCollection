package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bookcollection.com/rest-api/db"
	"github.com/labstack/echo/v4"
)

func scanBook(c echo.Context) error {
	fmt.Println(c.ParamNames())
	fmt.Println(c.Param("isbn"))
	isbn, err := strconv.Atoi(c.Param("isbn"))

	if err != nil {
		log.Printf("Failed to convert isbn to int: %s", err)
		return c.Redirect(http.StatusInternalServerError, "/books/"+c.Param("isbn. Invalid ISBN"))
	}

	response, err := db.ScanBook(isbn)

	if err != nil {
		log.Printf("Failed to scan book: %s", err)
		return c.Redirect(http.StatusInternalServerError, "/books/"+c.Param("isbn"))
	}

	return c.String(http.StatusOK, string(response))
}

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

	saved, err := book.Save()

	if err != nil {
		return c.Redirect(http.StatusInternalServerError, "Error:"+err.Error())
	}

	log.Println(saved)
	return nil
}
