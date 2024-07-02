package bookservice

import (
	"Book/internal/postgres"
	"Book/pb"
	"context"
	"database/sql"
	"log"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (b *BookService) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	book, err := postgres.StoreNewBook(b.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.AddBookResponse{BookId: book.BookId}, nil
}

func (b *BookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	book, err := postgres.StoreGetBook(b.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.Book{BookId: book.BookId, Title: book.Title, Author: book.Author}, nil
}
