package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	id            int
	ReferenceId   string
	Title         string
	Pages         int
	Author        string
	DatePublished time.Time
	Isbn          int
	Publisher     string
	ImageUrl      string
	Read          bool
}

type ComicBook struct {
	storageBox string
	sold       time.Time
	borrowed   time.Time
	borrowedBy string
	isForSale  bool
	saleURL    string
	Book
}

var DB *sql.DB

const mysqlUrl = "root:test@tcp(localhost:3306)/BOOK_COLLECTION?parseTime=true"

var books []Book

var comicBooks []ComicBook

func InitDB() {
	var err error

	DB, err = sql.Open("mysql", mysqlUrl)

	if err != nil {
		panic("Could not connect to db!!!")
	}

	fmt.Println("Connected to db")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func (b Book) Save() []Book {
	books = append(books, b)
	//query := `INSERT INTO trad_books (reference_id, title, pages, author, date_published, isbn, publisher, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)`
	//stmt, err := DB.Prepare(query)
	/*defer DB.Close()
	if err != nil {
		panic(err)
	}

	stmt.Exec(b.ReferenceId, b.Title, b.Pages, b.Author, b.DatePublished, b.Isbn, b.Publisher, b.ImageUrl)

	books = append(books, b)
	fmt.Println("Book saved")*/
	fmt.Println(books)
	return books
}

func GetAllTradBooks() ([]Book, error) {
	query := "SELECT * FROM trad_books"

	rows, err := DB.Query(query)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.id, &book.ReferenceId, &book.Title, &book.Pages, &book.Author, &book.DatePublished, &book.Isbn, &book.Publisher, &book.ImageUrl)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func GetTradBooksByReference(ref string) ([]Book, error) {
	query := "SELECT * FROM trad_books WHERE reference_id = ?"

	rows, err := DB.Query(query, ref)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.id, &book.ReferenceId, &book.Title, &book.Pages, &book.Author, &book.DatePublished, &book.Isbn, &book.Publisher, &book.ImageUrl)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
