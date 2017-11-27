package mysql

import (
	"testing"
)

const (
	//my sql config
	MYSQL_DB_HOST = "tcp(127.0.0.1:3306)"
	MYSQL_DB_NAME = "db_htgqd"
	MYSQL_DB_USER = "root"
	MYSQL_DB_PASS = "0304"
)

func TestMySQLClient_CreateMySqlClient(t *testing.T) {
	myclient := &MySQLClient{}
	err := myclient.CreateMySqlClient(MYSQL_DB_USER, MYSQL_DB_PASS, MYSQL_DB_HOST, MYSQL_DB_NAME)
	t.Log(err)
	if (myclient.db != nil) {
		t.Log("Connect DB Complete!")
	} else {
		t.Error("Faild to connect mysql server!")
	}
	myclient.Close()
}
