package sqlite

import "fmt"

type Payment struct {
	PaymentID  int
	ContractID int
	DatePaid   string
	Status     string
}

func AddPayment(storage Storage, payment Payment) error {
	const op = "AddPayment"
	query := `INSERT INTO Payments (payment_id, contract_id, date_paid, status) VALUES (?, ?, ?, ?)`

	_, err := storage.db.Exec(query, payment.PaymentID, payment.ContractID, payment.DatePaid, payment.Status)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func SelectPayments(storage Storage) ([]Payment, error) {
	const op = "SelectPayments"

	query := `SELECT payment_id, contract_id, date_paid, status FROM Payments`

	rows, err := storage.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var payments []Payment

	for rows.Next() {
		var payment Payment
		if err := rows.Scan(&payment.PaymentID, &payment.ContractID, &payment.DatePaid, &payment.Status); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		payments = append(payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return payments, nil
}
