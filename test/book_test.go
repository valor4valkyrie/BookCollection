package test

import (
	"fmt"
	"testing"
	"time"

	"bookcollection.com/rest-api/db"
)

func TestBookBorrowedOnSaleOrSold(t *testing.T) {

	var book = db.Book{
		ReferenceId: "abc1234", Title: "Blah", Pages: 122, Author: "Blah Blah", DatePublished: time.Now(), Isbn: 123456789, Publisher: "Blah Blah", ImageUrl: "http://blah.com", Borrowed: time.Now(), BorrowedBy: "Blah Blah",
	}

	var comicBook = db.ComicBook{
		StorageBox: "Blah Blah",
		SoldDate:   time.Now(),
		SaleURL:    "http://blah.com",
		Book:       book,
	}

	err := comicBook.ValidateComicBookStatus()

	if err == nil || err.Error() != fmt.Sprintf("%s cannot be Borrowed while having a Sold status or have a Sale URL for a pending sale", book.Title) {
		t.Error("Expected an error for this: ", err)
	}
}

func TestBookBorrowedNotOnSaleOrSold(t *testing.T) {

	var book = db.Book{
		ReferenceId: "abc1234", Title: "Blah", Pages: 122, Author: "Blah Blah", DatePublished: time.Now(), Isbn: 123456789, Publisher: "Blah Blah", ImageUrl: "", Borrowed: time.Time{}, BorrowedBy: "Blah Blah",
	}

	var comicBook = db.ComicBook{
		StorageBox: "Blah Blah",
		SoldDate:   time.Now(),
		SaleURL:    "http://blah.com",
		Book:       book,
	}

	err := comicBook.ValidateComicBookStatus()

	if err != nil {
		t.Error("No error expected for this book status: ", err)
	}
}

func TestBookReadBeforePublished(t *testing.T) {

	var book = db.Book{
		ReferenceId: "abc1234", Title: "Blah", Pages: 122, Author: "Blah Blah", DatePublished: time.Now().AddDate(0, 0, 1), Isbn: 123456789, Publisher: "Blah Blah", ImageUrl: "", Borrowed: time.Now(), BorrowedBy: "Blah Blah", Read: time.Now(),
	}

	err := book.ValidateBookStatus()

	if err != nil || err.Error() != fmt.Sprintf("%s cannot be Read before the book was Published", book.Title) {
		t.Error("Missing or incorrect error for this book status, Expected Published/Read error: ", err)
	}
}
