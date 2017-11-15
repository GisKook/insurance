package db

import ( 
	"database/sql"
	"github.com/giskook/insurance/base"
	"log"
)

const (
	SQL_USER_ADD string = "insert into picc_user(id, tel, name, passwd, prov, city, county, post_code) values($1, $2, $3, $4, $5, $6, $7, $8)"
	SQL_USER_ROLE_ADD string = "insert into picc_user_role(user_id, role_id) values($1, $2)"
	SQL_USER_NAME_OR_ID_EXISTS string = "select name, id from picc_user where id=$1 or name=$2"
)

func (db_socket *DBSocket) UserAdd(user *base.User) error {
	var name,id string
	e := db_socket.db.QueryRow(SQL_USER_NAME_OR_ID_EXISTS, user.ID, user.Name).Scan(&name, &id)
	if e != sql.ErrNoRows{ 
		return base.New(e, base.ERR_USER_ID_OR_TEL_EXIST_CODE, base.ERR_USER_ID_OR_TEL_EXIST_DESC)
	}

	tx, err := db_socket.db.Begin()
	if err != nil{
		tx.Rollback()
		log.Println(err.Error())
		return base.New(err, base.ERR_DB_BEGIN_TRANSCATION_CODE, base.ERR_DB_BEGIN_TRANSCATION_DESC)
	}
	_, err = tx.Exec(SQL_USER_ADD, user.ID, user.Tel, user.Name, user.Passwd, user.Prov, user.City, user.County, user.PostCode)
	if err != nil{
		tx.Rollback()
		log.Println(err.Error())
		return base.New(err, base.ERR_USER_ADD_CODE, base.ERR_USER_ADD_DESC)
	}
	for _, v := range user.Roles{
		_, err = tx.Exec(SQL_USER_ROLE_ADD, user.ID, v)
		if err != nil{
			tx.Rollback()
			log.Println(err.Error())
			return base.New(err, base.ERR_USER_ADD_CODE, base.ERR_USER_ADD_DESC)
		}
	}
	err = tx.Commit()
	if err != nil{
		log.Println(err.Error())
		return base.New(err, base.ERR_DB_COMMIT_TRANSCATION_CODE, base.ERR_DB_COMMIT_TRANSCATION_DESC)
	}

	return err
}
