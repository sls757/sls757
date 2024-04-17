// Q1. A) Write a program in GO language to accept the book details such
// as BookID, Title, Author, Price. Read and display the details of
// ‘n’ number of books

package main

import "fmt"

// Define a Book structure
type Book struct {
	BookID  int
	Title   string
	Author  string
	Price   float64
}

func main() {
	// Get the number of books from the user
	fmt.Print("Enter the number of books: ")
	var n int
	fmt.Scanln(&n)

	// Create a slice to store the book details
	books := make([]Book, n)

	// Accept book details from the user
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Book %d:\n", i+1)
		fmt.Print("BookID: ")
		fmt.Scanln(&books[i].BookID)
		fmt.Print("Title: ")
		fmt.Scanln(&books[i].Title)
		fmt.Print("Author: ")
		fmt.Scanln(&books[i].Author)
		fmt.Print("Price: ")
		fmt.Scanln(&books[i].Price)
	}

	// Display the details of the books
	fmt.Println("\nDetails of the Books:")
	for i, book := range books {
		fmt.Printf("\nBook %d:\n", i+1)
		fmt.Printf("BookID: %d\n", book.BookID)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Price: %.2f\n", book.Price)
	}
}
