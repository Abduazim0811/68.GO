package loanservice

import (
	"Loan/internal/postgres"
	"Loan/pb"
	"context"
	"database/sql"
	"log"
)

type LoanService struct{
	pb.UnimplementedLoanServiceServer
	db *sql.DB
}

func NewLoanServer(db *sql.DB) *LoanService{
	return &LoanService{db: db}
}

func (l *LoanService) IssueLoan(ctx context.Context, req *pb.IssueLoanRequest)(*pb.IssueLoanResponse, error){
	loan, err := postgres.StoreIssueleon(l.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.IssueLoanResponse{LoanId: loan.LoanId}, nil
}

func (l *LoanService) ReturnLoan(ctx context.Context, req *pb.ReturnLoanRequest)(*pb.ReturnLoanResponse, error){
	success, err := postgres.StoreReturnLoan(l.db, req)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	return &pb.ReturnLoanResponse{Success: success}, nil
}