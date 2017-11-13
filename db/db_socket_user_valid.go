package db

import (
	"crypto/md5"
	"fmt"
	"github.com/giskook/insurance/base"
	"io"
	"log"
	"strings"
)

const (
	SQL_GET_USER string = "select picc_user.id, picc_user.name,picc_user.passwd, picc_role.name, picc_entity.name,picc_role_privilege.can_add  , picc_role_privilege.can_read, picc_role_privilege.can_del, picc_role_privilege.can_update from picc_user, picc_user_role, picc_role, picc_entity, picc_role_privilege where picc_user.id=picc_user_role.user_id and picc_user_role.role_id=picc_role.id and picc_role.id=picc_role_privilege.role_id and picc_entity.id=picc_role_privilege.entity_id and picc_user.id=$1"
)

func (db_socket *DBSocket) valid_passwd(passwd_input string, secret string, passwd_store string) bool {
	m := md5.New()
	io.WriteString(m, passwd_input)
	log.Println(fmt.Sprintf("%x", m.Sum([]byte(secret))))

	if fmt.Sprintf("%x", m.Sum([]byte(secret))) == passwd_store {
		return true
	}

	return false
}

func (db_socket *DBSocket) gen_auth(entity_name string, can_add int, can_read int, can_del int, can_update int) string {
	auth := 0
	if can_add == 1 {
		auth |= 8
	}
	if can_read == 1 {
		auth |= 4
	}

	if can_del == 1 {
		auth |= 2
	}

	if can_update == 1 {
		auth |= 1
	}

	return entity_name + "." + fmt.Sprintf("%d", auth)
}

func (db_socket *DBSocket) combine_auth(auth_long string, auth2 string) string {
	if strings.Contains(auth_long, auth2) {
		return auth_long
	}
	return auth_long + "-" + auth2
}

func (db_socket *DBSocket) UserValid(user string, password string) (string, string, error) {
	rows, err := db_socket.db.Query(SQL_GET_USER, user)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	var id, name, passwd, role_name, entity_name, auth string
	var can_add, can_read, can_del, can_update int

	if rows.Next() {
		if err := rows.Scan(&id, &name, &passwd, &role_name, &entity_name, &can_add, &can_read, &can_del, &can_update); err != nil {
			return "", "", err
		}
		if db_socket.valid_passwd(password, db_socket.conf.Secret, passwd) {
			auth = db_socket.gen_auth(entity_name, can_add, can_read, can_del, can_update)
		} else {
			return "", "", base.ErrLoginPasswdError
		}
	} else {
		return "", "", base.ErrLoginUserNotFound
	}

	for rows.Next() {
		if err := rows.Scan(&id, &name, &passwd, &role_name, &entity_name, &can_add, &can_read, &can_del, &can_update); err != nil {
			return "", "", err
		}
		auth_tmp := db_socket.gen_auth(entity_name, can_add, can_read, can_del, can_update)
		auth = db_socket.combine_auth(auth, auth_tmp)
	}

	if err := rows.Err(); err != nil {
		return "", "", err
	}

	return auth, id, nil
}
