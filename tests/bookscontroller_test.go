package tests

import (
	"strconv"
	"testing"

	"jjrepos/golang/api/database"
	"jjrepos/golang/api/database/models"
	"jjrepos/golang/api/dtos"

	unitTest "github.com/Valiben/gin_unit_test"
	utils "github.com/Valiben/gin_unit_test/utils"
)

func TestCreateBook(t *testing.T) {
	defer teardownBooks()
	input := dtos.CreateBookDto{Title: "New Book", Author: "New Author"}
	var res models.Book
	if err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/books", "json", input, &res); err != nil {
		t.Errorf("TestCreateBook : %v\n", err)
		return
	}
	if res.Id == 0 {
		t.Errorf("Expected book id")
	}
	if res.Author != input.Author {
		t.Errorf("Expected %v but got %v", input.Author, res.Author)
	}
	if res.Title != input.Title {
		t.Errorf("Expected %v but got %v", input.Title, res.Title)
	}
}

func TestGetBook(t *testing.T) {
	books := setupBooks()
	defer teardownBooks()
	input := books[0]
	id := strconv.FormatInt(int64(input.Id), 10)
	var res models.Book
	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/books/"+id, "json", nil, &res); err != nil {
		t.Errorf("TestGetBook : %v\n", err)
	}
	if res.Id != input.Id {
		t.Errorf("Expected %v but got %v", input.Id, res.Id)
	}
	if res.Author != input.Author {
		t.Errorf("Expected %v but got %v", input.Author, res.Author)
	}
	if res.Title != input.Title {
		t.Errorf("Expected %v but got %v", input.Title, res.Title)
	}
}

func TestGetBooks(t *testing.T) {
	books := setupBooks()
	defer teardownBooks()
	var res []models.Book
	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/books", "json", nil, &res); err != nil {
		t.Errorf("TestGetBook : %v\n", err)
	}
	for _, book := range res {
		if !bookExists(book.Id, books) {
			t.Errorf("Expected response to contain book : %v\n", book)
			return
		}
	}
}

func TestUpdateBook(t *testing.T) {
	books := setupBooks()
	defer teardownBooks()
	id := strconv.FormatInt(int64(books[0].Id), 10)
	input := dtos.UpdateBookDto{Author: "New Author"}
	var res models.Book
	if err := unitTest.TestHandlerUnMarshalResp("PATCH", "/books/"+id, "json", input, &res); err != nil {
		t.Errorf("TestUpdateBook : %v\n", err)
		return
	}
	if res.Author != input.Author {
		t.Errorf("Expected %v but got %v", input.Author, res.Author)
	}

}

func TestDeleteBook(t *testing.T) {
	books := setupBooks()
	defer teardownBooks()
	id := strconv.FormatInt(int64(books[0].Id), 10)
	if _, err := unitTest.TestOrdinaryHandler(utils.DELETE, "/books/"+id, "json", nil); err != nil {
		t.Errorf("TestDeleteBook : %v\n", err)
		return
	}

	//Verify the book does not exist with a GET
	if err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/books/"+id, "json", nil, models.Book{}); err == nil {
		t.Errorf("TestDeleteBook : %v\n", err)
	}
}

func setupBooks() []models.Book {
	books := []models.Book{
		{Title: "Test Book1", Author: "Test Author1"},
		{Title: "Test Book2", Author: "Test Author2"},
		{Title: "Test Book3", Author: "Test Author3"},
		{Title: "Test Book4", Author: "Test Author4"},
		{Title: "Test Book5", Author: "Test Author5"},
		{Title: "Test Book6", Author: "Test Author6"},
		{Title: "Test Book7", Author: "Test Author7"},
		{Title: "Test Book8", Author: "Test Author8"},
	}
	database.Db.Create(&books)
	return books
}

func teardownBooks() {
	database.Db.Exec("DELETE FROM books")
}

func bookExists(id uint, books []models.Book) bool {
	for _, book := range books {
		if id == book.Id {
			return true
		}
	}
	return false
}
