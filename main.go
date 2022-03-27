package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// book represents data about a record book.
type book struct {
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed record book data.
var books = []book{
	{ISBN: "1", Title: "Blue Train", Author: "John Coltrane", Price: 56.99},
	{ISBN: "2", Title: "Jeru", Author: "Gerry Mulligan", Price: 17.99},
	{ISBN: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/", ping) // used for health checks!
	router.GET("/books", getBooks) // returns a list of books
	router.GET("/books/:isbn", getBookByISBN) // get a book by sibn
	router.POST("/books", postBooks) // post a book

	router.Run() // runs on http://localhost:8080 by default
}

// test with ping
func ping(c *gin.Context) {
	datetime := time.Now()
	c.JSON(http.StatusOK, "Runing a GO app in fargate! "+datetime.String())
	// c.IndentedJSON(http.StatusOK, books) //use this to send an indented JSON
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
	// c.IndentedJSON(http.StatusOK, books) //use this to send an indented JSON
}

func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// getBookByISBN locates the book whose ISBN value matches the ISBN parameter sent by the client, then returns that book as a response.
func getBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	//Loop over the list of books, looking for an book whose ISBN value matches the parameter.
	for _, book := range books {
		if book.ISBN == isbn {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
