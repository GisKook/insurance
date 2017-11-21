package db

import ( 
	"database/sql"
	"github.com/giskook/insurance/base"
	"log"
)

const (
	SQL_INSURANCE_SUBJECT_ADD string = "insert into picc_subject_attr(subject_name, subject_col) values($1, $2)"
	SQL_INSURANCE_SUBJECT_EXISTS string = "select distinct subject_name from picc_subject_attr where subject_name=$1"
)

func (db_socket *DBSocket) InsuranceSubjectAdd(subject *base.Subject) error {
	var name string
	e := db_socket.db.QueryRow(SQL_INSURANCE_SUBJECT_EXISTS,subject.Name).Scan(&name)
	if e != sql.ErrNoRows{ 
		return base.New(e, base.ERR_INSURANCE_SUBJECT_ADD_EXIST_CODE, base.ERR_INSURANCE_SUBJECT_ADD_EXIST_DESC)
	}

	tx, err := db_socket.db.Begin()
	if err != nil{
		tx.Rollback()
		log.Println(err.Error())
		return base.New(err, base.ERR_DB_BEGIN_TRANSCATION_CODE, base.ERR_DB_BEGIN_TRANSCATION_DESC)
	}
	for _, v := range subject.Cols{
		_, err = tx.Exec(SQL_INSURANCE_SUBJECT_ADD, subject.Name, v)
		if err != nil{
			tx.Rollback()
			log.Println(err.Error())
			return base.New(err, base.ERR_INSURANCE_SUBJECT_ADD_CODE, base.ERR_INSURANCE_SUBJECT_ADD_DESC)
		}
	}
	err = tx.Commit()
	if err != nil{
		log.Println(err.Error())
		return base.New(err, base.ERR_DB_COMMIT_TRANSCATION_CODE, base.ERR_DB_COMMIT_TRANSCATION_DESC)
	}

	return err
}
