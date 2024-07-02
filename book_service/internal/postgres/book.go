package postgres

import (
	"Book/pb"
	"database/sql"
	"log"
)

func StoreNewBook(db *sql.DB, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	var book pb.AddBookResponse
	query := `INSERT INTO books(title, author) VALUES($1, $2) RETURNING id;`
	row := db.QueryRow(query, req.Title, req.Author)
	if err := row.Scan(&book.BookId); err != nil {
		log.Fatal("unable to insert book:", err)
		return nil, err
	}

	return &book, nil
}

func StoreGetBook(db *sql.DB, req *pb.GetBookRequest) (*pb.Book, error) {
	var book pb.Book
	query := `SELECT *FROM books WHERE id = $1`
	row := db.QueryRow(query, req.BookId)
	if err := row.Scan(&book.BookId, &book.Title, &book.Author); err != nil {
		log.Println("unable to get book:", err)
		return nil, err
	}

	return &book, nil
}
