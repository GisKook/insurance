package db

import (
	"fmt"
	"github.com/giskook/insurance/base"
)

const (
	SQL_GET_USERS        string = "select name , id, tel, prov, city, county from picc_user where 1=1 %s limit $1 offset $2"
	TABLE_USERS_COL_NAME string = "name"
	TABLE_USERS_COL_ID   string = "id"
	TABLE_USERS_COL_TEL  string = "tel"
)

func (db_socket *DBSocket) User_gen_where_clause(name, id, tel string) string {
	var wc string
	if name != "" {
		wc = db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, TABLE_USERS_COL_NAME, name)
	}

	if id != "" {
		wc += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, TABLE_USERS_COL_ID, id)
	}

	if tel != "" {
		wc += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, TABLE_USERS_COL_TEL, tel)
	}

	return wc
}

func (db_socket *DBSocket) fmt_user_sql(f, where_clause string) string {
	return fmt.Sprintf(f, where_clause)
}

func (db_socket *DBSocket) UserGet(limit, offset int, wc string) ([]*base.User, error) {
	fmt_sql := db_socket.fmt_user_sql(SQL_GET_USERS, wc)
	var name, id, tel, prov, city, county string
	users := make([]*base.User, 0)
	rows, err := db_socket.db.Query(fmt_sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&name, &id, &tel, &prov, &city, &county); err != nil {
			return nil, err
		}
		users = append(users, &base.User{
			ID:     id,
			Name:   name,
			Tel:    tel,
			Prov:   prov,
			City:   city,
			County: county,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
