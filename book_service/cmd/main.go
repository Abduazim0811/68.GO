package main

import (
	"Book/internal/storage"
	"Book/pb"
	bookservice "Book/service/book_service"
	"log"
	"net"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	book_service := bookservice.NewBookService(db)

	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()
	s := grpc.NewServer()

	pb.RegisterBookServiceServer(s, book_service)
	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}

}
