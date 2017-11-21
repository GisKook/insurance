package db

import ()

const (
	SQL_INSURANCE_GET_SUBJECT string = "select distict subject_name from picc_subject_attr"
)

func (db_socket *DBSocket) InsuranceGetSubject() ([]string, error) {
	subjects := make([]string, 0)
	rows, err := db_socket.db.Query(SQL_INSURANCE_GET_SUBJECT)
	if err != nil {
		return subjects, err
	}
	defer rows.Close()

	var subject_name string
	for rows.Next() {
		if err := rows.Scan(&subject_name); err != nil {
			return subjects, err
		}
		subjects = append(subjects, subject_name)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subjects, err
}
