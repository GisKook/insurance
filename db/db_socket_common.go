package db

import (
	"fmt"
)

const (
	TABLE_USER                 string = "picc_user"
	SQL_WHERE_CLAUSE_EQ_STRING string = "and %s='%s'"
)

func (db_socket *DBSocket) gen_where_clause_string(op, col string, value string) string {
	return fmt.Sprintf(op, col, value)
}
