package sqlite

import (
	"Coursework/internal/storage"
	"fmt"
)

type Client struct {
	ClientID int
	Name     string
	Type     string
	Phone    string
}

func (s *Storage) AddClient(client Client) (int64, error) {
	const op = "storage.Client.AddClient"
	query := `INSERT INTO Clients (client_id, name, type, phone) VALUES (?, ?, ?, ?)`

	res, err := s.db.Exec(query, client.ClientID, client.Name, client.Type, client.Phone)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrExists)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SelectClients() ([]Client, error) {
	const op = "SelectClients"

	query := `SELECT client_id, name, type, phone FROM Clients`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var clients []Client

	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.ClientID, &client.Name, &client.Type, &client.Phone); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return clients, nil
}

func (s *Storage) DeleteClient(clientID int) error {
	const op = "storage.Client.DeleteClient"
	query := `DELETE FROM Clients WHERE client_id = ?`

	_, err := s.db.Exec(query, clientID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, storage.ErrExists)
	}

	return nil
}

func (s *Storage) UpdateClient(client Client) error {
	// Подготовить SQL-запрос для обновления данных клиента
	query := "UPDATE clients SET Name = ?, Type = ?, Phone = ? WHERE ClientID = ?"

	// Выполнить SQL-запрос с передачей аргументов
	_, err := s.db.Exec(query, client.Name, client.Type, client.Phone, client.ClientID)
	if err != nil {
		return err
	}

	return nil
}
