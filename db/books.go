package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"bookcollection.com/rest-api/properties"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type (
	Book struct {
		id            int
		ReferenceId   string    `json:"reference" validate:"required"`
		Title         string    `json:"title" validate:"required,max=100"`
		Pages         int       `json:"pages" validate:"required,max=4"`
		Author        string    `json:"author"`
		DatePublished time.Time `json:"date_published"`
		Isbn          int       `json:"isbn"`
		Publisher     string    `json:"publisher"`
		ImageUrl      string    `json:"image_url"`
		Borrowed      time.Time `json:"borrowed date"`
		BorrowedBy    string    `json:"borrowed by"`
		Read          time.Time `json:"read"`
	}
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type (
	ComicBook struct {
		StorageBox string    `json:"storage box" validate:"required"`
		SoldDate   time.Time `json:"sold date"`
		SaleURL    string    `json:"sale url"`
		Book
	}
)

var DB *sql.DB

var books []Book

var comicBooks []ComicBook

func InitDB() {
	username := properties.GetProperty("username")
	password := properties.GetProperty("password")
	host := properties.GetProperty("host")
	port := properties.GetProperty("port")

	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/BOOK_COLLECTION?parseTime=true", username, password, host, port))

	err = DB.Ping()

	if err != nil {
		log.Fatalf("Could not connect to db!!!: %s", err.Error())
	}

	log.Println("Connected to db")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}

func (b Book) Save() ([]Book, error) {
	validate = validator.New()
	err := validate.Struct(b)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if b.ValidateBookStatus() != nil {
		log.Println(err)
		return nil, err
	}

	books = append(books, b)
	uuid := uuid.New()
	query := `INSERT INTO trad_books (reference_id, title, pages, author, date_published, isbn, publisher, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := DB.Prepare(query)

	if err != nil {
		log.Printf("Could not prepare statement: %s", err.Error())
		return nil, err
	}

	exec, err := stmt.Exec(uuid, b.Title, b.Pages, b.Author, b.DatePublished, b.Isbn, b.Publisher, b.ImageUrl)
	if err != nil {
		log.Printf("Error saving Book: %s", err.Error())
		return nil, err
	}

	log.Println(exec)

	books = append(books, b)

	return books, nil
}

func GetAllTradBooks() ([]Book, error) {
	query := "SELECT * FROM trad_books"

	rows, err := DB.Query(query)

	if err != nil {
		fmt.Printf("Failed to query db: %s", err.Error())
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
		log.Printf("Failed to query db: %s", err.Error())
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

func (b Book) ValidateBookStatus() error {
	if b.DatePublished.After(b.Read) {
		return fmt.Errorf("%s cannot be Read before the book was Published", b.Title)
	}

	return nil
}

func (b ComicBook) ValidateComicBookStatus() error {
	if !b.Borrowed.IsZero() && (!b.SoldDate.IsZero() || b.SaleURL != "") {
		return fmt.Errorf("%s cannot be Borrowed while having a Sold status or have a Sale URL for a pending sale", b.Title)
	}

	return nil
}
