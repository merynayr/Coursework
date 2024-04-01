package sqlite

import "fmt"

type Box struct {
	BoxID      int
	BoxNumber  string
	Status     string
	Floor      int
	Area       float64
	RentAmount float64
}

func AddBox(storage Storage, box Box) error {
	const op = "AddBox"
	query := `INSERT INTO Boxes (box_id, box_number, status, floor, area, rent_amount) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := storage.db.Exec(query, box.BoxID, box.BoxNumber, box.Status, box.Floor, box.Area, box.RentAmount)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func SelectBoxes(storage Storage) ([]Box, error) {
	const op = "SelectBoxes"

	query := `SELECT box_id, box_number, status, floor, area, rent_amount FROM Boxes`

	rows, err := storage.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var boxes []Box

	for rows.Next() {
		var box Box
		if err := rows.Scan(&box.BoxID, &box.BoxNumber, &box.Status, &box.Floor, &box.Area, &box.RentAmount); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		boxes = append(boxes, box)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return boxes, nil
}
