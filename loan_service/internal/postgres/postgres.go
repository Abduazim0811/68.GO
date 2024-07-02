package postgres

import (
	"Loan/pb"
	"database/sql"
	"log"
)

func StoreIssueleon(db *sql.DB, req *pb.IssueLoanRequest)(*pb.IssueLoanResponse, error){
	var leon pb.IssueLoanResponse
	query := `INSERT INTO loans(book_id, user_id) VALUES($1, $2) RETURNING loan_id`
	row := db.QueryRow(query, req.BookId, req.UserId)
	if err := row.Scan(&leon.LoanId); err != nil{
		log.Println("unable to insert loan:", err)
		return nil, err
	}

	return &leon, nil
}

func StoreReturnLoan(db *sql.DB, req *pb.ReturnLoanRequest) (bool, error){
	query := `DELETE FROM loans WHERE loan_id = $1`
	result, err := db.Exec(query, req.LoanId)
	if err != nil {
		log.Println("unable to return loan:", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("unable to get affected rows:", err)
		return false, err
	}

	return rowsAffected > 0, nil
}