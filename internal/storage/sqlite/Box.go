package sqlite

import (
	"Coursework/internal/storage"
	"fmt"
)

type Box struct {
	BoxID          int
	Status         string
	Floor          int
	Area           float64
	Contract_id    string
	Contract_start string
	Contract_end   string
}

func (s *Storage) AddBox(box Box) (int64, error) {
	const op = "AddBox"
	query := `INSERT INTO Boxes (box_id, status, floor, area) VALUES (?, ?, ?, ?)`

	res, err := s.db.Exec(query, box.BoxID, box.Status, box.Floor, box.Area)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrExists)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SelectBoxes() ([]Box, error) {
	const op = "SelectBoxes"

	query := `SELECT Boxes.box_id, Boxes.status, Boxes.floor, Boxes.area, Contracts.contract_id, Contracts.start_date, Contracts.end_date From Boxes JOIN Contracts ON Boxes.box_id = Contracts.box_id`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var boxes []Box

	for rows.Next() {
		var box Box
		if err := rows.Scan(&box.BoxID, &box.Status, &box.Floor, &box.Area, &box.Contract_id, &box.Contract_start, &box.Contract_end); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		boxes = append(boxes, box)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return boxes, nil
}
