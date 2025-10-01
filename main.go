package main

import (
	"bookcollection.com/rest-api/db"
	"bookcollection.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
	//var book = db.Book{ReferenceId: "abc23", Title: "Westinster", Pages: 324, Author: "Grisholm", DatePublished: time.Now(), Isbn: 1234, Publisher: "Clearing House", ImageUrl: "http://image.com"}
	/*	var result, err = db.GetAllBooks()
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Print(result)
		}*/

}
