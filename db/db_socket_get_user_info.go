package db

import ()

const (
	SQL_GET_USER_INFO string = "select name from picc_user where id=$1"
)

func (db_socket *DBSocket) GetUserInfo(id string) (string, error) {
	var username string
	err := db_socket.db.QueryRow(SQL_GET_USER_INFO, id).Scan(&username)

	return username, err
}
