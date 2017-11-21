package db

import (
	"fmt"
	"github.com/giskook/insurance/base"
	"log"
)

const (
	SQL_INSURANCE_SUBJECT_GET string = "select subject_name, subject_col from picc_subject_attr where 1=1 %s order by id desc limit $1 offset $2"
	SQL_INSURANCE_GET_SUBJECTS_COUNT string = "select count(distinct subject_name) from picc_subject_attr"
	TABLE_INSURANCE_SUBJECT_COL_NAME = "subject_name"
)

func (db_socket *DBSocket) insurance_subject_gen_where_clause(name string) string {
	var wc string
	if name != "" {
		wc = db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, TABLE_INSURANCE_SUBJECT_COL_NAME, name)
	}

	return wc
}

func (db_socket *DBSocket) fmt_subject_sql(f, subject string) string {
	wc := db_socket.insurance_subject_gen_where_clause(subject)
	return fmt.Sprintf(f, wc)
}

func (db_socket *DBSocket) serilize_subject(subject, col string, subjects []*base.Subject) []*base.Subject{ 
	is_new := true
	for _, sub := range subjects{
		is_new = true 
		if subject == sub.Name{ 
			sub.Cols = append(sub.Cols, col)
			is_new = false
			break
		}
	}
	if is_new { 
		s := &base.Subject{
			Name:subject,
			Cols:[]string{
				col,
			},
		}
		subjects = append(subjects, s)
	}

	return subjects
}

func (db_socket *DBSocket) InsuranceSubjectGet(limit, offset int, subject string) ([]*base.Subject, error) {
	fmt_sql := db_socket.fmt_subject_sql(SQL_INSURANCE_SUBJECT_GET, subject)
	var name, col string
	subjects := make([]*base.Subject, 0)
	rows, err := db_socket.db.Query(fmt_sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&name, &col); err != nil {
			return nil, err
		} 
		
		log.Println(name)
		log.Println(col)
		subjects = db_socket.serilize_subject(name, col, subjects)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subjects, err
}

func (db_socket *DBSocket)InsuranceSubjectGetCount() (int ,error){
	var count int
	err := db_socket.db.QueryRow(SQL_INSURANCE_GET_SUBJECTS_COUNT).Scan(&count)

	return count, err
}
