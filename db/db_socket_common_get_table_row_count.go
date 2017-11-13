package db

import (
	"fmt"
)

const (
	SQL_GET_TABLE_ROW_COUNT string = "select count(*) from %s where 1=1 %s"
)

func (db_socket *DBSocket) CommonGetTableRowCount(table, where_clause string) (int, error) {
	var count int
	sql := fmt.Sprintf(SQL_GET_TABLE_ROW_COUNT, table, where_clause)
	err := db_socket.db.QueryRow(sql).Scan(&count)

	return count, err
}
