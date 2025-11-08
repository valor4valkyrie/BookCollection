package routes

import (
	_ "bookcollection.com/rest-api/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	books := e.Group("/books/trad")
	books.GET("/all", getTradBooks)
	books.GET("/:ref", getTradBooksByRef)
	books.PUT("/new", saveBook)
	books.PUT("/scan/:isbn", scanBook)
}
