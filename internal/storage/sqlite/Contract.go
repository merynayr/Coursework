package sqlite

import "fmt"

type Contract struct {
	ContractID int
	ClientID   int
	BoxID      int
	DateSigned string
	StartDate  string
	EndDate    string
	ClientName string
}

func (s *Storage) AddContract(contract Contract) (int64, error) {
	const op = "AddContract"
	query := `INSERT INTO Contracts (contract_id, client_id, box_id, date_signed, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?)`

	res, err := s.db.Exec(query, contract.ContractID, contract.ClientID, contract.BoxID, contract.DateSigned, contract.StartDate, contract.EndDate)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}
	return id, nil
}

func (s *Storage) SelectContracts() ([]Contract, error) {
	const op = "SelectContracts"

	query := `SELECT contract_id, Clients.name, Boxes.box_id, date_signed, start_date, end_date FROM Contracts JOIN Clients ON Contracts.client_id == Clients.client_id JOIN Boxes ON Contracts.box_id = Boxes.box_id`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var contracts []Contract

	for rows.Next() {
		var contract Contract
		if err := rows.Scan(&contract.ContractID, &contract.ClientName, &contract.BoxID, &contract.DateSigned, &contract.StartDate, &contract.EndDate); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		contracts = append(contracts, contract)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contracts, nil
}
