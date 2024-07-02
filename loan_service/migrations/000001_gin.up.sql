CREATE TABLE loans (
    loan_id serial PRIMARY KEY,
    book_id int NOT NULL,
    user_id int NOT NULL
);