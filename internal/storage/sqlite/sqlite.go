package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tables := []string{
		`CREATE TABLE IF NOT EXISTS Clients (
   client_id INT PRIMARY KEY,
   name VARCHAR(100),
   type VARCHAR(50),
   phone VARCHAR(20)
  );`,
		`CREATE TABLE IF NOT EXISTS Boxes (
   box_id INT PRIMARY KEY,
   status VARCHAR(50),
   floor INT,
   area DECIMAL(10, 2),
   rent_amount DECIMAL(10, 2)
  );`,
		`CREATE TABLE IF NOT EXISTS Contracts (
   contract_id INT PRIMARY KEY,
   client_id INT,
   box_id INT,
   date_signed DATE,
   start_date DATE,
   end_date DATE,
   FOREIGN KEY (client_id) REFERENCES Clients (client_id),
   FOREIGN KEY (box_id) REFERENCES Boxes (box_id)
  );`,
		`CREATE TABLE IF NOT EXISTS Payments (
   payment_id INT PRIMARY KEY,
   contract_id INT,
   date_paid DATE,
   status VARCHAR(50),
   FOREIGN KEY (contract_id) REFERENCES Contracts (contract_id)
  );`,
	}

	for _, table := range tables {
		stmt, err := db.Prepare(table)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		_, err = stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return &Storage{db: db}, nil
}
