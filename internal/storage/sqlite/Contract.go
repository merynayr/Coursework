package sqlite

import "fmt"

type Contract struct {
	ContractID int
	ClientID   int
	BoxID      int
	DateSigned string
	StartDate  string
	EndDate    string
}

func AddContract(storage Storage, contract Contract) error {
	const op = "AddContract"
	query := `INSERT INTO Contracts (contract_id, client_id, box_id, date_signed, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := storage.db.Exec(query, contract.ContractID, contract.ClientID, contract.BoxID, contract.DateSigned, contract.StartDate, contract.EndDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func SelectContracts(storage Storage) ([]Contract, error) {
	const op = "SelectContracts"

	query := `SELECT contract_id, client_id, box_id, date_signed, start_date, end_date FROM Contracts`

	rows, err := storage.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var contracts []Contract

	for rows.Next() {
		var contract Contract
		if err := rows.Scan(&contract.ContractID, &contract.ClientID, &contract.BoxID, &contract.DateSigned, &contract.StartDate, &contract.EndDate); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contracts = append(contracts, contract)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contracts, nil
}
